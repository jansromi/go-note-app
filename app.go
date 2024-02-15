package main

import (
	// tui "tui_example/tui"
	core "go-note-app/core"
	// tea "github.com/charmbracelet/bubbletea"
)

func main() {
	items := core.NewItems()
	items.AddItem("Test Item", "This is a test item", "Lorem ipsum dolor sit amet")
	items.AddItem("Test Item 2", "This is a test item", "Lorem ipsum")
	items.PrintItems()
	// m := tui.NewListPage()
	// p := tea.NewProgram(m, tea.WithAltScreen())

	// _, error := p.Run()
	// if error != nil {
	// 	panic(error)
	// }

}
