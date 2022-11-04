package net

import (
	"Frame/frame/iface"
	"Frame/frame/utils"
	"errors"
	"fmt"
	"io"
	"net"
)

// Connection 连接模块实现。包含套接字、连接ID、状态、绑定的回调函数、停止连接的Channel
type Connection struct {
	Server   iface.IServer
	Conn     *net.TCPConn
	ConnID   uint32
	IsClosed bool
	// 无缓冲管道，用于读、写Goroutine的消息通信
	MsgChan    chan []byte
	ExitChan   chan bool
	MsgHandler iface.IMsgHandler
}

// NewConnection 初始化连接方法
func NewConnection(server iface.IServer, conn *net.TCPConn, connID uint32, msgHandler iface.IMsgHandler) *Connection {
	c := &Connection{
		Server:     server,
		Conn:       conn,
		ConnID:     connID,
		IsClosed:   false,
		MsgChan:    make(chan []byte),
		MsgHandler: msgHandler,
		ExitChan:   make(chan bool, 1),
	}
	// 调用连接管理，发送连接到管理器内
	c.Server.GetConnMgr().ConnAdd(c)
	return c
}

// StartWriter 协程写业务
func (c *Connection) StartWriter() {
	fmt.Println("[Start Write Goroutine is running...]")
	// 如果客户端退出的话输出调试信息
	defer fmt.Printf("[Conn Writer exit! %s]\n", c.RemoteAddr().String())

	// 阻塞等待Channel消息
	for {
		select {
		// SendMsg接收到数据要写给客户端
		case data := <-c.MsgChan:
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("Send data error, ", err)
			}
		// 代表Reader已退出，此时Writer同时退出
		case <-c.ExitChan:
			return
		}

	}
}

// StartReader 协程读业务
func (c *Connection) StartReader() {
	fmt.Println("[Start Reader Goroutine is running...]")
	defer fmt.Printf("[Conn Reader exit! %s]\n", c.RemoteAddr().String())
	// 读业务中任意方法break后都会调用Stop，触发关闭channel
	defer c.Stop()
	// request ID
	var rid uint32 = 0
	// 阻塞等待读取到数据
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
			id:   rid,
			msg:  msg,
		}
		rid++
		// 判断是否开启了池，如果开启了才进行处理，未开启则直接调用Goroutine
		if utils.GlobalObject.WorkerPoolSize > 0 {
			c.MsgHandler.ReqToQueue(&req)
		} else {
			go c.MsgHandler.DoMsgHandler(&req)
		}

	}
}

func (c *Connection) Start() {
	fmt.Println("Connection Start().. Connection ID =", c.ConnID)
	go c.StartReader()
	go c.StartWriter()
	// 创建连接后即调用，执行Hook函数
	c.Server.CallOnConnStart(c)
}

// Stop 判断连接状态决定是否需要关闭socket和回收资源
func (c *Connection) Stop() {
	fmt.Println("Connection Stop().. Connection ID = ", c.ConnID)
	if c.IsClosed {
		return
	}
	c.IsClosed = true
	// 销毁连接前调用，执行Hook函数
	c.Server.CallOnConnStop(c)
	err := c.Conn.Close()
	if err != nil {
		fmt.Println("Close Connection error, ", err)
	}
	// 告知关闭Write
	c.ExitChan <- true
	c.Server.GetConnMgr().ConnDel(c)
	close(c.ExitChan)
	close(c.MsgChan)
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
	// 不再直接会写给客户端，通过Chan发送给写方法
	c.MsgChan <- binaryMsg
	return nil
}
