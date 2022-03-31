package main

import (
	"io"

	"net/http"
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

func HttpServiceListen(address string) {
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			Writer:     w,
			ReadCloser: r.Body,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))

	})
	http.ListenAndServe(address, nil)

}

func main() {
	RegisterService("User", new(User))
	HttpServiceListen(":1234")

}
