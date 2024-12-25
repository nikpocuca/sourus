package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

type model struct {
	progress       progress.Model
	percent        float64
	totalMemoryGB  float64
	usedMemoryGB   float64
	finished       bool
	activeTabIndex int // Track active tab index
	tabs           []string
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

		if m.percent >= 1.0 {
			if !m.finished {
				m.percent = 1.0
				cmd := m.progress.SetPercent(m.percent)
				m.finished = true
				return m, tea.Batch(cmd, finishedTick())
			}
		}
		cmd := m.progress.SetPercent(m.percent)
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

func (m model) View() string {
	// Define a style for the header
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("15")).
		Align(lipgloss.Center)

	// Define a style for the box
	boxStyle := lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("12")).
		Align(lipgloss.Center).
		Width(60)

	// Define a style for the instructions
	instructionsStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("14"))

	// Header to display on each view
	header := headerStyle.Render("SOURUS MONITOR")

	// Content of the active tab
	var tabView string
	switch m.activeTabIndex {
	case 0:
		// Memory Tab
		tabView = fmt.Sprintf("Host Memory: %0.2fG/%0.1fG\n%s\n", m.usedMemoryGB, m.totalMemoryGB, m.progress.View())
	case 1:
		// Other tab (placeholder for future content)
		tabView = "No Other System Running SOURUS Discovered On Network"
	}

	// Construct the final view with the header and tab content
	content := fmt.Sprintf(
		"%s\n%s\n\n",
		header,
		boxStyle.Render(tabView),
	)

	// Add instructions below the box
	instructions := instructionsStyle.Render("Press 'TAB' to switch tabs\nPress 'q' to quit")

	// Combine everything
	return content + instructions
}


func main() {
	colorOptions := progress.WithGradient("#00A5BF", "#BF008F")
	p := progress.New(colorOptions)

	m := model{
		progress:      p,
		percent:       0,
		usedMemoryGB:  0.0,
		totalMemoryGB: 0.0,
		finished:      false,
		activeTabIndex: 0,
		tabs:           []string{"Memory Usage", "Other Information"},
	}

	_, err := tea.NewProgram(m).Run()

	if err != nil {
		fmt.Println("Error with progress program.")
	}
}
