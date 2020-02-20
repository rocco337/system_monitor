package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting system monitor")
	monitor := new(SystemMonitor)

	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		monitor.GetCpuThreadsUsage()
		fmt.Println("--------------------------")
	}
}
