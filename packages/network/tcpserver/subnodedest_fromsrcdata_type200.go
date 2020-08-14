/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package tcpserver

import (
	"encoding/base64"
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

func Type200(r *network.SubNodeSrcDataRequest) (*network.SubNodeSrcDataResponse, error) {
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
	resp := &network.SubNodeSrcDataResponse{}
	resp.Hash = hash

	//
	NodePrivateKey, NodePublicKey := utils.GetNodeKeys()
	if len(NodePrivateKey) < 1 {
		log.WithFields(log.Fields{"type": consts.EmptyObject}).Error("node private key is empty")
		err = errors.New(`empty node private key`)
		return nil, err
	}
		TaskUUID:           r.TaskUUID,
		DataUUID:           r.DataUUID,
		AgentMode:          AgentMode,
		TranMode:           TranMode,
		Hash:               hash,
		DataInfo:           r.DataInfo,
		SubNodeSrcPubkey:   r.SubNodeSrcPubkey,
		SubNodeAgentPubkey: r.SubNodeAgentPubkey,
		SubNodeAgentIP:     r.SubNodeAgentIp,
		SubNodeDestPubkey:  r.SubNodeDestPubkey,
		SubNodeDestIP:      r.SubNodeDestIp,
		//Data:         r.Data,
		//Data:         data,
		Data:       []byte(encodeDataString),
		CreateTime: time.Now().Unix(),
	}

	err = SubNodeDestData.Create()
	if err != nil {
		log.WithError(err)
		return nil, err
	}

	return resp, nil
}
