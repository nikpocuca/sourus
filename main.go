package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	colorTheme := GenerateTheme()

	// handle configs

	colorOptions := progress.WithGradient(
		colorTheme.HostMemoryGradientColors.Left,
		colorTheme.HostMemoryGradientColors.Right,
	)

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
		colorTheme:     colorTheme,
	}

	// start cli program
	_, err := tea.NewProgram(m).Run()

	if err != nil {
		fmt.Println("Error with progress program.")
	}
}
