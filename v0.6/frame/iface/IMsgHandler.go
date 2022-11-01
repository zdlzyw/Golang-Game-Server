package iface

// IMsgHandler 路由管理层，执行路由方法、添加路由
type IMsgHandler interface {
	DoMsgHandler(request IRequest)
	AddRouter(msgID uint16, router IRouter)
}
