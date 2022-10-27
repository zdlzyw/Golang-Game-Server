package main

import (
	"fmt"
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
		if _, err := connect.Write([]byte("here from client")); err != nil {
			fmt.Println("write connect error, ", err)
			return
		}
		buffer := make([]byte, 512)
		content, err := connect.Read(buffer)
		if err != nil {
			fmt.Println("read buffer error, ", err)
			return
		}
		fmt.Printf("server call back:%s, content = %d\n", buffer, content)
		time.Sleep(1 * time.Second)
	}

}
