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
				BenhNhan{Mabn:"2020083198",Hoten:"Nguyễn Thị Mỹ Hương",Ngaysinh:"01/01/1969",Gioitinh:"Nữ",Cmnd:"",Diachi:"Phường An Thới, Quận Bình Thuỷ, Cần Thơ",Maxa:"9291831177"},
				BenhNhan{Mabn:"2020083191",Hoten:"Nguyễn Thị Sáu",Ngaysinh:"01/01/1939",Gioitinh:"Nữ",Cmnd:"",Diachi:"Phú Thạnh, Xã Long Phú, Huyện Tam Bình, Tỉnh Vĩnh Long",Maxa:"8686000000"},
				BenhNhan{Mabn:"2020083154",Hoten:"Đặng Thị Tám",Ngaysinh:"01/01/1951",Gioitinh:"Nữ",Cmnd:"",Diachi:"ấp thới hòa c, Xã Thới Xuân, Huyện Cờ Đỏ, Thành phố Cần Thơ",Maxa:"9292531277"},
				BenhNhan{Mabn:"2020083148",Hoten:"Nguyễn Thị Hồng Hoa",Ngaysinh:"28/08/1955",Gioitinh:"Nữ",Cmnd:"",Diachi:"80/28 PNL, Phường An Hòa, Quận Ninh Kiều, Thành phố Cần Thơ",Maxa:"9291631120"},
				BenhNhan{Mabn:"2020083138",Hoten:"Nguyễn Văn Dũng",Ngaysinh:"01/01/1957",Gioitinh:"Nam",Cmnd:"",Diachi:"ấp Phú Sơn C, Xã Long Phú, Huyện Tam Bình, Vĩnh Long",Maxa:"8686029752"},
				BenhNhan{Mabn:"2020083137",Hoten:"Nguyễn Quỳnh",Ngaysinh:"01/01/1940",Gioitinh:"Nam",Cmnd:"",Diachi:"KV Tân Lợi 2, Phường Tân Hưng, Quận Thốt Nốt, Thành phố Cần Thơ",Maxa:"9292331227"},
				BenhNhan{Mabn:"2020083131",Hoten:"Trần Hùng Sơn",Ngaysinh:"01/01/1968",Gioitinh:"Nam",Cmnd:"",Diachi:"Ấp Thạnh Lộc 2, Xã Trung An, Huyện Cờ Đỏ, Thành phố Cần Thơ",Maxa:"9292531222"},
				BenhNhan{Mabn:"2020083105",Hoten:"Nguyễn Thị Tư",Ngaysinh:"01/01/1942",Gioitinh:"Nữ",Cmnd:"",Diachi:"172/5 - Nguyễn Việt Hồng, Phường An Phú, Quận Ninh Kiều, Thành phố Cần Thơ",Maxa:"9291600000"},
				BenhNhan{Mabn:"2020083104",Hoten:"Nguyễn Thị Điếu",Ngaysinh:"01/01/1941",Gioitinh:"Nữ",Cmnd:"",Diachi:"Kv Thới Đông, Phường Phước Thới, Quận Ô Môn, Thành phố Cần Thơ",Maxa:"9291731162"},
				BenhNhan{Mabn:"2020083058",Hoten:"Đỗ Thị Gạo",Ngaysinh:"18/01/1933",Gioitinh:"Nữ",Cmnd:"",Diachi:"43/18/15 PNL, Phường Thới Bình, Quận Ninh Kiều, Thành phố Cần Thơ",Maxa:"9291600000"},
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
