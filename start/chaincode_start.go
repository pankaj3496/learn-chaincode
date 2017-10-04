package main

import ("errors"
		  "fmt"
		 "github.com/hyperledger/fabric/core/chaincode/shim")
type SimpleChaincode struct {}


func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte,error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect no. of arguments, Expecting 1")
	}
err :=stud.PutState("hello_world",[]byte(args[0]))
if err != nil{
	return nil,err
}
 return nil,nil
}
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte,error){
	fmt.Println("invoke is running" +function)

	if function =="init" {
		return t.Init(stub,"init",args)
	}else if function == "write" {
		return t.write(stub,args)
	}
	fmt.Println("invoke did not find func:" +function)
	return nil,errors.New("Received Unknown Function Invocation: "+function)
}
func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface,args []string) ([]byte,error) {
	var key, value string
	var err error
	fmt.Println("running write()")
	if len(args) !=2 {
		return nil,errors.New("Incorrect No. Of Arguments. Expecting 2. Name of key and value to set")
	}
key =args[0]
value = args[1]
err =  stub.PutState(key,[]byte(value))
if err != nil{
	return nil, err
}
return nil,nil
} 
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface,function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)
	
		if function == "read" {                            
			return t.read(stub, args)
		}
		fmt.Println("query did not find func: " + function)
	
	return nil, errors.New("Received unknown function query: " + function)
}

func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error

    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
    }

    key = args[0]
    valAsbytes, err := stub.GetState(key)
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
        return nil, errors.New(jsonResp)
    }
	return valAsbytes, nil
}

func main() {
    err := shim.Start(new(SimpleChaincode))
    if err != nil {
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}
