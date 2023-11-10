package main

import (
	"dependncy-injection/business"
	"dependncy-injection/database"
	"dependncy-injection/service"
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
	db := database.NewDatabase()
	logic := business.NewBusiness(db)
	svc := service.NewService(logic)
	client := NewClient(svc)

	fmt.Println(client.MakeRequest())
}
