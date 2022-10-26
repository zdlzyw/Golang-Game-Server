package net

import (
	"Server/server/iface"
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

// Start 服务启动
func (s *Server) Start() {
	fmt.Printf("[Start] net Listenner at %s : %d is starting!\n", s.IP, s.Port)

	go func() {
		address, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp add error :", err)
		}
		listenner, err := net.ListenTCP(s.IPVersion, address)
		if err != nil {
			fmt.Println("listen tcp error :", err)
		}
		fmt.Println("Start server success, ", s.Name, " success, Listerning...")

		for {
			connect, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept tcp error, ", err)
			}

			go func() {
				for {
					buffer := make([]byte, 512)
					content, err := connect.Read(buffer)
					if err != nil {
						fmt.Println("receive buffer error, ", err)
						continue
					}
					fmt.Printf("receive buff from client: %s, content = %d.\n", buffer, content)
					if _, err := connect.Write(buffer[:content]); err != nil {
						fmt.Println("write back buffer error, ", err)
						continue
					}
				}
			}()
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
