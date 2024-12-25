package main

import (
	"fmt"
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
	if gpuCall.Run() == nil {
		nvidiaCall = true
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
		finished:       false,
		activeTabIndex: 0,
		tabs:           []string{"HOST", "NEW"},
		gpuDetected:    nvidiaCall,
	}

	// start cli program
	_, err := tea.NewProgram(m).Run()

	if err != nil {
		fmt.Println("Error with progress program.")
	}
}
