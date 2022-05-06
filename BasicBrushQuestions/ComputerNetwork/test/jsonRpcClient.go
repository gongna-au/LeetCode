package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Client struct {
	*rpc.Client
}

func Dial(network string, address string) *Client {
	conn, err := net.Dial(network, address)
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	if err != nil {
		log.Fatal(err)
	}
	return &Client{
		client,
	}
}
func (c *Client) ClientCall(serviceMethod string, args string, reply *string) error {

	err := c.Call(serviceMethod+".Service", args, reply)
	if err != nil {
		log.Fatal(err)
		return err

	}
	return nil

}
func main() {
	client := Dial("tcp", "localhost:1234")
	var reply string
	err := client.ClientCall("User", "user request", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)

}
