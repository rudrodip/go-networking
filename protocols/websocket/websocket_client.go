package websocket

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	url        string
	connection *websocket.Conn
	done       chan struct{}
	interrupt  chan os.Signal
}

func NewClient(url string) *Client {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	return &Client{
		url:       url,
		done:      make(chan struct{}),
		interrupt: interrupt,
	}
}

func (c *Client) Connect() error {
	conn, _, err := websocket.DefaultDialer.Dial(c.url, nil)
	if err != nil {
		return err
	}
	c.connection = conn
	return nil
}

func (c *Client) StartReceivingMessages() {
	go func() {
		defer close(c.done)
		for {
			_, message, err := c.connection.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			fmt.Printf("Received message: %s\n", message)
		}
	}()
}

func (c *Client) SendMessage(message string) error {
	err := c.connection.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Close() {
	fmt.Println("Closing the WebSocket client...")
	closeMessage := websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")
	_ = c.connection.WriteMessage(websocket.CloseMessage, closeMessage)
	select {
	case <-c.done:
	case <-time.After(time.Second):
	}
}
