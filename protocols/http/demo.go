package http

import (
	"fmt"
	"time"

	"github.com/rudrodip/go-networking/pkg"
)

func TestHttp() {
	server, err := NewHttpServer("localhost:8080")
	if err != nil {
		fmt.Printf("Error creating server: %v", err)
		return
	}

	client, err := NewHttpClient()
	if err != nil {
		fmt.Printf("Error creating client: %v", err)
		return
	}

	go func() {
		defer server.Stop()
		server.Start()
	}()

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Available routes:")
	routes := GetRoutes()
	for i, route := range routes {
		fmt.Println(i+1, route)
	}
	fmt.Println(len(routes)+1, "/[dynamic]")

	for {
		route := pkg.GetUserInput("[enter \"exit\" to exit]Enter route: ")
		if route == "exit" {
			server.Stop()
			client.Stop()
			break
		}

		client.Get("http://localhost:8080" + route)
	}

	client.Stop()
}
