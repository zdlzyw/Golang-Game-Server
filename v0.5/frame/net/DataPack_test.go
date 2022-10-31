package net

import (
	"fmt"
	"io"
	"net"
	"testing"
)

// TestDataPack 封包拆包功能单元测试
func TestDataPack(t *testing.T) {
	listener, err := net.Listen("tcp4", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("Listener server error, ", err)
		return
	}

	go func() {
		for {
			connect, err := listener.Accept()
			if err != nil {
				fmt.Println("server accept error, ", err)
				return
			}
			go func(connect net.Conn) {
				dp := NewDataPack()
				for {
					headData := make([]byte, dp.GetHeadLen())
					if _, err := io.ReadFull(connect, headData); err != nil {
						fmt.Println("read head error, ", err)
						break
					}
					msgHead, err := dp.Unpack(headData)
					if err != nil {
						fmt.Println("server unpack head error, ", err)
						return
					}
					if msgHead.GetMsgLen() > 0 {
						msg := msgHead.(*Message)
						msg.Data = make([]byte, msg.GetMsgLen())

						if _, err := io.ReadFull(connect, msg.Data); err != nil {
							fmt.Println("server unpack data error, ", err)
							return
						}
						fmt.Printf("---> Recv MsgID %d, DataLength = %d, Data = %s.\n", msg.Id, msg.DataLen, msg.Data)
					} else {
						fmt.Printf("---> Recv MsgID %d, DataLength = %d.\n", msgHead.GetMsgID(), msgHead.GetMsgLen())
						continue
					}
				}

			}(connect)
		}
	}()

	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("client dial error, ", err)
	}

	dp := NewDataPack()
	msg1 := &Message{
		Id:      1,
		DataLen: 11,
		Data:    []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'},
	}
	msg2 := &Message{
		Id:      2,
		DataLen: 0,
		Data:    nil,
	}
	msg3 := &Message{
		Id:      3,
		DataLen: 5,
		Data:    []byte{'1', '2', '3', '4', '5'},
	}
	sendData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("client Pack msg1 error, ", err)
	}
	sendData2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("client Pack msg2 error, ", err)
	}
	sendData3, err := dp.Pack(msg3)
	if err != nil {
		fmt.Println("client Pack msg3 error, ", err)
	}
	sendData1 = append(sendData1, sendData2...)
	sendData1 = append(sendData1, sendData3...)
	if _, err := conn.Write(sendData1); err != nil {
		fmt.Println("client write data error, ", err)
		return
	}
	select {}
}
