package main

import (
	"fmt"
	"time"
	"os/exec"

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
	coreLoad       []float64
	finished       bool
	activeTabIndex int // Track active tab index
	tabs           []string
	gpuDetected    bool
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

		cpuLoads, _ :=  getCPULoads()
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

	boxStyleHeader := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("12")).
	Align(lipgloss.Center).
	Width(60)

	// Define a style for the box
	boxStyle := lipgloss.NewStyle().
		// Padding().
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
		memoryUsage := fmt.Sprintf("Host Memory: %0.2fG/%0.1fG\n%s\n", m.usedMemoryGB, m.totalMemoryGB, m.progress.View())

		// Construct Core Load Grid with Progress Bars
		gridRows := ""
		colsPerRow := 2 // Number of columns in the grid
		for i, load := range m.coreLoad {


      		colorOptions := progress.WithGradient("#B0FF00", "#FF0F00")
			coreProgressBar := progress.New(
			             colorOptions,
			             progress.WithWidth(10),
						 progress.WithoutPercentage(),
			             )

			// Set the percentage for the core's load
			barView := coreProgressBar.ViewAs(load / 100.0)

			gridRows += fmt.Sprintf("%5.1f%% %s", load, barView)

			if (i+1)%colsPerRow == 0 || i == len(m.coreLoad)-1 {
				gridRows += "\n" // New row after reaching column limit
			} else {
				gridRows += "\t" // Tab space between columns
			}
		}

		tabView = fmt.Sprintf("%s\n\nCore Loads\n%s", memoryUsage, gridRows)

		if m.gpuDetected {
    		gpuColorOptions := progress.WithGradient("#FFFB00", "#FF0084")
            gpuProgressBar := progress.New(
			             gpuColorOptions,)

            tabView += fmt.Sprintf("GPU Usage\n%s", gpuProgressBar.View())
		}

	case 1:
		// Placeholder for future content
		tabView = "No Other System Running SOURUS Discovered On Network"
	}

	// Construct the final view with the header and tab content
	content := fmt.Sprintf(
		"%s\n%s\n\n",
		boxStyleHeader.Render(header),
		boxStyle.Render(tabView),
	)

	// Add instructions below the box
	instructions := instructionsStyle.Render("Press 'TAB' to switch tabs\nPress 'q' to quit")

	// Combine everything
	return content + instructions
}



func main() {
    cpuLoads, _ := getCPULoads()

	colorOptions := progress.WithGradient("#00A5BF", "#BF008F")
	p := progress.New(colorOptions)

    // check if a GPU is available.
    gpuCall := exec.Command("nvidia-smi")
    nvidiaCall := false
	if gpuCall.Run() == nil {
		nvidiaCall = true
	}

	m := model{
		progress:      p,
		percent:       0,
		usedMemoryGB:  0.0,
		totalMemoryGB: 0.0,
		coreLoad:      cpuLoads,
		finished:      false,
		activeTabIndex: 0,
		tabs:           []string{"Memory Usage", "Other Information"},
		gpuDetected:     nvidiaCall,
	}

	_, err := tea.NewProgram(m).Run()

	if err != nil {
		fmt.Println("Error with progress program.")
	}
}
