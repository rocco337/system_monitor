package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "127.0.0.1:8080", "http service address")
var monitor = SystemMonitor{}
var processRefreshRate = 2 * time.Second

func main() {
	fmt.Println("Starting system monitor")
	//monitor := new(SystemMonitor)

	http.HandleFunc("/hostInfo", serveHostInfo)
	http.HandleFunc("/processes", serveProcesses)
	http.HandleFunc("/cpuusage", serveCpuusage)

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

func serveHostInfo(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}

	defer ws.Close()

	ws.WriteJSON(monitor.GetHostInfo())
}

func serveProcesses(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}

	defer ws.Close()

	ticker := time.NewTicker(processRefreshRate)

	for range ticker.C {
		ws.WriteJSON(monitor.GetProcesses())
	}
}

func serveCpuusage(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}

	defer ws.Close()

	ticker := time.NewTicker(processRefreshRate)

	for range ticker.C {
		ws.WriteJSON(monitor.GetCPUThreadsUsage())
	}
}
