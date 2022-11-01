package net

import "Frame/frame/iface"

type Request struct {
	conn iface.IConnection
	msg  iface.IMessage
}

func (r *Request) GetConnection() iface.IConnection {
	return r.conn
}

func (r *Request) GetMsgData() []byte {
	return r.msg.GetMsgData()
}

func (r *Request) GetMsgID() uint16 {
	return r.msg.GetMsgID()
}
