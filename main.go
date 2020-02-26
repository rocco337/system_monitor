package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

//env GOOS=linux GOARCH=arm GOARM=5 go build

var addr = flag.String("addr", "127.0.0.1:8081", "http service address")
var refreshRate = 2 * time.Second

func main() {
	flag.Parse()
	fmt.Printf("Starting system monitor on address: %s \n", *addr)
	monitor := new(SystemMonitor)

	http.HandleFunc("/hostInfo", func(w http.ResponseWriter, r *http.Request) {
		serveJSON(w, r, false, func() interface{} {
			return monitor.GetHostInfo()
		})
	})

	http.HandleFunc("/processes", func(w http.ResponseWriter, r *http.Request) {
		serveJSON(w, r, true, func() interface{} {
			return monitor.GetProcesses()
		})
	})

	http.HandleFunc("/cpuusage", func(w http.ResponseWriter, r *http.Request) {
		serveJSON(w, r, true, func() interface{} {
			return monitor.GetCPUThreadsUsage()
		})
	})

	http.HandleFunc("/memorystat", func(w http.ResponseWriter, r *http.Request) {
		serveJSON(w, r, true, func() interface{} {
			return monitor.GetMemoryUsage()
		})
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serveFile(w, r, "home.html")
	})
	http.HandleFunc("/home.js", func(w http.ResponseWriter, r *http.Request) {
		serveFile(w, r, "home.js")
	})

	log.Fatal(http.ListenAndServe(*addr, nil))
}

func serveFile(w http.ResponseWriter, r *http.Request, filename string) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, filename)
}

var upgrader = websocket.Upgrader{}

func serveJSON(w http.ResponseWriter, r *http.Request, useTicker bool, monitorCall func() interface{}) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}

	defer ws.Close()

	if useTicker {
		ticker := time.NewTicker(refreshRate)
		for range ticker.C {
			ws.WriteJSON(monitorCall())
		}
	} else {
		ws.WriteJSON(monitorCall())
	}

}
