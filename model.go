package main

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type cursorPos struct {
	row int
	col int
}

type cursorState struct {
	selected bool
}

type gridSize struct {
	numRows int
	numCols int
}

type model struct {
	cursor   cursorPos                 // The current position of the cursor
	grid     map[cursorPos]cursorState // This is the entire playable grid. With each block having certain properties (only `selectedStyle` for now)
	gridSize gridSize                  // Size of the playable grid
}

func newModel() *model {
	return &model{
		cursor: cursorPos{
			row: 0,
			col: 0,
		},
		grid: make(map[cursorPos]cursorState),
		gridSize: gridSize{
			// Let's take this as the default size of our playable grid
			numRows: 16,
			numCols: 16,
		},
	}
}

func (m *model) Init() tea.Cmd {
	// Returning `nil` means "no I/O right now, please."
	return nil
}

func (m *model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keyQuit):
			return m, tea.Quit

		case key.Matches(msg, keyUp):
			if m.cursor.row == 0 {
				m.cursor.row = m.gridSize.numRows - 1
			} else {
				m.cursor.row--
			}

		case key.Matches(msg, keyDown):
			if m.cursor.row == m.gridSize.numRows-1 {
				m.cursor.row = 0
			} else {
				m.cursor.row++
			}

		case key.Matches(msg, keyLeft):
			if m.cursor.col == 0 {
				m.cursor.col = m.gridSize.numCols - 1
			} else {
				m.cursor.col--
			}

		case key.Matches(msg, keyRight):
			if m.cursor.col == m.gridSize.numCols-1 {
				m.cursor.col = 0
			} else {
				m.cursor.col++
			}

		case key.Matches(msg, keySelect):
			currPos := cursorPos{
				row: m.cursor.row,
				col: m.cursor.col,
			}

			m.grid[currPos] = cursorState{
				selected: !m.grid[currPos].selected,
			}
		}

	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m *model) View() string {
	// Header
	view := "\n"
	header := "Play Mines\n"
	underline := "----------\n\n"
	movement := "Movement: \u2190 (h), \u2191 (k), \u2192 (l), or \u2193 (j)\n"
	selection := "Selection/Deselection: Space \" \"\n"

	helper := movement + selection
	view += header + underline + helper + "\n\n"

	for row := range m.gridSize.numRows {
		for col := range m.gridSize.numCols {
			focused := m.cursor.row == row && m.cursor.col == col
			selected := m.grid[cursorPos{row, col}].selected
			currElement := "0"
			if selected {
				currElement = "1"
			}

			switch {
			case focused && selected:
				view += selectedFocusedStyle.Render(currElement)
			case focused:
				view += cursorStyle.Render(currElement)
			case selected:
				view += selectedUnfocusedStyle.Render(currElement)
			default:
				view += currElement
			}

			view += " "
		}
		view += "\n"
	}

	view += "\nPress q or ctrl+c to quit"

	return view
}
