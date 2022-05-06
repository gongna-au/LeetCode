package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type ServerServiceInterface = interface {
	Service(request string, reply *string) error
}

func RegisterService(name string, svc ServerServiceInterface) error {
	return rpc.RegisterName(name, svc)

}

type User struct {
}

func (u *User) Service(request string, reply *string) error {

	*reply = request + " " + "Success!"
	return nil
}

func ListenService(network string, address string) {
	listener, err := net.Listen(network, address)
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeCodec(jsonrpc.NewServerCodec(conn))

}

func main() {
	RegisterService("User", new(User))
	ListenService("tcp", ":1234")

}
