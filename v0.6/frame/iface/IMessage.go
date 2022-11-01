package iface

type IMessage interface {
	GetMsgID() uint16
	GetMsgLen() uint16
	GetMsgData() []byte

	SetMsgID(uint16)
	SetMsgLen(uint16)
	SetMsgData([]byte)
}
