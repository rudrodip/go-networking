package tcp

import (
	"fmt"
	"net"
)

type TCPServer struct {
	listener net.Listener
}

func NewTCPServer() (*TCPServer, error) {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		return nil, err
	}

	return &TCPServer{listener: listener}, nil
}

func (s *TCPServer) Start() error {
	fmt.Println("Starting TCP server")
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection")
			return err
		}

		go s.handleConnection(conn)
	}
}

func (s *TCPServer) handleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	fmt.Printf("Received data: %s\n", buffer[:n])

	// Close the connection
	conn.Close()
}

func (s *TCPServer) Stop() error {
	return s.listener.Close()
}
