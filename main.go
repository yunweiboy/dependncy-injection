package main

import (
	"dependncy-injection/service"
	"dependncy-injection/wire"
	"fmt"
)

// Client struct
type Client struct {
	service service.Service
}

// NewClient Constructor
func NewClient(service service.Service) *Client {
	return &Client{service: service}
}

// MakeRequest Call Service
func (c Client) MakeRequest() string {
	return "Client request: " + c.service.HandleRequest()
}

func main() {
	// make dependency injection manually
	svc := wire.InitializeService()
	client := NewClient(svc)

	fmt.Println(client.MakeRequest())
}
