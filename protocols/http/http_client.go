package http

import (
	"fmt"
	"io"
	"net/http"
)

type HttpClient struct {
	client *http.Client
}

func NewHttpClient() (*HttpClient, error) {
	client := &HttpClient{
		client: &http.Client{},
	}
	return client, nil
}

func (c *HttpClient) Get(url string) string {
	defer c.Stop()

	resp, err := c.client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response: %v\n", string(data))
	return string(data)
}

func (c *HttpClient) Stop() {
	c.client.CloseIdleConnections()
}
