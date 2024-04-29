package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// NewTestModel().View()
	if _, err := tea.NewProgram(NewTestModel()).Run(); err != nil {
		panic(err)
	}
}
