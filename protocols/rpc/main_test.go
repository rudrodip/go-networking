package rpc

import (
	"testing"
	"time"
)

func TestMainFunc(t *testing.T) {
	// Start RPC server in a goroutine
	go func() {
		Start()
	}()

	// Wait for the server to initialize
	time.Sleep(time.Second)

	// Make RPC calls
	agr1 := Args{3, 4}
	agr2 := Args{10, -9}

	result1, err := CallAddProcedure(agr1)
	assertNoError(t, err)
	assertEqual(t, result1, Calculator(7), "Unexpected response from server")

	result2, err := CallAddProcedure(agr2)
	assertNoError(t, err)
	assertEqual(t, result2, Calculator(1), "Unexpected response from server")

	result3, err := CallMultiplyProcedure(agr1)
	assertNoError(t, err)
	assertEqual(t, result3, Calculator(12), "Unexpected response from server")

	result4, err := CallMultiplyProcedure(agr2)
	assertNoError(t, err)
	assertEqual(t, result4, Calculator(-90), "Unexpected response from server")
}

func assertEqual(t *testing.T, actual Calculator, expected Calculator, message string) {
	t.Helper()
	if actual != expected {
		t.Errorf("%s: expected '%+v', got '%+v'", message, expected, actual)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}
