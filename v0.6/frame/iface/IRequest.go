package iface

// IRequest 请求模块抽象，得到连接和消息内容
type IRequest interface {
	GetConnection() IConnection
	GetMsgData() []byte
	GetMsgID() uint16
}
