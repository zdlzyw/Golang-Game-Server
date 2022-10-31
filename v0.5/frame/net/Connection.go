package net

import (
	"Frame/frame/iface"
	"errors"
	"fmt"
	"io"
	"net"
)

// Connection 连接模块实现。包含套接字、连接ID、状态、绑定的回调函数、停止连接的Channel
type Connection struct {
	Conn     *net.TCPConn
	ConnID   uint32
	IsClosed bool
	ExitChan chan bool
	Router   iface.IRouter
}

// NewConnection 初始化连接方法
func NewConnection(conn *net.TCPConn, connID uint32, router iface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		IsClosed: false,
		Router:   router,
		ExitChan: make(chan bool, 1),
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
		dp := NewDataPack()
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read head error, ", err)
			break
		}
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("server unpack head error, ", err)
			break
		}
		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("server unpack data error, ", err)
				return
			}
		}
		msg.SetMsgData(data)

		req := Request{
			conn: c,
			msg:  msg,
		}
		go func(request iface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)

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

func (c *Connection) SendMsg(msgId uint16, data []byte) error {
	if c.IsClosed {
		return errors.New("connection closed when send message")
	}
	dp := NewDataPack()
	binaryMsg, err := dp.Pack(NewMessage(msgId, data))
	if err != nil {
		fmt.Println("Send data pack error, ", err)
		return errors.New("pack errors message")
	}
	if _, err := c.Conn.Write(binaryMsg); err != nil {
		fmt.Printf("Write msg id%d, error. err:%s.\n", msgId, err)
		return errors.New("conn write error")
	}
	return nil
}
