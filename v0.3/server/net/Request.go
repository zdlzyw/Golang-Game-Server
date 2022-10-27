package net

import "Server/server/iface"

type Request struct {
	conn iface.IConnection
	data []byte
}

func (r *Request) GetConnection() iface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}
