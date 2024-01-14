package upd

import (
	"fmt"
	"time"

	"github.com/rudrodip/go-networking/pkg"
)

func TestUDP() {
	udpServer, err := NewUDPServer("localhost:8080")
	if err != nil {
		fmt.Printf("Error creating server: %v", err)
		return
	}

	udpClient, err := NewUDPClient("localhost:8080")
	if err != nil {
		fmt.Printf("Error creating client: %v", err)
		return
	}

	go func() {
		defer udpServer.conn.Close()
		udpServer.Start()
	}()

	time.Sleep(100 * time.Millisecond)

	for {
		message := pkg.GetUserInput("[enter \"exit\" to exit]Enter message: ")
		if message == "exit" {
			break
		}
		udpClient.SendMessage(message)
		udpClient.ReadResponse()
	}

	defer udpServer.conn.Close()
	time.Sleep(100 * time.Millisecond)
	defer udpClient.conn.Close()
	time.Sleep(100 * time.Millisecond)
}
