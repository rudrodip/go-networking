package upd

import (
	"sync"
	"testing"
	"time"
)

func TestNewUDPServer(t *testing.T) {
	server, err := NewUDPServer("localhost:8080")
	if err != nil {
		t.Errorf("Error creating server: %v", err)
		return
	}

	client, err := NewUDPClient("localhost:8080")
	if err != nil {
		t.Errorf("Error creating client: %v", err)
		return
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
	server.conn.Close()
	wg.Wait()
}

func assertEqual(t *testing.T, actual, expected, message string) {
	t.Helper()
	if actual != expected {
		t.Errorf("%s: expected '%s', got '%s'", message, expected, actual)
	}
}
