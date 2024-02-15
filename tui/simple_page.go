package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// MODEL DATA

type SimplePage struct{ text string }

func NewSimplePage(text string) SimplePage {
	return SimplePage{text: text}
}

func (s SimplePage) Init() tea.Cmd { return nil }

// VIEW

func (s SimplePage) View() string {
	textLen := len(s.text)
	topAndBottomBar := strings.Repeat("*", textLen+4)
	return fmt.Sprintf(
		"%s\n* %s *\n%s\n\nPress Ctrl+C to exit",
		topAndBottomBar, s.text, topAndBottomBar,
	)
}

// UPDATE

func (s SimplePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return s, tea.Quit

		default:
			key := msg.String()
			text := fmt.Sprintf("You pressed '%s'", key)
			s = NewSimplePage(text)
			return s, nil
		}

	}
	return s, nil
}
