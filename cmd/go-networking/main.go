package main

import (
	"fmt"

	"github.com/rudrodip/go-networking/pkg"
	http "github.com/rudrodip/go-networking/protocols/http"
	tcp "github.com/rudrodip/go-networking/protocols/tcp"
	udp "github.com/rudrodip/go-networking/protocols/udp"
)

var protocolMap = map[string]func(){
	"tcp":  tcp.TestTCP,
	"udp":  udp.TestUDP,
	"http": http.TestHttp,
}

func main() {
	for {
		fmt.Println("Available protocols:")
		for protocol := range protocolMap {
			fmt.Printf(" -> %s\n", protocol)
		}
		protocol := pkg.GetUserInput("[enter \"exit\" to exit]Enter protocol: ")
		if protocol == "exit" {
			break
		}
		if test, ok := protocolMap[protocol]; ok {
			test()
		} else {
			fmt.Println("Invalid protocol")
		}
	}
}
