package main

import tea "github.com/charmbracelet/bubbletea"

func main() {
	// fmt.Println(NewTestModel().View())
	if _, err := tea.NewProgram(NewTestModel()).Run(); err != nil {
		panic(err)
	}
}
