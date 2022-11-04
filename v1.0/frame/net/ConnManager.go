package net

import (
	"Frame/frame/iface"
	"errors"
	"fmt"
	"sync"
)

type ConnManager struct {
	connections map[uint32]iface.IConnection
	connLock    sync.RWMutex
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]iface.IConnection),
	}
}

// ConnAdd 向map中添加连接，先写锁保护，使用defer在调用完毕后释放，mapID为conn ID
func (cm *ConnManager) ConnAdd(conn iface.IConnection) {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	cm.connections[conn.GetConnID()] = conn
	fmt.Printf("[ConnManager Add Conn = %d.Now amount = %d .]\n]", conn.GetConnID(), cm.ConnAmount())
}
func (cm *ConnManager) ConnDel(conn iface.IConnection) {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	tmp := conn.GetConnID()
	delete(cm.connections, conn.GetConnID())
	fmt.Printf("[ConnManager Delete Conn = %d.Now amount = %d .\n]", tmp, cm.ConnAmount())
}
func (cm *ConnManager) ConnFindById(connID uint32) (iface.IConnection, error) {
	cm.connLock.RLock()
	defer cm.connLock.RUnlock()
	conn, ok := cm.connections[connID]
	if ok {
		fmt.Println("Found Out Connection.]")
		return conn, nil
	}
	return nil, errors.New("[Conn No Found!]")
}
func (cm *ConnManager) ConnAmount() int {
	return len(cm.connections)
}
func (cm *ConnManager) ConnClear() {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	cm.connections = make(map[uint32]iface.IConnection)
	for connID, conn := range cm.connections {
		conn.Stop()
		delete(cm.connections, connID)
	}
	fmt.Printf("[Cleat all connections success! Now amount = %d.]\n", cm.ConnAmount())
}
