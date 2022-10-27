package iface

import "net"

// IConnection 连接模块抽象。包含启动、停止、连接套接字、ID、远端信息、数据发送方法
type IConnection interface {
	Start()
	Stop()
	GetTCPConnection() *net.TCPConn
	GetConnID() uint32
	RemoteAddr() net.Addr
	Send(data []byte) error
}

// HandleFunc 处理连接业务方法
type HandleFunc func(*net.TCPConn, []byte, int) error
