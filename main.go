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
	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/note", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("foo"))
	})

	mux := http.NewServeMux()
	mux.Handle("/api/", apiMux)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		t.Execute(w, nil)
	})

	err := http.ListenAndServe(":80", mux)
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
