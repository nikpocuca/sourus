package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	// handle configs
	colorOptions := progress.WithGradient("#00A5BF", "#BF008F")
	p := progress.New(colorOptions)

	// check if a GPU is available.
	gpuCall := exec.Command("nvidia-smi")
	nvidiaCall := false
	var nvidiaInfoGPU *GPUInfo

	if gpuCall.Run() == nil {
		nvidiaCall = true

		infoGPU, err := monitorGPU()
		if err != nil {
			fmt.Println("Error monitoring GPU:", err)
			os.Exit(1)
		}

		nvidiaInfoGPU = &infoGPU
	}

	// get preliminary
	cpuLoads, _ := getCPULoads()

	// declare model
	m := model{
		progress:       p,
		percent:        0,
		usedMemoryGB:   0.0,
		totalMemoryGB:  0.0,
		coreLoad:       cpuLoads,
		activeTabIndex: 0,
		tabs:           []string{"HOST", "NEW"},
		gpuDetected:    nvidiaCall,
		gpuInfo:        nvidiaInfoGPU,
	}

	// start cli program
	_, err := tea.NewProgram(m).Run()

	if err != nil {
		fmt.Println("Error with progress program.")
	}
}
