/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package block

import (
	"github.com/IBAX-io/go-ibax/packages/transaction"
	"github.com/IBAX-io/go-ibax/packages/types"
	"github.com/IBAX-io/go-ibax/packages/utils"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var (
	ErrIncorrectRollbackHash = errors.New("Rollback hash doesn't match")
	ErrEmptyBlock            = errors.New("Block doesn't contain transactions")
	ErrIncorrectBlockTime    = utils.WithBan(errors.New("Incorrect block time"))
)

// Block is storing block data
type Block struct {
	*types.BlockData
	PrevRollbacksHash []byte
	Transactions      []*transaction.Transaction
	SysUpdate         bool
	GenBlock          bool // it equals true when we are generating a new block
	Notifications     []types.Notifications
}

// GetLogger is returns logger
func (b *Block) GetLogger() *log.Entry {
	return log.WithFields(log.Fields{"block_id": b.Header.BlockId, "block_time": b.Header.Timestamp, "block_wallet_id": b.Header.KeyId,
		"block_state_id": b.Header.EcosystemId, "block_hash": b.Header.BlockHash, "block_version": b.Header.Version})
}

func (b *Block) IsGenesis() bool {
	return b.Header.BlockId == 1
}

func (b *Block) limitMode() transaction.LimitMode {
	if b == nil {
		return transaction.GetLetPreprocess()
	}
	if b.GenBlock {
		return transaction.GetLetGenBlock()
	}
	return transaction.GetLetParsing()
}

func (b *Block) Play(dbTx *sqldb.DbTransaction) (batchErr error) {
	var (
		playTxs sqldb.AfterTxs
	)
	logger := b.GetLogger()
	limits := transaction.NewLimits(b.limitMode())
	rand := random.NewRand(b.Header.Time)
	processedTx := make([]*transaction.Transaction, 0, len(b.Transactions))
	defer func() {
		if b.GenBlock {
			b.Transactions = processedTx
		}
		if err := sqldb.AfterPlayTxs(dbTx, b.Header.BlockID, playTxs, b.GenBlock, b.IsGenesis()); err != nil {
			batchErr = err
			return
		}
	}()

	for curTx, t := range b.Transactions {
		err := dbTx.Savepoint(consts.SetSavePointMarkBlock(curTx))
		if err != nil {
			logger.WithFields(log.Fields{"type": consts.DBError, "error": err, "tx_hash": t.Hash()}).Error("using savepoint")
			return err
		}

		t.Notifications = notificator.NewQueue()
		t.DbTransaction = dbTx
		t.TxCheckLimits = limits
		t.PreBlockData = b.PrevHeader
		t.GenBlock = b.GenBlock
		t.SqlDbSavePoint = curTx
		t.Rand = rand.BytesSeed(t.Hash())
		err = t.Play()
		if err != nil {
			if err == transaction.ErrNetworkStopping {
				// Set the node in a pause state
				node.PauseNodeActivity(node.PauseTypeStopingNetwork)
				return err
			}
			errRoll := t.DbTransaction.RollbackSavepoint(consts.SetSavePointMarkBlock(curTx))
			if errRoll != nil {
				t.GetLogger().WithFields(log.Fields{"type": consts.DBError, "error": err, "tx_hash": t.Hash()}).Error("rolling back to previous savepoint")
				return errRoll
			}
			if b.GenBlock {
				if err == transaction.ErrLimitStop {
					if curTx == 0 {
						return err
					}
					break
				}
				if strings.Contains(err.Error(), script.ErrVMTimeLimit.Error()) {
					err = script.ErrVMTimeLimit
				}
			}
			if t.IsSmartContract() {
				transaction.BadTxForBan(t.KeyID())
			}
			_ = transaction.MarkTransactionBad(t.DbTransaction, t.Hash(), err.Error())
			if t.SysUpdate {
				if err := syspar.SysUpdate(t.DbTransaction); err != nil {
					log.WithFields(log.Fields{"type": consts.DBError, "error": err}).Error("updating syspar")
					return err
				}
				t.SysUpdate = false
			}
			if b.GenBlock {
				continue
			}

			return err
		}

		if t.SysUpdate {
			b.SysUpdate = true
			t.SysUpdate = false
		}
		if err := sqldb.SetTransactionStatusBlockMsg(t.DbTransaction, t.BlockData.BlockID, t.TxResult, t.Hash()); err != nil {
			t.GetLogger().WithFields(log.Fields{"type": consts.DBError, "error": err, "tx_hash": t.Hash()}).Error("updating transaction status block id")
			return err
		}
		if t.Notifications.Size() > 0 {
			b.Notifications = append(b.Notifications, t.Notifications)
		}
		playTxs.UsedTx = append(playTxs.UsedTx, t.Hash())
		playTxs.TxExecutionSql = append(playTxs.TxExecutionSql, t.DbTransaction.ExecutionSql...)
		var (
			eco      int64
			contract string
		)
		if t.IsSmartContract() {
			eco = t.SmartContract().TxSmart.EcosystemID
			contract = t.SmartContract().TxContract.Name
		}
		playTxs.Lts = append(playTxs.Lts, &sqldb.LogTransaction{
			Block:        b.Header.BlockID,
			Hash:         t.Hash(),
			TxData:       t.FullData,
			Timestamp:    t.Timestamp(),
			Address:      t.KeyID(),
			EcosystemID:  eco,
			ContractName: contract,
		})
		playTxs.Rts = append(playTxs.Rts, t.RollBackTx...)
		processedTx = append(processedTx, t)
	}
	return nil
}

// Check is checking block
func (b *Block) Check() error {
	if b.IsGenesis() {
		return nil
	}
	logger := b.GetLogger()
	if b.PrevHeader == nil || b.PrevHeader.BlockID != b.Header.BlockID-1 {
		if err := b.readPreviousBlockFromBlockchainTable(); err != nil {
			logger.WithFields(log.Fields{"type": consts.InvalidObject}).Error("block id is larger then previous more than on 1")
			return err
		}
	}
	if b.Header.Time > time.Now().Unix() {
		logger.WithFields(log.Fields{"type": consts.ParameterExceeded}).Error("block time is larger than now")
		return ErrIncorrectBlockTime
	}

	// is this block too early? Allowable error = error_time
	if b.PrevHeader != nil {
		var (
			exists bool
			err    error
		)
		if syspar.GetRunModel() == consts.HonorNodeMode {
			// skip time validation for first block
			exists, err = protocols.NewBlockTimeCounter().BlockForTimeExists(time.Unix(b.Header.Time, 0), int(b.Header.NodePosition))
		}
		if err != nil {
			logger.WithFields(log.Fields{"type": consts.BlockError, "error": err}).Error("calculating block time")
			return err
		}

		if exists {
			logger.WithFields(log.Fields{"type": consts.BlockError, "error": err}).Warn("incorrect block time")
			return utils.WithBan(fmt.Errorf("%s %d", ErrIncorrectBlockTime, b.PrevHeader.Time))
		}
	}

	// check each transaction
	txCounter := make(map[int64]int)
	txHashes := make(map[string]struct{})
	for i, t := range b.Transactions {
		hexHash := string(converter.BinToHex(t.Hash()))
		// check for duplicate transactions
		if _, ok := txHashes[hexHash]; ok {
			logger.WithFields(log.Fields{"tx_hash": hexHash, "type": consts.DuplicateObject}).Warning("duplicate transaction")
			return utils.ErrInfo(fmt.Errorf("duplicate transaction %s", hexHash))
		}
		txHashes[hexHash] = struct{}{}

		// check for max transaction per user in one block
		txCounter[t.KeyID()]++
		if txCounter[t.KeyID()] > syspar.GetMaxBlockUserTx() {
			return utils.WithBan(utils.ErrInfo(fmt.Errorf("max_block_user_transactions")))
		}

		err := t.Check(b.Header.Time)
		if err != nil {
			transaction.MarkTransactionBad(t.DbTransaction, t.Hash(), err.Error())
			delete(txHashes, hexHash)
			b.Transactions = append(b.Transactions[:i], b.Transactions[i+1:]...)
			return errors.Wrap(err, "check transaction")
		}
	}

	// hash compare could be failed in the case of fork
	_, err := b.CheckHash()
	if err != nil {
		transaction.CleanCache()
		return err
	}
	return nil
}

// CheckHash is checking hash
func (b *Block) CheckHash() (bool, error) {
	logger := b.GetLogger()
	if b.IsGenesis() {
		return true, nil
	}
	if conf.Config.IsSubNode() {
		return true, nil
	}
	// check block signature
	if b.PrevHeader != nil {
		var (
			nodePublicKey   []byte
			resultCheckSign bool
			err             error
		)

		nodePublicKey, err = syspar.GetNodePublicKeyByPosition(b.Header.NodePosition)
		if err != nil {
			return false, utils.ErrInfo(err)
		}
		if len(nodePublicKey) == 0 {
			logger.WithFields(log.Fields{"type": consts.EmptyObject}).Error("node public key is empty")
			return false, utils.ErrInfo(fmt.Errorf("empty nodePublicKey"))
		}

		signSource := b.Header.ForSign(b.PrevHeader, b.MrklRoot)

		resultCheckSign, err = utils.CheckSign(
			[][]byte{nodePublicKey},
			[]byte(signSource),
			b.Header.Sign,
			true)

		if err != nil {
			if err == asymalgo.ErrIncorrectSign {
				if !bytes.Equal(b.PrevRollbacksHash, b.PrevHeader.RollbacksHash) {
					return false, ErrIncorrectRollbackHash
				}
			}
			logger.WithFields(log.Fields{"error": err, "type": consts.CryptoError}).Error("checking block header sign")
			return false, utils.ErrInfo(fmt.Errorf("err: %v / block.PrevHeader.BlockID: %d /  block.PrevHeader.Hash: %x / ", err, b.PrevHeader.BlockID, b.PrevHeader.Hash))
		}

		return resultCheckSign, nil
	}

	return true, nil
}

// InsertBlockWOForks is inserting blocks
func InsertBlockWOForks(data []byte, genBlock, firstBlock bool) error {
	block, err := ProcessBlockWherePrevFromBlockchainTable(data, !firstBlock)
	block, err := ProcessBlockWherePrevFromBlockchainTable(data, !firstBlock)
	block, err := ProcessBlockByBinData(data, !firstBlock)
	if err != nil {
		return err
	}

	block.GenBlock = genBlock
	if err := block.Check(); err != nil {
		return err
	}

	err = block.PlaySafe()
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{"block_id": block.Header.BlockID}).Debug("block was inserted successfully")
	log.WithFields(log.Fields{"block_id": block.Header.BlockID}).Debug("block was inserted successfully")
	log.WithFields(log.Fields{"block_id": block.Header.BlockId}).Debug("block was inserted successfully")
	return nil
}

// ProcessBlockWherePrevFromBlockchainTable is processing block with in table previous block
func ProcessBlockWherePrevFromBlockchainTable(data []byte, checkSize bool) (*Block, error) {
	if checkSize && int64(len(data)) > syspar.GetMaxBlockSize() {
		log.WithFields(log.Fields{"check_size": checkSize, "size": len(data), "max_size": syspar.GetMaxBlockSize(), "type": consts.ParameterExceeded}).Error("binary block size exceeds max block size")
		return nil, utils.WithBan(types.ErrMaxBlockSize(syspar.GetMaxBlockSize(), len(data)))
	}

	block, err := UnmarshallBlock(bytes.NewBuffer(data), true)
	if err != nil {
		return nil, errors.Wrap(utils.WithBan(types.ErrUnmarshallBlock), err.Error())
	}
	block.BinData = data

	if err := block.readPreviousBlockFromBlockchainTable(); err != nil {
		return nil, err
	}

	return block, nil
}

// ProcessBlockWherePrevFromBlockchainTable is processing block with in table previous block
func ProcessBlockWherePrevFromBlockchainTable(data []byte, checkSize bool) (*Block, error) {
	if checkSize && int64(len(data)) > syspar.GetMaxBlockSize() {
		log.WithFields(log.Fields{"check_size": checkSize, "size": len(data), "max_size": syspar.GetMaxBlockSize(), "type": consts.ParameterExceeded}).Error("binary block size exceeds max block size")
		return nil, utils.WithBan(types.ErrMaxBlockSize(syspar.GetMaxBlockSize(), len(data)))
	}

	block, err := UnmarshallBlock(bytes.NewBuffer(data), true)
	if err != nil {
		return nil, errors.Wrap(utils.WithBan(types.ErrUnmarshallBlock), err.Error())
	}
	block.BinData = data

	if err := block.readPreviousBlockFromBlockchainTable(); err != nil {
		return nil, err
	}

	return block, nil
}
