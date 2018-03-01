package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Launching... well.. me")

	ln, err := net.Listen("tcp", ":23")
	if err != nil {
		panic(err)
	}

	for {
		conn, _ := ln.Accept()
		go func() {
			conn.Write([]byte("Ping\n"))
			conn.Close()
		}()
	}
}
