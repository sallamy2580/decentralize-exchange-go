/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package modes

import (
	"errors"

	"github.com/IBAX-io/go-ibax/packages/transaction"

	"github.com/IBAX-io/go-ibax/packages/conf"
	"github.com/IBAX-io/go-ibax/packages/consts"
	"github.com/IBAX-io/go-ibax/packages/model"
	"github.com/IBAX-io/go-ibax/packages/types"
	log "github.com/sirupsen/logrus"
)

var ErrDiffKey = errors.New("Different keys")

type blockchainTxPreprocessor struct{}

func (p blockchainTxPreprocessor) ProcessClientTxBatches(txDatas [][]byte, key int64, le *log.Entry) (retTx []string, err error) {
	var rtxs []*model.RawTx
	for _, txData := range txDatas {
		rtx := &transaction.Transaction{}
		if err = rtx.Processing(txData); err != nil {
			return nil, err
		}
		rtxs = append(rtxs, rtx.SetRawTx())
		retTx = append(retTx, rtx.HashStr())
	}
	err = model.SendTxBatches(rtxs)
	return
}

type ObsTxPreprocessor struct{}

/*
func (p ObsTxPreprocessor) ProcessClientTranstaction(txData []byte, key int64, le *log.Entry) (string, error) {

	tx, err := transaction.UnmarshallTransaction(bytes.NewBuffer(txData), true)
	if err != nil {
		le.WithFields(log.Fields{"type": consts.ParseError, "error": err}).Error("on unmarshaling user tx")
		return "", err
	}

	ts := &model.TransactionStatus{
		BlockID:  1,
		Hash:     tx.TxHash,
		Time:     time.Now().Unix(),
		WalletID: key,
		Type:     tx.Rtx.Type(),
	}

	if err := ts.Create(); err != nil {
		le.WithFields(log.Fields{"type": consts.DBError, "error": err}).Error("on creating tx status")
		return "", err
	}

	res, _, err := tx.CallOBSContract()
	if err != nil {
		le.WithFields(log.Fields{"type": consts.ParseError, "error": err}).Error("on execution contract")
		return "", err
	}

	if err := model.SetTransactionStatusBlockMsg(nil, 1, res, tx.TxHash); err != nil {
		le.WithFields(log.Fields{"type": consts.DBError, "error": err, "tx_hash": tx.TxHash}).Error("updating transaction status block id")
		return "", err
	}

	return string(converter.BinToHex(tx.TxHash)), nil
}*/

func (p ObsTxPreprocessor) ProcessClientTxBatches(txData [][]byte, key int64, le *log.Entry) ([]string, error) {
	return nil, nil
}

func GetClientTxPreprocessor() types.ClientTxPreprocessor {
	if conf.Config.IsSupportingOBS() {
		return ObsTxPreprocessor{}
	}

	return blockchainTxPreprocessor{}
}

// BlockchainSCRunner implementls SmartContractRunner for blockchain mode
type BlockchainSCRunner struct{}

// RunContract runs smart contract on blockchain mode
func (runner BlockchainSCRunner) RunContract(data, hash []byte, keyID, tnow int64, le *log.Entry) error {
	if err := transaction.CreateTransaction(data, hash, keyID, tnow); err != nil {
		le.WithFields(log.Fields{"type": consts.ContractError, "error": err}).Error("Executing contract")
		return err
	}

	return nil
}

// OBSSCRunner implementls SmartContractRunner for obs mode
type OBSSCRunner struct{}

// RunContract runs smart contract on obs mode
func (runner OBSSCRunner) RunContract(data, hash []byte, keyID, tnow int64, le *log.Entry) error {
	proc := GetClientTxPreprocessor()
	_, err := proc.ProcessClientTxBatches([][]byte{data}, keyID, le)
	if err != nil {
		le.WithFields(log.Fields{"error": consts.ContractError}).Error("on run internal NewUser")
		return err
	}

	return nil
}

// GetSmartContractRunner returns mode boundede implementation of SmartContractRunner
func GetSmartContractRunner() types.SmartContractRunner {
	if !conf.Config.IsSupportingOBS() {
		return BlockchainSCRunner{}
	}

	return OBSSCRunner{}
}
