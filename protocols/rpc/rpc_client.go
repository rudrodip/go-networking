package rpc

import (
	"fmt"
	"net/rpc"
)

func CallAddProcedure(message Args) (Calculator, error) {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer client.Close()

	var addResult int
	err = client.Call("Calculator.Add", message, &addResult)
	if err != nil {
		fmt.Println("Error calling Add method:", err)
		return 0, err
	}
	fmt.Printf("Add result: %d\n", addResult)
	return Calculator(addResult), nil
}

func CallMultiplyProcedure(message Args) (Calculator, error) {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer client.Close()

	var multiplyResult int
	err = client.Call("Calculator.Multiply", message, &multiplyResult)
	if err != nil {
		fmt.Println("Error calling Multiply method:", err)
		return 0, err
	}
	fmt.Printf("Multiply result: %d\n", multiplyResult)
	return Calculator(multiplyResult), nil
}
