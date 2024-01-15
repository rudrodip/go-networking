package rpc

import (
	"fmt"
	"net"
	"net/rpc"
)

type Calculator int

type Args struct {
	A, B int
}

func (c *Calculator) Add(args Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}

func (c *Calculator) Multiply(args Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func Start() {
	calculator := new(Calculator)
	rpc.Register(calculator)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server listening on port 1234")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
