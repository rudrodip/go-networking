package tcp

func main() {
	server, err := NewTCPServer()
	if err != nil {
		panic(err)
	}

	err = server.Start()
	if err != nil {
		panic(err)
	}

	defer server.Stop()
}
