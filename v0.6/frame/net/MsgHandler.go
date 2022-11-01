package net

import (
	"Frame/frame/iface"
	"fmt"
)

type MsgHandler struct {
	Apis map[uint16]iface.IRouter
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis: make(map[uint16]iface.IRouter),
	}
}

func (mh *MsgHandler) DoMsgHandler(request iface.IRequest) {
	handler, ok := mh.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("Api msgID ", request.GetMsgID(), " is not found!")
	}
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)

}
func (mh *MsgHandler) AddRouter(msgID uint16, router iface.IRouter) {
	if _, ok := mh.Apis[msgID]; ok {
		fmt.Printf("Msg Router ID %d is duplicate.\n", msgID)
	}
	mh.Apis[msgID] = router
	fmt.Printf("Add api success! ID=%d.\n", msgID)
}
