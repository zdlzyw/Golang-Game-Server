package main

import "Server/server/net"

func main() {
	s := net.NewServer("game")
	s.Serve()
}
