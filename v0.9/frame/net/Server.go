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
	ConnManager    iface.IConnManager

	OnConnStart func(conn iface.IConnection)
	OnConnStop  func(conn iface.IConnection)
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
		fmt.Printf("Start server success, %s success, Listening...\n", s.Name)
		// cid 临时进行ID自增
		var cid uint32 = 0
		for {
			connect, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept tcp error, ", err)
				continue
			}
			if s.ConnManager.ConnAmount() >= utils.GlobalObject.MaxConn {
				// to do  发送给客户端超出最大连接数的错误包
				fmt.Println("[Too Many Connection ,closed!]")
				if err := connect.Close(); err != nil {
					fmt.Println(err)
				}
				continue
			}

			// 处理新连接业务的方法由Connection进行绑定及处理
			dealConn := NewConnection(s, connect, cid, s.MsgHandler)
			cid++
			go dealConn.Start()
		}
	}()
}

// Stop 服务器停止
func (s *Server) Stop() {
	fmt.Println("[Server Stop! Clear Connections Manager!]")
	s.ConnManager.ConnClear()
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

func (s *Server) GetConnMgr() iface.IConnManager {
	return s.ConnManager
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
		ConnManager:    NewConnManager(),
	}
	return s
}

/*
	注册、调用钩子方法。调用时需判断当前指针是否有方法，如果没有则不调用
*/
func (s *Server) SetOnConnStart(hook func(connection iface.IConnection)) {
	s.OnConnStart = hook
}
func (s *Server) SetOnConnStop(hook func(connection iface.IConnection)) {
	s.OnConnStop = hook
}
func (s *Server) CallOnConnStart(conn iface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("--->Call OnConnStart...")
		s.OnConnStart(conn)
	}
}
func (s *Server) CallOnConnStop(conn iface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("--->Call OnConnStop...")
		s.OnConnStop(conn)
	}
}
