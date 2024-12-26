package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

// getCPULoads returns the CPU load percentages for each core.
func getCPULoads() ([]float64, error) {
	// Specify the interval to calculate CPU usage (e.g., 1 second).
	interval := time.Millisecond * 750

	// Fetch CPU usage for each core.
	percentages, err := cpu.Percent(interval, true)
	if err != nil {
		return nil, err
	}

	return percentages, nil
}

func monitorRam() (float64, float64, float64) {
	virtualMemory, errorMemory := mem.VirtualMemory()
	if errorMemory != nil {
		fmt.Println("Cannot ready virtual memory")
		return -1.0, -1.0, -1.0
	}

	usedMemory, totalMemory := virtualMemory.Used/1024/1024, virtualMemory.Total/1024/1024

	return float64(usedMemory), float64(totalMemory), virtualMemory.UsedPercent
}

func monitorGPU() (GPUInfo, error) {
	cmd := exec.Command("nvidia-smi", "--query-gpu=name,temperature.gpu,utilization.gpu,utilization.memory,memory.total,memory.free,memory.used", "--format=csv,noheader,nounits")
	output, err := cmd.Output()
	if err != nil {
		return GPUInfo{}, err
	}

	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	var gpuInfo GPUInfo

	if scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ", ")
		if len(fields) == 7 {
			gpuInfo.Name = fields[0]

			// Parse numeric fields
			gpuInfo.Temperature, _ = strconv.ParseFloat(fields[1], 64)
			gpuInfo.UtilizationGPU, _ = strconv.ParseFloat(fields[2], 64)
			gpuInfo.UtilizationMem, _ = strconv.ParseFloat(fields[3], 64)
			gpuInfo.MemoryTotal, _ = strconv.ParseFloat(fields[4], 64)
			gpuInfo.MemoryFree, _ = strconv.ParseFloat(fields[5], 64)
			gpuInfo.MemoryUsed, _ = strconv.ParseFloat(fields[6], 64)
		}
	}

	if err := scanner.Err(); err != nil {
		return GPUInfo{}, err
	}

	return gpuInfo, nil
}
