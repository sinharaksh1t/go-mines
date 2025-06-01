package main

import "github.com/charmbracelet/lipgloss"

var (
	mainColor = lipgloss.Color("#825DF2")

	cursorStyle            lipgloss.Style
	selectedFocusedStyle   lipgloss.Style
	selectedUnfocusedStyle lipgloss.Style
)

func initStyles() {
	cursorStyle = lipgloss.NewStyle().Background(mainColor).Foreground(lipgloss.Color("#FFFFFF"))
	selectedFocusedStyle = lipgloss.NewStyle().Background(lipgloss.Color("#FFB703" /* amber/gold */)).Foreground(lipgloss.Color("#FFFFFF"))
	selectedUnfocusedStyle = lipgloss.NewStyle().Background(lipgloss.Color("#3FC1C9" /* teal */)).Foreground(lipgloss.Color("#FFFFFF"))
}
