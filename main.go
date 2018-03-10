package main

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
)

func serveGame() {
	ln, err := net.Listen("tcp", ":23")
	if err != nil {
		panic(err)
	}

	for {
		conn, _ := ln.Accept()
		go spawnGame(conn)
	}
}

func serveWeb() {
	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("foo"))
	})
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
		log.Fatal("could not start server,", err)
	}
}

func main() {
	log.Println("Launching... well.. me")

	store, err := NewStore()
	if err != nil {
		log.Fatalln("could not set up store", err)
	}
	fmt.Println(store)

	go serveGame()
	go serveWeb()

	var forever chan int
	<-forever
}
