package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/go-ini/ini"
	"github.com/pkg/errors"
)

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

func main() {
	log.Println("┌starting up")
	cfg, err := ini.Load("ergo.ini")
	if err != nil {
		log.Fatalln(errors.Wrap(err, "├could not read config"))
	}
	cfg.BlockMode = false

	section := cfg.Section(ini.DEFAULT_SECTION)
	port := section.Key("port").MustInt(1337)

	store, err := NewStore()
	if err != nil {
		log.Fatalln("└could not set up store", err)
	}
	log.Println(store)

	mux := http.NewServeMux()
	// Note mangement
	mux.HandleFunc("/api/note", baseRoute("POST", func(w http.ResponseWriter, r *http.Request, ctx context.Context) error {
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

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		t.Execute(w, nil)
	})

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("└alive @ %s", server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("└could not start server,", err)
	}
}
