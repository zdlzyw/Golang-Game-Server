package main

import (
	"Frame/frame/iface"
	"Frame/frame/net"
	"fmt"
)

type PingRouter struct {
	net.BaseRouter
}

func (*PingRouter) PreHandle(request iface.IRequest) {
	fmt.Println("Call Router PreHandle...")
	if _, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping...\n")); err != nil {
		fmt.Println("Call Back Before Ping error, ", err)
	}
}
func (*PingRouter) Handle(request iface.IRequest) {
	fmt.Println("Call Router Handle...")
	if _, err := request.GetConnection().GetTCPConnection().Write([]byte("ping ping...\n")); err != nil {
		fmt.Println("Call Back Ping error, ", err)
	}
}
func (*PingRouter) PostHandle(request iface.IRequest) {
	fmt.Println("Call Router PostHandle...")
	if _, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping...\n")); err != nil {
		fmt.Println("Call Back after Ping error, ", err)
	}
}
func main() {
	s := net.NewServer()
	s.AddRouter(&PingRouter{})
	s.Serve()
}
