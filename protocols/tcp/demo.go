package tcp

import (
	"fmt"
	"time"

	"github.com/rudrodip/go-networking/pkg"
)

func TestTCP() {
	tcpServer, err := NewTCPServer("localhost", 8080)
	if err != nil {
		fmt.Printf("Error creating server: %v", err)
		return
	}

	tcpClient, err := NewTCPClient("localhost", 8080)
	if err != nil {
		fmt.Printf("Error creating client: %v", err)
		return
	}

	go func() {
		defer tcpServer.listener.Close()
		tcpServer.Start()
	}()

	time.Sleep(100 * time.Millisecond)

	for {
		message := pkg.GetUserInput("[enter \"exit\" to exit]Enter message: ")
		if message == "exit" {
			break
		}
		tcpClient.SendMessage(message)
		tcpClient.ReadResponse()
	}

	defer tcpServer.listener.Close()
	time.Sleep(100 * time.Millisecond)
	defer tcpClient.conn.Close()
	time.Sleep(100 * time.Millisecond)
}
