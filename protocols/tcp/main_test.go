package tcp

import (
	"sync"
	"testing"
	"time"
)

func TestMainFunctionality(t *testing.T) {
	server, err := NewTCPServer("localhost", 8080)
	if err != nil {
		t.Fatalf("Error creating server: %v", err)
	}

	client, err := NewTCPClient("localhost", 8080)
	if err != nil {
		t.Fatalf("Error creating client: %v", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		server.Start()
	}()

	time.Sleep(100 * time.Millisecond)

	client.SendMessage("message 1")
	response1 := client.ReadResponse()
	assertEqual(t, response1, response1, "Unexpected response from server")

	client.SendMessage("message 2")
	response2 := client.ReadResponse()
	assertEqual(t, response2, response2, "Unexpected response from server")

	client.Close()

	server.listener.Close()

	wg.Wait()
}

func assertEqual(t *testing.T, actual, expected, message string) {
	t.Helper()
	if actual != expected {
		t.Errorf("%s: expected '%s', got '%s'", message, expected, actual)
	}
}
