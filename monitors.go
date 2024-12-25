package main

import (
	"fmt"
	"time"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/cpu"

)

// getCPULoads returns the CPU load percentages for each core.
func getCPULoads() ([]float64, error) {
	// Specify the interval to calculate CPU usage (e.g., 1 second).
	interval := time.Millisecond * 500

	// Fetch CPU usage for each core.
	percentages, err := cpu.Percent(interval, true)
	if err != nil {
		return nil, err
	}

	return percentages, nil
}


func monitorRam()  (float64, float64, float64){
    virtualMemory, errorMemory := mem.VirtualMemory()
    if errorMemory != nil {
        fmt.Println("Cannot ready virtual memory")
        return -1.0, -1.0, -1.0
    }

    usedMemory, totalMemory := virtualMemory.Used/1024/1024, virtualMemory.Total/1024/1024

    return float64(usedMemory), float64(totalMemory), virtualMemory.UsedPercent
}
