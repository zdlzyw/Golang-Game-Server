package net

import (
	"Frame/frame/iface"
	"Frame/frame/utils"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

type DataPack struct {
}

func NewDataPack() *DataPack {
	return &DataPack{}
}

func (dp *DataPack) GetHeadLen() uint32 {
	// DataLen(uint16-2byte) + ID(uint16-2byte)
	return 4
}

// Pack 通过获得的数据组合数据包：包长|包ID|数据（小端模式）
func (dp *DataPack) Pack(msg iface.IMessage) ([]byte, error) {
	dataBuff := bytes.NewBuffer([]byte{})

	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		fmt.Println("Data Pack → Length error.")
		return nil, err
	}

	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgID()); err != nil {
		fmt.Println("Data Pack → ID error.")
		return nil, err
	}

	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgData()); err != nil {
		fmt.Println("Data Pack → Data error.")
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

// Unpack 通过获得的字节流解析数据包，最终分别对Message方法赋值长度、ID，返回数据包（小端模式）
func (dp *DataPack) Unpack(binaryData []byte) (iface.IMessage, error) {
	dataBuff := bytes.NewReader(binaryData)

	msg := &Message{}
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		fmt.Println("Data Unpack → Length error.")
		return nil, err
	}

	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id); err != nil {
		fmt.Println("Data Unpack → Length error.")
		return nil, err
	}

	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		fmt.Println("Too large msg data receive, nothing to do")
		return nil, errors.New(fmt.Sprint("Msg ID: ", &msg.Id, "is too large!"))
	}

	return msg, nil
}
