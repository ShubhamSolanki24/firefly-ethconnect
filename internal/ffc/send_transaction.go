// Copyright © 2022 Kaleido, Inc.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ffc

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"strings"

	"github.com/hyperledger/firefly-ethconnect/internal/errors"
	"github.com/hyperledger/firefly-ethconnect/internal/eth"
	"github.com/hyperledger/firefly-transaction-manager/pkg/ffcapi"
)

func (s *ffcServer) sendTransaction(ctx context.Context, payload []byte) (interface{}, ffcapi.ErrorReason, error) {

	var req ffcapi.SendTransactionRequest
	err := json.Unmarshal(payload, &req)
	if err != nil {
		return nil, ffcapi.ErrorReasonInvalidInputs, err
	}

	nonce := json.Number(req.Nonce.Int().String())
	gas := json.Number(req.Gas.Int().String())
	var gasPrice json.Number
	if req.GasPrice != nil {
		err := json.Unmarshal([]byte(*req.GasPrice), &gasPrice)
		if err != nil {
			return nil, ffcapi.ErrorReasonInvalidInputs, errors.Errorf(errors.FFCInvalidGasPrice, string(*req.GasPrice), err)
		}
	}
	var value json.Number
	if req.Value != nil {
		value = json.Number(req.Value.Int().String())
	}
	txData, err := hex.DecodeString(strings.TrimPrefix(req.TransactionData, "0x"))
	if err != nil {
		return nil, ffcapi.ErrorReasonInvalidInputs, errors.Errorf(errors.FFCInvalidTXData, req.TransactionData, err)
	}
	tx, err := eth.NewRawSendTxn(nil, req.From, req.To, nonce, value, gas, gasPrice, txData)
	if err != nil {
		return nil, ffcapi.ErrorReasonInvalidInputs, err
	}

	err = tx.Send(ctx, s.rpc)
	if err != nil {
		return nil, s.mapSendError(err.Error()), err
	}
	return &ffcapi.SendTransactionResponse{
		TransactionHash: tx.Hash,
	}, "", nil

}

func (s *ffcServer) mapSendError(errString string) ffcapi.ErrorReason {
	switch {
	case strings.Contains(errString, "nonce too low"):
		return ffcapi.ErrorReasonNonceTooLow
	case strings.Contains(errString, "insufficient funds"):
		return ffcapi.ErrorReasonInsufficientFunds
	case strings.Contains(errString, "transaction underpriced"):
		return ffcapi.ErrorReasonTransactionUnderpriced
	default:
		return ""
	}
}
