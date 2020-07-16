/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package tcpserver

import (
	"errors"
	"fmt"
	"time"

	"github.com/IBAX-io/go-ibax/packages/consts"
	"github.com/IBAX-io/go-ibax/packages/converter"
	"github.com/IBAX-io/go-ibax/packages/crypto"
	"github.com/IBAX-io/go-ibax/packages/crypto/ecies"
	"github.com/IBAX-io/go-ibax/packages/model"
	"github.com/IBAX-io/go-ibax/packages/network"
	"github.com/IBAX-io/go-ibax/packages/utils"

	log "github.com/sirupsen/logrus"
)

func Type102(r *network.VDEAgentDataRequest) (*network.VDEAgentDataResponse, error) {
	nodePrivateKey, err := utils.GetNodePrivateKey()
	if err != nil || len(nodePrivateKey) < 1 {
		if err == nil {
			log.WithFields(log.Fields{"type": consts.EmptyObject}).Error("node private key is empty")
			return nil, errors.New("Incorrect private key length")
		}
		return nil, err
	}

	data, err := ecies.EccDeCrypto(r.Data, nodePrivateKey)
	if err != nil {
		fmt.Println("EccDeCrypto err!")
		log.WithError(err)
		return nil, err
	}

	//hash, err := crypto.HashHex(r.Data)
	hash, err := crypto.HashHex(data)
	if err != nil {
		log.WithError(err)
		return nil, err
	}
	resp := &network.VDEAgentDataResponse{}
	resp.Hash = hash
	AgentMode := converter.StrToInt64(r.AgentMode)
	VDEDestData := model.VDEDestData{
		TaskUUID:       r.TaskUUID,
		DataUUID:       r.DataUUID,
		AgentMode:      AgentMode,
		Hash:           hash,
		DataInfo:       r.DataInfo,
		VDESrcPubkey:   r.VDESrcPubkey,
		VDEDestPubkey:  r.VDEDestPubkey,
		VDEDestIp:      r.VDEDestIp,
		//Data:         r.Data,
		Data:       data,
		CreateTime: time.Now().Unix(),
	}

	err = VDEDestData.Create()
	if err != nil {
		log.WithError(err)
		return nil, err
	}

	return resp, nil
}
