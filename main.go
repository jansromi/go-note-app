package main

import (
	// tui "tui_example/tui"
	core "go-note-app/core"
	// tea "github.com/charmbracelet/bubbletea"
)

func main() {
	app := core.NewApp()
	app.AddItem("title", "desc", "content")
	// m := tui.NewListPage()
	// p := tea.NewProgram(m, tea.WithAltScreen())

	// _, error := p.Run()
	// if error != nil {
	// 	panic(error)
	// }

}
