/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	iot "github.com/hyperledger/fabric-samples/asset-transfer-abac/chaincode-go/smart-contract"
)

func main() {
	iotSmartContract, err := contractapi.NewChaincode(&iot.SmartContract{})
	if err != nil {
		log.Panicf("Error creating iot chaincode: %v", err)
	}

	if err := iotSmartContract.Start(); err != nil {
		log.Panicf("Error starting iot chaincode: %v", err)
	}
}
