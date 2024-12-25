package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/mem"
)



func monitorRam()  (float64, float64, float64){
    virtualMemory, errorMemory := mem.VirtualMemory()
    if errorMemory != nil {
        fmt.Println("Cannot ready virtual memory")
        return -1.0, -1.0, -1.0
    }

    usedMemory, totalMemory := virtualMemory.Used/1024/1024, virtualMemory.Total/1024/1024

    return float64(usedMemory), float64(totalMemory), virtualMemory.UsedPercent
}
