package main

import (
	core "go-note-app/core"
	tui "go-note-app/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	app := core.NewApp()
	m := tui.NewListPage(app)
	p := tea.NewProgram(m, tea.WithAltScreen())

	_, error := p.Run()
	if error != nil {
		panic(error)
	}

}
