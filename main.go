package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	newFileInput           textinput.Model
	createFileInputVisible bool
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit
		case "ctrl+n":
			m.createFileInputVisible = true
			return m, nil
		}

	}

	return m, nil
}

func (m model) View() string {

	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		PaddingLeft(2).
		PaddingRight(2)

	welcome := style.Render("Welcome to Navnote")

	help := "Ctrl+N:New File - Ctrl+L: List - Esc: Back/Save - Ctrl+S: Save - Ctrl+Q: Quit"

	view := ""
	if m.createFileInputVisible {
		view = m.newFileInput.View()
	}
	return fmt.Sprintf("\n%s\n\n%s\n\n%s", welcome, view, help)
}

func initializeMode() model {
	//starting new file input
	ti := textinput.New()
	ti.Placeholder = "What would you like to call it ?"
	ti.Focus()
	ti.CharLimit = 156

	return model{
		newFileInput:           ti,
		createFileInputVisible: false,
	}
}

func main() {
	p := tea.NewProgram(initializeMode())
	if _, err := p.Run(); err != nil {
		fmt.Printf("there is an error : %v", err)
		os.Exit(1)
	}

}
