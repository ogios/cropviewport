package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ogios/cropviewport"
)

const (
	HEIGHT = 6
	WIDTH  = HEIGHT * 2
)

var (
	Blue        = "#0066ff"
	BgBlueStyle = lipgloss.NewStyle().Background(lipgloss.Color(Blue))
	FgBlueStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(Blue))
	BorderStyle = lipgloss.NewStyle().
			Width(WIDTH).Height(HEIGHT).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#b31a66"))

	CONTENT = func() string {
		var buf strings.Builder
		ss := []string{
			"Hello, World!",
			"你好，世界！",
			"こんにちは、世界よ！",
			// "مرحباً أيها العالم!",
			"Bonjour à tous !",
			"Γεια σου, κόσμε!",
			"안녕하세요, 세상 여러분!",
			"Привет, мир!",
			"Hej, världen!",
		}
		for i, v := range ss {
			buf.WriteString(v + BgBlueStyle.Render(v) + FgBlueStyle.Render(v) + v)
			if i < len(ss)-1 {
				buf.WriteString("\n")
			}
		}
		return buf.String()
	}()
)

type TestViewModel struct {
	CropViewModel tea.Model
}

func NewTestModel() tea.Model {
	t := &TestViewModel{}
	crop := cropviewport.NewCropViewportModel().(*cropviewport.CropViewportModel)
	crop.SetBlock(0, 0, WIDTH, HEIGHT)
	crop.SetContent(CONTENT)
	t.CropViewModel = crop
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

	m, cmd := t.CropViewModel.Update(msg)
	t.CropViewModel = m
	return t, cmd
}

func (t *TestViewModel) View() string {
	return BorderStyle.Render(t.CropViewModel.View())
}

func main() {
	fmt.Println(CONTENT)
	if _, err := tea.NewProgram(NewTestModel()).Run(); err != nil {
		panic(err)
	}
}
