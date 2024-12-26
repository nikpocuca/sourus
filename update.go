package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

type tickMsg struct{}
type finishedMsg struct{}

func tick() tea.Cmd {
	return tea.Tick(time.Second/20, func(time.Time) tea.Msg {
		return tickMsg{}
	})
}

func finishedTick() tea.Cmd {
	return func() tea.Msg {
		// Sleep to allow progress bar to update properly
		time.Sleep(time.Millisecond * 500)
		return finishedMsg{}
	}
}

func (m model) Init() tea.Cmd {
	return tick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "tab":
			// Cycle through tabs when "TAB" is pressed
			m.activeTabIndex = (m.activeTabIndex + 1) % len(m.tabs)
		}
	case tickMsg:
		usedMemoryMB, totalMemoryMB, percentageMemory := monitorRam()

		m.percent = percentageMemory / 100.0
		m.usedMemoryGB = usedMemoryMB / 1000
		m.totalMemoryGB = totalMemoryMB / 1000

		cpuLoads, _ := getCPULoads()
		m.coreLoad = cpuLoads

		if m.percent >= 1.0 {
			if !m.finished {
				m.percent = 1.0
				cmd := m.progress.SetPercent(m.percent)
				m.finished = true
				return m, tea.Batch(cmd, finishedTick())
			}
		}
		cmd := m.progress.SetPercent(m.percent)

		gpuInfo, err := monitorGPU()

		if err != nil {
			fmt.Println("Could not update GPU Information, error called")
			fmt.Println(err)
			return m, tea.Quit
		}

		m.gpuInfo = gpuInfo

		return m, tea.Batch(cmd, tick())
	case finishedMsg:
		return m, tea.Quit
	}

	var cmd tea.Cmd
	var progressModel tea.Model
	progressModel, cmd = m.progress.Update(msg)
	m.progress = progressModel.(progress.Model) // Type assertion
	return m, cmd
}
