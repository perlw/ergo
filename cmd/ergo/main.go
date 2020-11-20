package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var port int
	var webBaseDir string

	flag.StringVar(
		&webBaseDir, "web-base-dir", "./web", "sets the base dir of the webapp",
	)
	flag.IntVar(&port, "port", 80, "set port to use")
	flag.Parse()

	if webBaseDir == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if webBaseDir[len(webBaseDir)-1] != '/' {
		webBaseDir += "/"
	}

	logger := log.New(os.Stdout, "ergo: ", log.LstdFlags)

	templates := template.New("")
	templates.Funcs(template.FuncMap{})
	templates, err := templates.ParseFiles(webBaseDir + "template/index.tpl")
	if err != nil {
		logger.Fatalf("could not prepare templates: %s", err.Error())
	}

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir(webBaseDir + "static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		err := templates.ExecuteTemplate(w, "index.tpl", nil)
		if err != nil {
			logger.Println("ERR:", err)
			http.Error(w, "cat on keyboard", http.StatusInternalServerError)
		}
	})

	mux.HandleFunc(
		"/about", func(w http.ResponseWriter, req *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("TBD"))
		},
	)

	mux.HandleFunc(
		"/healthcheck", func(w http.ResponseWriter, req *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Service OK"))
		},
	)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		ReadTimeout:  90 * time.Second,
		WriteTimeout: 90 * time.Second,
		Handler:      mux,
	}

	logger.Printf("Set up complete, using: %+v\n", map[string]interface{}{
		"port":         port,
		"web_base_dir": webBaseDir,
	})

	logger.Printf("Listening to %s\n", server.Addr)

	serverErr := make(chan error)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			serverErr <- err
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	select {
	case err := <-serverErr:
		logger.Fatalf("fatal server issue: %s", err.Error())
	case <-stop:
		logger.Println("Winding down...")
	}
}
