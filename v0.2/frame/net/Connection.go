package net

import (
	"Frame/frame/iface"
	"fmt"
	"net"
)

// Connection 连接模块实现。包含套接字、连接ID、状态、绑定的回调函数、停止连接的Channel
type Connection struct {
	Conn      *net.TCPConn
	ConnID    uint32
	IsClosed  bool
	HandleAPI iface.HandleFunc
	ExitChan  chan bool
}

// NewConnection 初始化连接方法
func NewConnection(conn *net.TCPConn, connID uint32, callBackApi iface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		IsClosed:  false,
		HandleAPI: callBackApi,
		ExitChan:  make(chan bool, 1),
	}
	return c
}

// StartWriter 协程写业务
func (c *Connection) StartWriter() {

}

// StartReader 协程读业务
func (c *Connection) StartReader() {
	fmt.Println("Start Reader Goroutine is running...")
	defer fmt.Printf("Connection ID = %d Reader is exit, remote addr is %s .\n", c.ConnID, c.RemoteAddr().String())
	defer c.Stop()

	for {
		buffer := make([]byte, 512)
		content, err := c.Conn.Read(buffer)
		if err != nil {
			fmt.Println("receive buffer error, ", err)
			continue
		}
		if err := c.HandleAPI(c.Conn, buffer, content); err != nil {
			fmt.Printf("Connect ID %d handle is error, %s.\n", c.ConnID, err)
			break
		}
	}
}

func (c *Connection) Start() {
	fmt.Println("Connection Start().. Connection ID =", c.ConnID)
	go c.StartReader()
}

// Stop 判断连接状态决定是否需要关闭socket和回收资源
func (c *Connection) Stop() {
	fmt.Println("Connection Stop().. Connection ID = ", c.ConnID)
	if !c.IsClosed {
		return
	}
	c.IsClosed = true
	err := c.Conn.Close()
	if err != nil {
		fmt.Println("Close Connection error, ", err)
	}
	close(c.ExitChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(data []byte) error {
	data = nil
	return nil
}
