package iface

// IConnManager 连接管理模块，抽象出添加、删除、ID查找、数量、删除所有（退出服务）
type IConnManager interface {
	ConnAdd(conn IConnection)
	ConnDel(conn IConnection)
	ConnFindById(connID uint32) (IConnection, error)
	ConnAmount() int
	ConnClear()
}
