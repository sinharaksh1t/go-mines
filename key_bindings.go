package main

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

var (
	// Quit keys
	keyQuit = key.NewBinding(key.WithKeys(tea.KeyCtrlC.String(), "q"))

	// Direction Keys
	keyUp    = key.NewBinding(key.WithKeys(tea.KeyUp.String(), "k"))
	keyDown  = key.NewBinding(key.WithKeys(tea.KeyDown.String(), "j"))
	keyLeft  = key.NewBinding(key.WithKeys(tea.KeyLeft.String(), "h"))
	keyRight = key.NewBinding(key.WithKeys(tea.KeyRight.String(), "l"))

	// Selection keys
	keySelect = key.NewBinding(key.WithKeys(tea.KeySpace.String(), tea.KeyEnter.String()))
)
