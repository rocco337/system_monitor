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

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/hostInfo", serveHostInfo)
	http.HandleFunc("/processes", serveProcesses)

	http.HandleFunc("/home.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "home.js")
	})

	log.Fatal(http.ListenAndServe(*addr, nil))

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
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

	for _ = range ticker.C {
		ws.WriteJSON(monitor.GetProcesses())
	}

}
