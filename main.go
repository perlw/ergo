package main

import (
	"fmt"
	"net"
)

func telnet() {
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

func web() {
	ln, err := net.Listen("tcp", ":80")
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

func main() {
	fmt.Println("Launching... well.. me")

	var forever chan int = nil
	go telnet()
	go web()
	<-forever
}
