package iface

import "net"

// IConnection 连接模块抽象。包含启动、停止、连接套接字、ID、远端信息、数据发送方法
type IConnection interface {
	Start()
	Stop()
	GetTCPConnection() *net.TCPConn
	GetConnID() uint32
	RemoteAddr() net.Addr
	SendMsg(msgId uint16, data []byte) error

	// 连接属性的增删查属性
	SetProperty(key string, value interface{})
	GetProperty(key string) (interface{}, error)
	DelProperty(key string)
}

// HandleFunc 处理连接业务方法
type HandleFunc func(*net.TCPConn, []byte, int) error
