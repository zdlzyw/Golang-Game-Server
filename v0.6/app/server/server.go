package main

import (
	"Frame/frame/iface"
	"Frame/frame/net"
	"fmt"
)

type PingRouter struct {
	net.BaseRouter
}
type Hello struct {
	net.BaseRouter
}

func (*PingRouter) Handle(request iface.IRequest) {
	fmt.Print("Call Router Handle...\t")
	fmt.Printf("reve from client mstID=:%d, data=%s.\n", request.GetMsgID(), request.GetMsgData())
	if err := request.GetConnection().SendMsg(uint16(200), []byte("ping...ping...ping...")); err != nil {
		fmt.Println("Send msg error, ", err)
	}
}
func (*Hello) Handle(request iface.IRequest) {
	fmt.Print("Call Router Handle...\t")
	fmt.Printf("reve from client mstID=:%d, data=%s.\n", request.GetMsgID(), request.GetMsgData())
	if err := request.GetConnection().SendMsg(uint16(201), []byte("Hello...")); err != nil {
		fmt.Println("Send msg error, ", err)
	}
}

func main() {
	s := net.NewServer()
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &Hello{})
	s.Serve()
}
