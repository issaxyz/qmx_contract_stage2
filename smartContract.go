/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	iot "github.com/issaxyz/qmx_contract_stage2/smart-contract"
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
