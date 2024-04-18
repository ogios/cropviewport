package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ogios/clipviewport"
)

const (
	HEIGHT  = 6
	WIDTH   = HEIGHT * 2
	CONTENT = `Hello, World! Hello, World! Hello, World! Hello, World!
你好，世界！ 你好，世界！ 你好，世界！ 你好，世界！
こんにちは、世界よ！ こんにちは、世界よ！ こんにちは、世界よ！ こんにちは、世界よ！
مرحباً أيها العالم! مرحباً أيها العالم! مرحباً أيها العالم! مرحباً أيها العالم!
Bonjour à tous ! Bonjour à tous ! Bonjour à tous ! Bonjour à tous !
Γεια σου, κόσμε! Γεια σου, κόσμε! Γεια σου, κόσμε! Γεια σου, κόσμε!
안녕하세요, 세상 여러분! 안녕하세요, 세상 여러분! 안녕하세요, 세상 여러분! 안녕하세요, 세상 여러분!
Привет, мир! Привет, мир! Привет, мир! Привет, мир!
Hej, världen! Hej, världen! Hej, världen! Hej, världen!`
)

var BorderStyle = lipgloss.NewStyle().
	Width(WIDTH).Height(HEIGHT).
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("#b31a66"))

type TestViewModel struct {
	ClipViewModel tea.Model
}

func NewTestModel() tea.Model {
	t := &TestViewModel{}
	clip := clipviewport.NewClipViewportModel().(*clipviewport.ClipViewportModel)
	clip.SetBlock(0, 0, WIDTH, HEIGHT)
	clip.SetContent(CONTENT)
	t.ClipViewModel = clip
	return t
}

func (t *TestViewModel) Init() tea.Cmd {
	return nil
}

func (t *TestViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			fallthrough
		case "ctrl+c":
			return t, tea.Quit
		}
	}

	m, cmd := t.ClipViewModel.Update(msg)
	t.ClipViewModel = m
	return t, cmd
}

func (t *TestViewModel) View() string {
	return BorderStyle.Render(t.ClipViewModel.View())
}

func main() {
	// NewTestModel().View()
	if _, err := tea.NewProgram(NewTestModel()).Run(); err != nil {
		panic(err)
	}
}
