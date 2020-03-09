package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "0.0.0.0:8081", "http service address")
var deafultRefreshRate = 2 * time.Second

func main() {
	flag.Parse()
	fmt.Printf("Starting system monitor on address: %s \n", *addr)
	monitor := new(SystemMonitor)

	http.HandleFunc("/hostInfo", func(w http.ResponseWriter, r *http.Request) {
		serveJSON(w, r, false, deafultRefreshRate, func() interface{} {
			return monitor.GetHostInfo()
		})
	})

	http.HandleFunc("/processes", func(w http.ResponseWriter, r *http.Request) {
		serveJSON(w, r, true, deafultRefreshRate, func() interface{} {
			return monitor.GetProcesses()
		})
	})

	http.HandleFunc("/cpuusage", func(w http.ResponseWriter, r *http.Request) {
		serveJSON(w, r, true, deafultRefreshRate, func() interface{} {
			return monitor.GetCPUThreadsUsage()
		})
	})

	http.HandleFunc("/memorystat", func(w http.ResponseWriter, r *http.Request) {
		serveJSON(w, r, true, deafultRefreshRate, func() interface{} {
			return monitor.GetMemoryUsage()
		})
	})

	http.HandleFunc("/temperatures", func(w http.ResponseWriter, r *http.Request) {
		serveJSON(w, r, true, 1*time.Second, func() interface{} {
			return GetTemperatures()
		})
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serveFile(w, r, "index.html")
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

func serveJSON(w http.ResponseWriter, r *http.Request, useTicker bool, refreshRate time.Duration, monitorCall func() interface{}) {
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
			fmt.Printf("Serving: %s \n", r.RequestURI)
			err := ws.WriteJSON(monitorCall())
			if err != nil {
				fmt.Println(err)
				break
			}
		}
	} else {
		ws.WriteJSON(monitorCall())
	}

}
