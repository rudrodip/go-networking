package upd

import (
	"fmt"
	"net"
)

type UDPServer struct {
	conn *net.UDPConn
}

func NewUDPServer(address string) (*UDPServer, error) {
	udpAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return nil, err
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return nil, err
	}

	server := &UDPServer{
		conn: conn,
	}

	return server, nil
}

func (s *UDPServer) HandleClient() {
	buffer := make([]byte, 2048)

	for {
		n, addr, err := s.conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("[server log] Error reading from UDP:", err)
			return
		}

		clientData := buffer[:n]
		fmt.Printf("Received data from %s: %s\n", addr.String(), string(clientData))

		response := []byte(clientData)
		_, err = s.conn.WriteToUDP(response, addr)
		if err != nil {
			fmt.Println("[server log] Error writing to UDP:", err)
			return
		}
	}
}

func (s *UDPServer) Start() {
	fmt.Println("UDP server started at", s.conn.LocalAddr().String())
	s.HandleClient()
}
