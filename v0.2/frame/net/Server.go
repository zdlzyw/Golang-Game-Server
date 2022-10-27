package net

import (
	"Frame/frame/iface"
	"errors"
	"fmt"
	"net"
)

// Server 服务结构体
type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      uint16
}

// CallBackToClient 定义客户端绑定的handle api（应由用户自定义）
func CallBackToClient(conn *net.TCPConn, data []byte, length int) error {
	fmt.Println("[Connection Handle] Call Back TO Client ...")
	if _, err := conn.Write(data[:length]); err != nil {
		fmt.Println("write back buffer error, ", err)
		return errors.New("CallBackToClient error")
	}
	return nil
}

// Start 服务启动
func (s *Server) Start() {
	fmt.Printf("[Start] net Listenner at %s : %d is starting!\n", s.IP, s.Port)

	go func() {
		address, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp add error :", err)
		}
		listener, err := net.ListenTCP(s.IPVersion, address)
		if err != nil {
			fmt.Println("listen tcp error :", err)
		}
		fmt.Println("Start server success, ", s.Name, " success, Listening...")
		var cid uint32 = 0
		for {
			connect, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept tcp error, ", err)
			}
			// 处理新连接业务的方法由Connection进行绑定及处理
			dealConn := NewConnection(connect, cid, CallBackToClient)
			cid++
			go dealConn.Start()
		}
	}()
}

// Stop 服务器停止
func (s *Server) Stop() {

}

// Serve 运行服务器
func (s *Server) Serve() {
	s.Start()
	// 阻塞状态，避免启动后立刻停止
	select {}
}

// NewServer 初始化Server模块
func NewServer(name string) iface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8000,
	}
	return s
}
