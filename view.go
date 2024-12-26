package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/lipgloss"
)

func HostView(m *model) string {
	/*
		Main view of the CLI app, will contain things like RAM memory on host side,
		and if there is a GPU detected will create another bar.
	*/

	var tabView string

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
			gpuColorOptions)

		// handle gpu info in model.

		usedMemoryPercentage := m.gpuInfo.MemoryUsed / m.gpuInfo.MemoryTotal
		tabView += fmt.Sprintf("\nGPU Usage:  %0.2fMiB/%0.1fMiB \n%s\n %s",
			m.gpuInfo.MemoryUsed,
			m.gpuInfo.MemoryTotal,
			gpuProgressBar.ViewAs(usedMemoryPercentage),
			m.gpuInfo.Name)
	}

	return tabView
}

func RemoteView(m *model, tabIndex int) string {

	var tabView string

	tabView = "Nothing"

	return tabView

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

	for index, identifier := range m.tabs {

		// if its host tab just kill it early
		if m.activeTabIndex == 0 {
			tabView = HostView(&m)
			break
		}

		if index == m.activeTabIndex {
			tabView = RemoteView(&m, index)
		}

		if identifier == "NEW" {
			// Placeholder for future content
			tabView = "No Other System Running SOURUS Discovered On Network"
		}
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
