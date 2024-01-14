package upd

import (
	"fmt"
	"net"
)

type UDPClient struct {
	conn *net.UDPConn
}

func NewUDPClient(address string) (*UDPClient, error) {
	udpAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return nil, err
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return nil, err
	}

	client := &UDPClient{
		conn: conn,
	}

	return client, nil
}

func (c *UDPClient) SendMessage(message string) {
	_, err := c.conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending message via UDP:", err)
		return
	}
}

func (c *UDPClient) ReadResponse() string {
	buffer := make([]byte, 2048)
	n, _, err := c.conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Printf("[client %s log] Error reading: %s\n", c.conn.LocalAddr().String(), err.Error())
		return ""
	}

	response := string(buffer[:n])
	fmt.Printf("[client %s log] Received from serverside: %s\n", c.conn.LocalAddr().String(), response)
	return response
}

func (c *UDPClient) Close() {
	c.conn.Close()
}
