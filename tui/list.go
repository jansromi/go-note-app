package tui

import (
	core "go-note-app/core"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// KEYS FOR INPUT EVENTS
type inputKeyMap struct {
	Save   key.Binding
	Delete key.Binding
}

var inputEventKeys = inputKeyMap{
	Save: key.NewBinding(
		key.WithKeys("ctrl+s"),
		key.WithHelp("ctrl+s", "save"),
	),
	Delete: key.NewBinding(
		key.WithKeys("ctrl+d"),
		key.WithHelp("ctrl+d", "delete item"),
	),
}

// sets the view style
var docStyle = lipgloss.NewStyle().Margin(1, 2)

// data structure for list items
type item struct {
	id          int
	title, desc string
}

// getters for list items
func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

// model for list page
type model struct {
	list list.Model
	// a reference to the core app
	app core.App
}

func (m model) Init() tea.Cmd {
	return nil
}

// Handle input events
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "d" {
			deleteId := m.list.SelectedItem().(item).id
			m.app.DeleteItem(deleteId)
			m.list.SetItems(GetListItems(&m.app))
		}
		if msg.String() == "enter" {
			note := m.list.SelectedItem().(item)
			print(note.id)
		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// Render the list page
func (m model) View() string {
	return docStyle.Render(m.list.View())
}

// Get all list items from the core app
func GetListItems(noteapp *core.App) []list.Item {
	allItems := noteapp.GetItemsForListView()
	var listItems []list.Item

	for _, v := range allItems {
		listItems = append(listItems, item{id: v.ID, title: v.Title, desc: v.Desc})
	}
	return listItems
}

// Create a new list page
func NewListPage(noteapp *core.App) tea.Model {

	m := model{list: list.New(GetListItems(noteapp), list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Note App"
	keys := []key.Binding{inputEventKeys.Save, inputEventKeys.Delete}
	m.list.AdditionalShortHelpKeys = func() []key.Binding {
		return append([]key.Binding{}, keys...)
	}
	m.app = *noteapp

	return m
}
