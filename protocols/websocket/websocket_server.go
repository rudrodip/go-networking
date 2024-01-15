package websocket

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type Server struct {
	port     int
	server   *http.Server
	upgrader *websocket.Upgrader
}

func NewServer(port int) *Server {
	upgrader := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	return &Server{
		port:     port,
		upgrader: upgrader,
	}
}

func (s *Server) Start() {
	http.HandleFunc("/ws", s.handleWebSocket)
	address := fmt.Sprintf(":%d", s.port)
	s.server = &http.Server{Addr: address}
	fmt.Printf("WebSocket server started on ws://localhost%s\n", address)
	err := s.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func (s *Server) Close() {
	fmt.Println("Shutting down the WebSocket server...")
	if err := s.server.Close(); err != nil {
		fmt.Printf("Error shutting down server: %s\n", err)
	}
}

func (s *Server) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Error upgrading to WebSocket: %s\n", err)
		return
	}
	defer conn.Close()

	fmt.Println("WebSocket connection established.")

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("Error reading message: %s\n", err)
			return
		}
		fmt.Printf("Message type: %d\n", messageType)
		fmt.Printf("Message: %s\n", p)
	}
}
