package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/process"
)

func main() {
	fmt.Println("Starting system monitor")
	monitor := new(SystemMonitor)
	monitor.GetProcesses()
}

//SystemMonitor ...
type SystemMonitor struct{}

//GetProcesses ...
func (m *SystemMonitor) GetProcesses() []ProcessInfo {
	list := make([]ProcessInfo, 0)

	v, _ := process.Pids()

	for i := 0; i < len(v); i++ {
		proc, err := process.NewProcess(v[i])

		if err == nil {
			procInfo := ProcessInfo{}
			procInfo.Pid = proc.Pid
			procInfo.Name, _ = proc.Name()
			procInfo.Status, _ = proc.Status()
			procInfo.CPUPercent, _ = proc.CPUPercent()
			procInfo.MemPercent, _ = proc.MemoryPercent()
			procInfo.Username, _ = proc.Username()

			list = append(list, procInfo)
		}
	}

	// sort.Slice(list[:], func(i, j int) bool {
	// 	return list[i].CPUPercent > list[j].CPUPercent
	// })

	return list
}

//ProcessInfo ...
type ProcessInfo struct {
	Pid        int32
	Name       string
	Status     string
	CPUPercent float64
	MemPercent float32
	Username   string
	Uptime     time.Duration
}

//PrintLine ...
func (p *ProcessInfo) PrintLine() {
	fmt.Printf("%d - %s - %s - %f - %f - %s\n", p.Pid, p.Name, p.Status, p.CPUPercent, p.MemPercent, p.Username)
}
