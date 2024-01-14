package tcp

import (
	"fmt"
	"io"
	"net"
)

type TCPServer struct {
	listener net.Listener
}

func NewTCPServer(address string, port int) (*TCPServer, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		return nil, err
	}
	return &TCPServer{listener: listener}, nil
}

func (s *TCPServer) HandleClient(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 2048)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("[server log] Client closed the connection.")
			} else {
				fmt.Println("[server log] Error reading:", err.Error())
			}
			return
		}

		fmt.Println("[server log] Received from client: ", string(buffer[:n]))

		conn.Write(buffer[:n])
	}
}

func (s *TCPServer) Start() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println("[server log] Error accepting:", err.Error())
			return
		}
		fmt.Println("[server log] Accepted connection from:", conn.RemoteAddr().String())

		go s.HandleClient(conn)
	}
}
