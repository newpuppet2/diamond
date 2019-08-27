package main

/**
 * Shows how to use the history
 **/

import (
	// For printing messages on console
	"fmt"

	// The shim package
	"github.com/hyperledger/fabric/core/chaincode/shim"

	// // peer.Response is in the peer package
	"github.com/hyperledger/fabric/protos/peer"

	// JSON Encoding
	"encoding/json"

	// KV Interface

	"strconv"
)

type DiamondChaincode struct {
}

type Diamondore struct {
	DocType			string  `json:"docType"`
	Uniqmid			string  `json:"uniqmid"`
    Oreid           string  `json:"oreid"`
	Caratw			int    `json:"caratw"`
	Clarity			int    `json:"clarity"`
	Location	    string  `json:"location"`
	Deducted        int    `json:"deducted"`
}

type Diamondcut struct {
	DocType			string  `json:"docType"`
	Uniqid			string  `json:"uniqid"`
	Oreid           string  `json:"oreid"`
	Caratw			int    `json:"caratw"`
	Clarity			int    `json:"clarity"`
	Color           string  `json:"color"`
	Cut             string  `json:"cut"`
	Shape           string  `json:"shape"`
	Certified       string  `json:"certified"`
    Owner           string  `json:"owner"`	
}

const	DocType	= "DiamondAsset"

func (diamond *DiamondChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	// Simply print a message
	fmt.Println("Init executed in DiamondAssetRegistory")

	// Return success
	return shim.Success(nil)
}

func (diamond *DiamondChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	// Get the function name and parameters
	funcName, args := stub.GetFunctionAndParameters()

	if (funcName == "CreateOre") {
		return  CreateOre(stub, args)
     } else if (funcName == "CreateCut") {
	       return CreateCut(stub, args)
	}
	
	
	 return shim.Error(("Bad Function Name = !!!"))
}

func CreateOre(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	
	uniqmid1 := string(args[0])
	Oreid1 := string(args[1])
	Caratw1, _ := strconv.Atoi(string(args[2]))
	Clarity1, _ := strconv.Atoi(string(args[3]))
	Location1 := string(args[4])
	
	diamondore := Diamondore{DocType: DocType, Uniqmid: uniqmid1, Oreid: Oreid1, Caratw: Caratw1, Clarity: Clarity1, Location: Location1, Deducted: Caratw1}
	jsondiamond, _ := json.Marshal(diamondore)
	// Key = VIN#, Value = Car's JSON representation
	stub.PutState(Oreid1, jsondiamond)
        return shim.Success([]byte("Success"))
}

func CreateCut(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	
	uniqid1 := string(args[0])
	Oreid1 := string(args[1])
	Caratw1, _ := strconv.Atoi(string(args[2]))
	Clarity1, _ := strconv.Atoi(string(args[3]))
	Color1 := string(args[4])
	Cut1 := string(args[5])
	Shape1 := string(args[6])
	Certified1 := string(args[7])
	Owner1 := string(args[8])
	
	diamondcut := Diamondcut{DocType: DocType, Uniqid: uniqid1, Oreid: Oreid1, Caratw: Caratw1, Clarity: Clarity1, Color: Color1, Cut: Cut1, Shape: Shape1, Certified: Certified1, Owner: Owner1}
	jsondiamond, _ := json.Marshal(diamondcut)
	stub.PutState(uniqid1, jsondiamond)

	var diamond1 Diamondore
	_ = json.Unmarshal([]byte(Oreid1), &diamond1)

	diamond1.Deducted = diamond1.Deducted - Caratw1
	jsondiamond, _ = json.Marshal(diamond1)

	stub.PutState(Oreid1, jsondiamond)
return shim.Success([]byte("Success"))
}

func main() {
	fmt.Printf("Started Chaincode. Diamond-Tracking\n")
	err := shim.Start(new(DiamondChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}







