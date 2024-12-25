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
        // you need this added in so the progress bar update actually reaches 100
        time.Sleep(time.Millisecond * 500)
        return finishedMsg{}
    }
}


type model struct {
	progress progress.Model
	percent  float64
	totalMemoryGB float64
	usedMemoryGB float64
	finished bool
}

func (m model) Init() tea.Cmd {
	return tick()
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
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
    message := "Monitoring Press 'q' to quit \n"

    if m.finished {
        message = "DANGER, one or more of your devices are at maximum usage!"
    }

    maxIterations := m.totalMemoryGB // Define the maximum number of iterations
    currentIteration := m.usedMemoryGB

	return fmt.Sprintf(" Sauros \n Host Memory %s %0.2fG/%0.1fG\n%s\n \n",
	   m.progress.View(),
    	currentIteration,
		maxIterations,
        message,
	   )
}

func main() {

	colorOptions := progress.WithGradient("#00A5BF", "#BF008F")
	p := progress.New(colorOptions)

	m := model{
		progress: p,
		percent:  0,
		usedMemoryGB: 0.0,
		totalMemoryGB: 0.0,
		finished: false,
	}

	_, err := tea.NewProgram(m).Run()

	if err != nil {
	   fmt.Println("Error with progress program.")
	}
}
