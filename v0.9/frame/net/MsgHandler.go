package net

import (
	"Frame/frame/iface"
	"Frame/frame/utils"
	"fmt"
)

type MsgHandler struct {
	Apis           map[uint16]iface.IRouter
	MsgQueue       []chan iface.IRequest
	WorkerPoolSize uint32
}

// NewMsgHandler 定义消息集合，读取worker池数量，定义消息队列池（数量和worker池数量一致）
func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis:           make(map[uint16]iface.IRouter),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
		MsgQueue:       make([]chan iface.IRequest, utils.GlobalObject.WorkerPoolSize),
	}
}

func (mh *MsgHandler) DoMsgHandler(request iface.IRequest) {
	handler, ok := mh.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("Api msgID ", request.GetMsgID(), " is not found!")
	}
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)

}
func (mh *MsgHandler) AddRouter(msgID uint16, router iface.IRouter) {
	if _, ok := mh.Apis[msgID]; ok {
		fmt.Printf("Msg Router ID %d is duplicate.\n", msgID)
	}
	mh.Apis[msgID] = router
	fmt.Printf("Add api success! ID=%d.\n", msgID)
}

// StartWorkerPool 初始化Worker工作池。每次服务器启动只需执行一次（只能有一个Worker池），通过接口暴露出方法。每个Worker对应一个消息队列
func (mh *MsgHandler) StartWorkerPool() {
	fmt.Printf("[Start Worker Pool, Size = %d.]\n", utils.GlobalObject.WorkerPoolSize)
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		// 初始化消息队列，每个队列可存放消息数量固定
		mh.MsgQueue[i] = make(chan iface.IRequest, utils.GlobalObject.MaxQueueSize)
		go mh.startOneWorker(i, mh.MsgQueue[i])
	}
}

// StartOneWorker 开启一个Worker的阻塞Goroutine执行消息任务。对外不应暴露，根据size大小创建Goroutine
func (mh *MsgHandler) startOneWorker(workerID int, msgQueue chan iface.IRequest) {
	fmt.Printf("[Worker Started! ID = %d. \n", workerID)
	for {
		select {
		// 过来的消息为客户端的Request请求，当得到请求后执行DoMsgHandler方法
		case request := <-msgQueue:
			mh.DoMsgHandler(request)

		}
	}
}

// ReqToQueue 请求消息进入消息队列，根据算法平均轮训分配（简单）。每个request使用自增ID，通过池长度取余，平均分配到0-9池中
func (mh *MsgHandler) ReqToQueue(request iface.IRequest) {
	// 可根据连接ID进行分配
	workerID := request.GetRequestID() % mh.WorkerPoolSize
	fmt.Printf("[Add ConnID = %d,Request ID = %d, Msg ID = %d, Worker ID = %d .]\n ", request.GetConnection().GetConnID(), request.GetRequestID(), request.GetMsgID(), workerID)
	mh.MsgQueue[workerID] <- request
}
