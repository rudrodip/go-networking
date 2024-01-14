package tcp

import (
	"fmt"
	"net"
)

type TCPClient struct {
	conn net.Conn
}

func NewTCPClient(serverAddress string, port int) (*TCPClient, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverAddress, port))
	if err != nil {
		return nil, err
	}
	return &TCPClient{conn: conn}, nil
}

func (c *TCPClient) SendMessage(message string) {
	c.conn.Write([]byte(message))
}

func (c *TCPClient) ReadResponse() string {
	buffer := make([]byte, 2048)
	n, err := c.conn.Read(buffer)
	if err != nil {
		fmt.Printf("[client %s log] Error reading: %s\n", c.conn.LocalAddr().String(), err.Error())
		return ""
	}

	response := string(buffer[:n])
	fmt.Printf("[client %s log] Received from serverside: %s\n", c.conn.LocalAddr().String(), response)
	return response
}

func (c *TCPClient) Close() {
	c.conn.Close()
}
