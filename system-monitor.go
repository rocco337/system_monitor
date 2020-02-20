package main

import (
	"fmt"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/process"

	"time"
)

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

//GetHostInfo ...
func (m *SystemMonitor) GetHostInfo() *HostInfo {
	hostInfo := new(HostInfo)
	info, err := host.Info()
	if err == nil {
		hostInfo.Hostname = info.Hostname
		hostInfo.Uptime = time.Unix(int64(info.Uptime), 0)
		hostInfo.Procs = info.Procs
		hostInfo.OS = info.OS
	}

	hostInfo.PrintLine()
	return hostInfo
}

//HostInfo ...
type HostInfo struct {
	Hostname string
	Uptime   time.Time
	Procs    uint64
	OS       string
}

func (h *HostInfo) PrintLine() {
	fmt.Printf("%s - %d - %s \n", h.Hostname, h.Procs, h.OS)
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
