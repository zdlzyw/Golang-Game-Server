package main

import (
	"Frame/frame/iface"
	"Frame/frame/net"
	"fmt"
)

type PingRouter struct {
	net.BaseRouter
}

func (*PingRouter) Handle(request iface.IRequest) {
	fmt.Print("Call Router Handle...\t")
	fmt.Printf("reve from client mstID=:%d, data=%s.\n", request.GetMsgID(), request.GetMsgData())
	if err := request.GetConnection().SendMsg(uint16(1), []byte("ping...ping...ping...")); err != nil {
		fmt.Println("Send msg error, ", err)
	}
}

func main() {
	s := net.NewServer()
	s.AddRouter(&PingRouter{})
	s.Serve()
}
