/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */

package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the bn structure, with 4 properties.  Structure tags are used by encoding/json library
type BenhNhan struct {
	Mabn     string `json:"mabn"`
	Hoten    string `json:"hoten"`
	Ngaysinh string `json:"ngaysinh"`
	Gioitinh string `json:"gioitinh"`
	Cmnd     string `json:"cmnd"`
	Diachi   string `json:"diachi"`
	Maxa     string `json:"maxa"`
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryBN" {
		return s.queryBN(APIstub, args)

	} else if function == "initLedger" {
		return s.initLedger(APIstub)

	} else if function == "createBN" {
		return s.createBN(APIstub, args)

	} else if function == "queryAllBN" {
		return s.queryAllBN(APIstub)

	} else if function == "changeBN" {
		return s.changeBN(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) queryBN(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	bnAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(bnAsBytes)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	/*
		bns := []BenhNhan{
				BenhNhan{Mabn:"2020083198",Hoten:"Nguy???n Th??? M??? H????ng",Ngaysinh:"01/01/1969",Gioitinh:"N???",Cmnd:"",Diachi:"Ph?????ng An Th???i, Qu???n B??nh Thu???, C???n Th??",Maxa:"9291831177"},
				BenhNhan{Mabn:"2020083191",Hoten:"Nguy???n Th??? S??u",Ngaysinh:"01/01/1939",Gioitinh:"N???",Cmnd:"",Diachi:"Ph?? Th???nh, X?? Long Ph??, Huy???n Tam B??nh, T???nh V??nh Long",Maxa:"8686000000"},
				BenhNhan{Mabn:"2020083154",Hoten:"?????ng Th??? T??m",Ngaysinh:"01/01/1951",Gioitinh:"N???",Cmnd:"",Diachi:"???p th???i h??a c, X?? Th???i Xu??n, Huy???n C??? ?????, Th??nh ph??? C???n Th??",Maxa:"9292531277"},
				BenhNhan{Mabn:"2020083148",Hoten:"Nguy???n Th??? H???ng Hoa",Ngaysinh:"28/08/1955",Gioitinh:"N???",Cmnd:"",Diachi:"80/28 PNL, Ph?????ng An H??a, Qu???n Ninh Ki???u, Th??nh ph??? C???n Th??",Maxa:"9291631120"},
				BenhNhan{Mabn:"2020083138",Hoten:"Nguy???n V??n D??ng",Ngaysinh:"01/01/1957",Gioitinh:"Nam",Cmnd:"",Diachi:"???p Ph?? S??n C, X?? Long Ph??, Huy???n Tam B??nh, V??nh Long",Maxa:"8686029752"},
				BenhNhan{Mabn:"2020083137",Hoten:"Nguy???n Qu???nh",Ngaysinh:"01/01/1940",Gioitinh:"Nam",Cmnd:"",Diachi:"KV T??n L???i 2, Ph?????ng T??n H??ng, Qu???n Th???t N???t, Th??nh ph??? C???n Th??",Maxa:"9292331227"},
				BenhNhan{Mabn:"2020083131",Hoten:"Tr???n H??ng S??n",Ngaysinh:"01/01/1968",Gioitinh:"Nam",Cmnd:"",Diachi:"???p Th???nh L???c 2, X?? Trung An, Huy???n C??? ?????, Th??nh ph??? C???n Th??",Maxa:"9292531222"},
				BenhNhan{Mabn:"2020083105",Hoten:"Nguy???n Th??? T??",Ngaysinh:"01/01/1942",Gioitinh:"N???",Cmnd:"",Diachi:"172/5 - Nguy???n Vi???t H???ng, Ph?????ng An Ph??, Qu???n Ninh Ki???u, Th??nh ph??? C???n Th??",Maxa:"9291600000"},
				BenhNhan{Mabn:"2020083104",Hoten:"Nguy???n Th??? ??i???u",Ngaysinh:"01/01/1941",Gioitinh:"N???",Cmnd:"",Diachi:"Kv Th???i ????ng, Ph?????ng Ph?????c Th???i, Qu???n ?? M??n, Th??nh ph??? C???n Th??",Maxa:"9291731162"},
				BenhNhan{Mabn:"2020083058",Hoten:"????? Th??? G???o",Ngaysinh:"18/01/1933",Gioitinh:"N???",Cmnd:"",Diachi:"43/18/15 PNL, Ph?????ng Th???i B??nh, Qu???n Ninh Ki???u, Th??nh ph??? C???n Th??",Maxa:"9291600000"},
		}

		i := 0
		for i < len(bns) {
			fmt.Println("i is ", i)
			bnAsBytes, _ := json.Marshal(bns[i])
			APIstub.PutState("BN"+strconv.Itoa(i), bnAsBytes)
			fmt.Println("Added", bns[i])
			i = i + 1
		}
	*/
	return shim.Success(nil)
}

func (s *SmartContract) createBN(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 8")
	}

	var bn = BenhNhan{Mabn: args[1], Hoten: args[2], Ngaysinh: args[3], Gioitinh: args[4], Cmnd: args[5], Diachi: args[6], Maxa: args[7]}

	bnAsBytes, _ := json.Marshal(bn)
	APIstub.PutState(args[0], bnAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryAllBN(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "BN0"
	endKey := "BN999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllBN:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) changeBN(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	bnAsBytes, _ := APIstub.GetState(args[0])
	bn := BenhNhan{}

	json.Unmarshal(bnAsBytes, &bn)
	bn.Hoten = args[1]

	bnAsBytes, _ = json.Marshal(bn)
	APIstub.PutState(args[0], bnAsBytes)

	return shim.Success(nil)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
