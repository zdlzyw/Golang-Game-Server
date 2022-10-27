package iface

// IRouter 处理请求数据的方法路由，分为预处理、处理、后处理
type IRouter interface {
	PreHandle(request IRequest)
	Handle(request IRequest)
	PostHandle(request IRequest)
}
