package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"github.com/shirou/gopsutil/v3/mem"
)

func getFanSpeed() (string, error) {
	// Path to fan speed file, this path may vary based on your system configuration
	// You can find the fan speed by checking /sys/class/hwmon/ or using the `sensors` command.
	filePath := "/sys/class/hwmon/hwmon0/fan1_input"

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(data)), nil
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
