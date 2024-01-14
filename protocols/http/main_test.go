package http

import (
	"sync"
	"testing"
	"time"
)

func TestMainFunctionality(t *testing.T) {
	server, err := NewHttpServer("localhost:8080")
	if err != nil {
		t.Fatalf("Error creating server: %v", err)
	}

	client, err := NewHttpClient()
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

	assertEqual(t, client.Get("http://localhost:8080/res"), "res", "Unexpected response from server")
	assertEqual(t, client.Get("http://localhost:8080/res2"), "res2", "Unexpected response from server")

	server.Stop()
}

func assertEqual(t *testing.T, actual, expected, message string) {
	t.Helper()
	if actual != expected {
		t.Errorf("%s: expected '%s', got '%s'", message, expected, actual)
	}
}
