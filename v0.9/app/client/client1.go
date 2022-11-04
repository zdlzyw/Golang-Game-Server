package main

import (
	net2 "Frame/frame/net"
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	fmt.Println("client start...")
	time.Sleep(1 * time.Second)
	connect, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("client start error, exit!.", err)
		return
	}
	for {
		// 写数据
		dp := net2.NewDataPack()
		binaryMsg, err := dp.Pack(net2.NewMessage(uint16(1), []byte("Game client test message11111111111111!")))
		if err != nil {
			fmt.Println("client pack error, ", err)
			return
		}
		if _, err := connect.Write(binaryMsg); err != nil {
			fmt.Println("write binary error, ", err)
			return
		}

		// 读数据
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(connect, headData); err != nil {
			fmt.Println("client read head error, ", err)
			break
		}
		msgHead, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("client unpack head error, ", err)
			return
		}
		if msgHead.GetMsgLen() > 0 {
			msg := msgHead.(*net2.Message)
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

		time.Sleep(1 * time.Second)
	}

}
