package main

import "github.com/charmbracelet/bubbles/progress"

type GPUInfo struct {
	Name           string
	Temperature    float64
	UtilizationGPU float64
	UtilizationMem float64
	MemoryTotal    float64
	MemoryFree     float64
	MemoryUsed     float64
}

type model struct {
	/*
		Track Host Information
			- cpu core load
			- total RAM usage
			- tabs on the view
	*/
	progress       progress.Model
	percent        float64
	totalMemoryGB  float64
	usedMemoryGB   float64
	coreLoad       []float64
	finished       bool
	activeTabIndex int
	tabs           []string

	/*
		Track GPU Information
			- Temperature
			- Utilization
			- Memory

		gpuInfo could be nil if a gpu isnt detected
	*/
	gpuDetected bool
	gpuInfo     *GPUInfo
}
