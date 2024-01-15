# Networking Concepts in Go

This repository provides examples and demonstrations of various networking concepts implemented in the Go programming language. Each section includes both server and client implementations, along with test and demo applications.

## Project Structure

```
.
├── cmd
│   └── go-networking
│       └── main.go
├── go.mod
├── go.sum
├── Makefile
├── pkg
│   └── userinput.go
└── protocols
    ├── http
    │   ├── demo.go
    │   ├── http_client.go
    │   ├── http_server.go
    │   └── main_test.go
    ├── rpc
    │   ├── main_test.go
    │   ├── rpc_client.go
    │   └── rpc_server.go
    ├── tcp
    │   ├── demo.go
    │   ├── main_test.go
    │   ├── tcp_client.go
    │   └── tcp_server.go
    ├── udp
    │   ├── demo.go
    │   ├── main_test.go
    │   ├── udp_server.go
    │   └── upd_client.go
    └── websocket
        ├── websocket_client.go
        └── websocket_server.go
```

## TODO Structure

- [x] **TCP**
    - [x] **TCP Server**
    - [x] **TCP Client**
    - [x] **TCP Test**
    - [x] **TCP Demo**

- [x] **UDP**
    - [x] **UDP Server**
    - [x] **UDP Client**
    - [x] **UDP Test**
    - [x] **UDP Demo**

- [x] **HTTP**
    - [x] **HTTP Server**
    - [x] **HTTP Client**
    - [x] **HTTP Test**
    - [x] **HTTP Demo**

- [x] **RPC**
    - [x] **RPC Server**
    - [x] **RPC Client**
    - [x] **RPC Test**
    - [ ] **RPC Demo**

- [ ] **Websocket**
    - [x] **Websocket Server**
    - [x] **Websocket Client**
    - [ ] **Websocket Test**
    - [ ] **Websocket Demo**


Feel free to explore each section's source code, tests, and demos to better understand the implementation of these networking concepts in Go. Contributions and feedback are welcome!