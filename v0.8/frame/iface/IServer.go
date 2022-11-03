package iface

// IServer 定义服务器的三个方法接口(开启、停止、运行)
type IServer interface {
	Start()
	Stop()
	Serve()
	AddRouter(msgID uint16, router IRouter)
}
