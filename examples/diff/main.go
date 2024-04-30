package main

import "fmt"

func main() {
	fmt.Println(NewTestModel().View())
	// if _, err := tea.NewProgram(NewTestModel()).Run(); err != nil {
	// 	panic(err)
	// }
}
