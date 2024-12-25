package main

import "github.com/charmbracelet/bubbles/progress"

type model struct {
	progress       progress.Model
	percent        float64
	totalMemoryGB  float64
	usedMemoryGB   float64
	coreLoad       []float64
	finished       bool
	activeTabIndex int // Track active tab index
	tabs           []string
	gpuDetected    bool
}
