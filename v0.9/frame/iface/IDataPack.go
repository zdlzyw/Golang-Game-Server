package iface

// IDataPack 数据包处理方法，获取Head长度，封包、拆包
type IDataPack interface {
	GetHeadLen() uint32
	Pack(msg IMessage) ([]byte, error)
	Unpack(msg IMessage) (IMessage, error)
}
