package iface

// IMsgHandler 路由管理层，执行路由方法、添加路由、启动工作池、请求进入队列
type IMsgHandler interface {
	DoMsgHandler(request IRequest)
	AddRouter(msgID uint16, router IRouter)
	StartWorkerPool()
	ReqToQueue(request IRequest)
}
