package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

func main() {
	opts := []tea.ProgramOption{tea.WithAltScreen()}
	p := tea.NewProgram(newModel(), opts...)

	initStyles()

	if _, err := p.Run(); err != nil {
		log.Fatalln("Alas, unable to initialize the new model", err)
	}
}
