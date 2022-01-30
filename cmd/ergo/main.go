package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/fs"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/gomarkdown/markdown"
)

var musings map[string]string

type CacheData struct {
	t time.Time
	d []byte
}

const CacheSize = 32

type Cache map[string]CacheData

func (c Cache) get(key string) ([]byte, bool) {
	if cd, ok := c[key]; ok {
		cd.t = time.Now()
		return cd.d, true
	}
	return nil, false
}

func (c Cache) set(key string, data []byte) {
	cd, ok := c[key]
	if !ok {
		if len(c) >= CacheSize {
			var longest time.Duration = 0
			var longestKey string
			now := time.Now()
			for k, v := range c {
				diff := now.Sub(v.t)
				if diff > longest {
					longest = diff
					longestKey = k
				}
			}
			delete(c, longestKey)
		}

		c[key] = CacheData{
			t: time.Now(),
			d: data,
		}
	} else {
		cd.t = time.Now()
		cd.d = data
	}
}

var cachedMusings Cache

func getMusing(musing string) ([]byte, error) {
	if data, ok := cachedMusings.get(musing); ok {
		return data, nil
	}

	path, ok := musings[musing]
	if !ok {
		return nil, fmt.Errorf("not-found: %s", musing)
	}

	rawMd, err := ioutil.ReadFile(path)
	if !ok {
		return nil, fmt.Errorf("not-found: %w", err)
	}

	result := markdown.ToHTML(markdown.NormalizeNewlines(rawMd), nil, nil)
	cachedMusings.set(musing, result)
	return result, nil
}

func parameterize(format, path string) (map[string]string, error) {
	result := make(map[string]string)

	formatParts := strings.Split(format, "/")
	pathParts := strings.Split(path, "/")

	if len(formatParts) == len(pathParts) {
		for i, p := range formatParts {
			if len(p) > 0 && p[0] == ':' {
				result[p[1:]] = pathParts[i]
			}
		}
		return result, nil
	}

	return nil, fmt.Errorf("invalid-params: expected %s got %s", format, path)
}

type WakaAPIStatsBlock struct {
	Decimal      string
	Digital      string
	Hours        int
	Minutes      int
	Name         string
	Percent      float64
	Seconds      int
	Text         string
	TotalSeconds float64 `json:"total_seconds"`
}

var wakatimeMutex sync.RWMutex
var wakatimeStats struct {
	CummulativeTotal struct {
		Decimal string
		Digital string
		Seconds float64
		Text    string
	} `json:"cummulative_total"`
	Data []struct {
		Categories   []WakaAPIStatsBlock
		Dependencies []WakaAPIStatsBlock
		Editors      []WakaAPIStatsBlock
		GrandTotal   struct {
			Decimal      string
			Digital      string
			Hours        int
			Minutes      int
			Text         string
			TotalSeconds float64 `json:"total_seconds"`
		} `json:"grand_total"`
		Languages        []WakaAPIStatsBlock
		Machines         []WakaAPIStatsBlock
		OperatingSystems []WakaAPIStatsBlock `json:"operating_systems"`
		Projects         []WakaAPIStatsBlock
		Range            struct {
			Date     string
			End      time.Time
			Start    time.Time
			Text     string
			Timezone string
		}
	}
}

func runWakatimeUpdate(apikeyb64 string, logger *log.Logger) error {
	logger.Println("Updating wakatime stats...")
	req, err := http.NewRequest(http.MethodGet, "https://wakatime.com/api/v1/users/current/summaries?range=Last%207%20Days", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+apikeyb64)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	wakatimeMutex.Lock()
	err = json.NewDecoder(resp.Body).Decode(&wakatimeStats)
	if err != nil {
		return err
	}
	wakatimeMutex.Unlock()

	return nil
}

var wakatimeStatsOn bool

func wakatimeStatsUpdater(apikey string, logger *log.Logger) {
	apikeyb64 := base64.StdEncoding.EncodeToString([]byte(apikey))
	if err := runWakatimeUpdate(apikeyb64, logger); err != nil {
		logger.Println("Could not initialize wakatime stats:", err)
		wakatimeStatsOn = false
		return
	}

	ticker := time.NewTicker(time.Hour * 24)
	for range ticker.C {
		if err := runWakatimeUpdate(apikeyb64, logger); err != nil {
			logger.Println("Could not update wakatime stats:", err)
			wakatimeStatsOn = false
			return
		}
	}
}

func main() {
	var port int
	var webBaseDir string
	var devMode bool

	flag.StringVar(
		&webBaseDir, "web-base-dir", "./", "sets the base dir of the webapp",
	)
	flag.IntVar(&port, "port", 80, "set port to use")
	flag.BoolVar(&devMode, "dev", false, "dev mode")
	flag.Parse()

	if webBaseDir == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if webBaseDir[len(webBaseDir)-1] != '/' {
		webBaseDir += "/"
	}

	logger := log.New(os.Stdout, "ergo: ", log.LstdFlags)

	pageTmpls := make(map[string]*template.Template)
	pageTmpls["home"] = template.Must(template.ParseFiles(webBaseDir+"template/page.tpl", webBaseDir+"template/home.tpl"))
	pageTmpls["musings"] = template.Must(template.ParseFiles(webBaseDir+"template/page.tpl", webBaseDir+"template/musings.tpl"))
	pageTmpls["musing"] = template.Must(template.ParseFiles(webBaseDir+"template/page.tpl", webBaseDir+"template/musing.tpl"))
	pageTmpls["why"] = template.Must(template.ParseFiles(webBaseDir+"template/page.tpl", webBaseDir+"template/why.tpl"))

	musings = make(map[string]string)
	cachedMusings = make(Cache)
	musingLinks := make(map[string][]string)
	filepath.WalkDir(webBaseDir+"musings/", func(path string, d fs.DirEntry, err error) error {
		if !devMode && d.IsDir() && strings.Contains(path, "/dev") {
			return filepath.SkipDir
		}

		if !d.IsDir() && filepath.Ext(path) == ".md" {
			parts := strings.Split(path, "/")
			date := parts[len(parts)-2]
			title := strings.TrimRight(parts[len(parts)-1], ".md")
			ident := date + "_" + title
			musings[ident] = path
			if _, ok := musingLinks[date]; !ok {
				musingLinks[date] = []string{}
			}
			musingLinks[date] = append(musingLinks[date], title)
		}
		return nil
	})

	for _, env := range os.Environ() {
		parts := strings.Split(env, "=")
		if parts[0] == "WAKATIME_APIKEY" {
			go wakatimeStatsUpdater(parts[1], logger)
			wakatimeStatsOn = true
			break
		}
	}
	if !wakatimeStatsOn {
		logger.Println("WARN: No WAKATIME_APIKEY in env, skipping wakatimeStatsUpdater.")
	}

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir(webBaseDir + "static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		var err error
		if wakatimeStatsOn {
			var totalSeconds float64
			wakaSeconds := make(map[string]float64)
			wakaPercentage := make(map[string]int)
			wakatimeMutex.RLock()
			for _, data := range wakatimeStats.Data {
				for _, lang := range data.Languages {
					if _, ok := wakaSeconds[lang.Name]; !ok {
						wakaSeconds[lang.Name] = 0
					}
					wakaSeconds[lang.Name] += lang.TotalSeconds
					totalSeconds += lang.TotalSeconds
				}
			}
			wakatimeMutex.RUnlock()

			for k, v := range wakaSeconds {
				wakaPercentage[k] = int(math.Ceil((v / totalSeconds) * 100.0))
			}

			err = pageTmpls["home"].Execute(w, map[string]interface{}{
				"wakaStats": wakaPercentage,
			})
		} else {
			err = pageTmpls["home"].Execute(w, nil)
		}
		if err != nil {
			logger.Println("ERR:", err)
			http.Error(w, "cat on keyboard", http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("/musings", func(w http.ResponseWriter, req *http.Request) {
		if err := pageTmpls["musings"].Execute(w, map[string]interface{}{
			"musingLinks": musingLinks,
		}); err != nil {
			logger.Println("ERR:", err)
			http.Error(w, "cat on keyboard", http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("/musing/", func(w http.ResponseWriter, req *http.Request) {
		parameters, err := parameterize("/musing/:date/:title", req.URL.Path)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		html, err := getMusing(parameters["date"] + "_" + parameters["title"])
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		if err := pageTmpls["musing"].Execute(w, map[string]interface{}{
			"title":   parameters["title"],
			"content": template.HTML(html),
		}); err != nil {
			logger.Println("ERR:", err)
			http.Error(w, "cat on keyboard", http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("/why", func(w http.ResponseWriter, req *http.Request) {
		if err := pageTmpls["why"].Execute(w, nil); err != nil {
			logger.Println("ERR:", err)
			http.Error(w, "cat on keyboard", http.StatusInternalServerError)
		}
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

	templates := make([]string, 0, 10)
	for k := range pageTmpls {
		templates = append(templates, k)
	}
	logger.Printf("Set up complete, using: %+v\n", map[string]interface{}{
		"port":         port,
		"web_base_dir": webBaseDir,
		"templates":    templates,
		"musings":      musings,
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
