package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting system monitor")
	monitor := new(SystemMonitor)
	monitor.GetHostInfo()
}
