package iface

// IServer 定义服务器的三个方法接口(开启、停止、运行)，添加路由，通过server获取当前连接管理
type IServer interface {
	Start()
	Stop()
	Serve()
	AddRouter(msgID uint16, router IRouter)
	GetConnMgr() IConnManager
	// 钩子函数注册、调用方法
	SetOnConnStart(func(connection IConnection))
	SetOnConnStop(func(connection IConnection))
	CallOnConnStart(connection IConnection)
	CallOnConnStop(connection IConnection)
}
