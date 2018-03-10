package main

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
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

type Router struct {
	mu sync.RWMutex
	m  map[string]routerEntry
}

type routerEntry struct {
	h       http.Handler
	pattern string
}

func NewRouter() *Router {
	return &Router{}
}

func (mux *Router) Handler(r *http.Request) http.Handler {
	mux.mu.RLock()
	defer mux.mu.RUnlock()

	v, ok := mux.m[r.URL.Path]
	if ok {
		return v.h
	}

	fmt.Println("here I am", r.URL.Path)
	path := strings.Split(r.URL.Path, "/")
	if len(path) > 1 {
		path = path[1:]
	}
	fmt.Println(path)

	longest := -1
	for _, entry := range mux.m {
		entryPath := strings.Split(entry.pattern, "/")
		if len(entryPath) > 1 {
			entryPath = entryPath[1:]
		}
		fmt.Println("entry", entryPath)

		count := 0
		for t := range entryPath {
			if path[t] != entryPath[t] {
				break
			} else {
				count++
			}
		}
		if count > 0 && count > longest {
			fmt.Println("marking", longest)
			longest = count
		}
		fmt.Println("longest", path, "vs", entryPath, longest)
	}

	if longest > -1 {
		subPath := "/" + strings.Join(path[:longest], "/")
		fmt.Println("going with", subPath)
		v, ok = mux.m[subPath]
		if ok {
			return v.h
		}
	}

	return http.NotFoundHandler()
}

func (mux *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "*" {
		if r.ProtoAtLeast(1, 1) {
			w.Header().Set("Connection", "close")
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h := mux.Handler(r)

	h.ServeHTTP(w, r)
}

func (mux *Router) Handle(pattern string, handler http.Handler) {
	mux.mu.Lock()
	defer mux.mu.Unlock()

	if pattern == "" {
		panic("http: invalid pattern")
	}
	if handler == nil {
		panic("http: nil handler")
	}
	if _, exist := mux.m[pattern]; exist {
		panic("http: multiple registrations for " + pattern)
	}

	if mux.m == nil {
		mux.m = make(map[string]routerEntry)
	}
	mux.m[pattern] = routerEntry{h: handler, pattern: pattern}
}

func serveWeb() {
	router := NewRouter()

	router2 := NewRouter()
	router2.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("inside api"))
	}))
	router2.Handle("/note", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("note me senpai"))
	}))

	router.Handle("/foo", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("foo"))
	}))
	router.Handle("/api", router2)
	router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		t.Execute(w, nil)
	}))

	err := http.ListenAndServe(":80", router)
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
