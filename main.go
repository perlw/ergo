package main

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"strings"
)

func spawnGame(conn net.Conn) {
	defer conn.Close()

	conn.Write([]byte(">"))

	r := bufio.NewReader(conn)
	buffer := make([]byte, 1)
	line := bytes.Buffer{}
	for {
		n, err := r.Read(buffer)
		if err != nil {
			log.Println("could not read from client, ", err)
			return
		} else if n <= 0 {
			continue
		}

		if buffer[0] == 0xff {
			cmd := make([]byte, 2)
			n, err := r.Read(cmd)
			if n <= 0 || err != nil {
				log.Println("could not read cmd from client, ", err)
				return
			}
			fmt.Println("cmd: ", cmd)
		} else if buffer[0] == '\n' {
			fmt.Println("client says, ", line.Bytes(), line.String())

			if strings.TrimRight(line.String(), "\r") == "quit" {
				conn.Write([]byte("BYE"))
				return
			}

			line.Reset()
			conn.Write([]byte(">"))
		} else {
			line.WriteByte(buffer[0])
		}
	}
}

func telnet() {
	ln, err := net.Listen("tcp", ":23")
	if err != nil {
		panic(err)
	}

	for {
		conn, _ := ln.Accept()
		go spawnGame(conn)
	}
}

func web() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		t.Execute(w, nil)
	})
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("could not start server, ", err)
	}
}

func main() {
	fmt.Println("Launching... well.. me")

	go telnet()
	go web()

	var forever chan int = nil
	<-forever
}
