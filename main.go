package main

import (
	"context"
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"net"
	"net/http"
	"time"
)

func serveGame() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, _ := ln.Accept()
		go spawnGame(conn)
	}
}

type middlewareFunc func(routeFunc) routeFunc
type routeFunc func(http.ResponseWriter, *http.Request, context.Context) error

// Define own writer?
func baseRoute(method string, handler routeFunc, funcs ...middlewareFunc) http.HandlerFunc {
	handler = func() routeFunc {
		for t := len(funcs) - 1; t >= 0; t-- {
			handler = funcs[t](handler)
		}
		return handler
	}()
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		if r.Method != method {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		if err := handler(w, r, ctx); err != nil {
			// TODO: Only catch uncaught errors
			log.Println("\tsomething went wrong;", err)
		}
	}
}

/*type contextKey string

func (c contextKey) String() string {
	return "me context key " + string(c)
}

var (
	contextKeyBodyJson = contextKey("body/json")
)*/

func isErr(err error) string {
	if err != nil {
		return "FAIL"
	}
	return "OK"
}

func withHitLogger() middlewareFunc {
	return func(handler routeFunc) routeFunc {
		return func(w http.ResponseWriter, r *http.Request, ctx context.Context) error {
			start := time.Now()
			err := handler(w, r, ctx)
			log.Printf("%s [%s] in %.2fms", r.URL.Path, isErr(err), float64(time.Since(start))/float64(time.Millisecond))
			return err
		}
	}
}

func withJsonBody() middlewareFunc {
	return func(handler routeFunc) routeFunc {
		return func(w http.ResponseWriter, r *http.Request, ctx context.Context) error {
			if r.Header.Get("Content-Type") != "application/json" {
				err := errors.New("incorrect content-type, expected json")
				http.Error(w, err.Error(), http.StatusBadRequest)
				return err
			}

			return handler(w, r, ctx)
		}
	}
}

func serveWeb() {
	// Note mangement
	http.HandleFunc("/api/note", baseRoute("POST", func(w http.ResponseWriter, r *http.Request, ctx context.Context) error {
		var note struct {
			Note string `json:"note"`
		}
		err := json.NewDecoder(r.Body).Decode(&note)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return nil
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("note says: " + note.Note))
		return nil
	}, withHitLogger(), withJsonBody()))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		t.Execute(w, nil)
	})

	err := http.ListenAndServe(":8000", nil)
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
	log.Println(store)

	go serveGame()
	go serveWeb()

	var forever chan int
	<-forever
}
