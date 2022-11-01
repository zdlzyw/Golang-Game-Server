package net

type Message struct {
	Id      uint16
	DataLen uint16
	Data    []byte
}
type DataType struct {
	int
}

func NewMessage(id uint16, data []byte) *Message {
	return &Message{
		Id:      id,
		DataLen: uint16(len(data)),
		Data:    data,
	}
}

func (m *Message) GetMsgID() uint16 {
	return m.Id
}

func (m *Message) GetMsgLen() uint16 {
	return m.DataLen
}

func (m *Message) GetMsgData() []byte {
	return m.Data
}

func (m *Message) SetMsgID(id uint16) {
	m.Id = id
}

func (m *Message) SetMsgLen(len uint16) {
	m.DataLen = len
}

func (m *Message) SetMsgData(data []byte) {
	m.Data = data
}
