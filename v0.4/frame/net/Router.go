package net

import "Frame/frame/iface"

// BaseRouter Router基类，以基类实现Router，使用时可根据需求继承、重写该方法
// 实体类继承抽象类后需要实现全部抽象方法，在此通过实体类先实现形成基类，重写该方法时因为是实体类可指定重写的方法，不需要实现所有方法
type BaseRouter struct {
}

func (br *BaseRouter) PreHandle(iface.IRequest)  {}
func (br *BaseRouter) Handle(iface.IRequest)     {}
func (br *BaseRouter) PostHandle(iface.IRequest) {}
