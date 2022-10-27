package main

import "Frame/frame/net"

func main() {
	s := net.NewServer("game")
	s.Serve()
}
