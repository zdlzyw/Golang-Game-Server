package net

import (
	"Frame/frame/iface"
	"Frame/frame/utils"
	"fmt"
	"net"
)

// Server 服务结构体
type Server struct {
	Name           string
	IPVersion      string
	IP             string
	Port           uint16
	MsgHandler     iface.IMsgHandler
	MaxConn        int
	MaxPackageSize uint16
}

// Start 服务启动
func (s *Server) Start() {
	fmt.Printf("%s has [Start] ,Listener at %s:%d.Max Connect is %d, Max Package is %d.\n ", s.Name, s.IP, s.Port, s.MaxConn, s.MaxPackageSize)

	go func() {
		// 开启消息队列
		s.MsgHandler.StartWorkerPool()
		address, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp add error :", err)
		}
		listener, err := net.ListenTCP(s.IPVersion, address)
		if err != nil {
			fmt.Println("listen tcp error :", err)
		}
		fmt.Printf("Start server success, %s success, Listening...", s.Name)
		// cid 临时进行ID自增
		var cid uint32 = 0
		for {
			connect, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept tcp error, ", err)
			}
			// 处理新连接业务的方法由Connection进行绑定及处理
			dealConn := NewConnection(connect, cid, s.MsgHandler)
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

func (s *Server) AddRouter(msgID uint16, router iface.IRouter) {
	s.MsgHandler.AddRouter(msgID, router)

	fmt.Println("Add Route Success!")
}

// NewServer 初始化Server模块
func NewServer() iface.IServer {
	s := &Server{
		Name:           utils.GlobalObject.Name,
		IPVersion:      utils.GlobalObject.IPVersion,
		IP:             utils.GlobalObject.Host,
		Port:           utils.GlobalObject.TcpPort,
		MsgHandler:     NewMsgHandler(),
		MaxConn:        utils.GlobalObject.MaxConn,
		MaxPackageSize: utils.GlobalObject.MaxPackageSize,
	}
	return s
}
