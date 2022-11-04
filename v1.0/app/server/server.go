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

func HookDoBegin(conn iface.IConnection) {
	fmt.Println("======================>Begin Hook Called")
	if err := conn.SendMsg(203, []byte("Begin Hook!!")); err != nil {
		fmt.Println(err)
	}
	conn.SetProperty("key", "value")
}

func HookDoStop(conn iface.IConnection) {
	fmt.Printf("======================>Stop Hook Called, ID = %d.\n", conn.GetConnID())
	if name, err := conn.GetProperty("key"); err == nil {
		fmt.Println("Name = ", name)
	}
}

func main() {
	s := net.NewServer()
	// 服务器开启后即注册Hook函数
	s.SetOnConnStart(HookDoBegin)
	s.SetOnConnStop(HookDoStop)

	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &Hello{})
	s.Serve()
}
