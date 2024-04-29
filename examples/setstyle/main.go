package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	"github.com/ogios/cropviewport"
)

const (
	HEIGHT = 20
	WIDTH  = HEIGHT * 2
)

var BorderStyle = lipgloss.NewStyle().
	Width(WIDTH).Height(HEIGHT).
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("#b31a66"))

var CONTENT string

type TestViewModel struct {
	CropViewModel tea.Model
}

func NewTestModel() tea.Model {
	_c, err := os.ReadFile("./test")
	if err != nil {
		panic(err)
	}
	CONTENT = string(_c)
	t := &TestViewModel{}
	crop := cropviewport.NewCropViewportModel().(*cropviewport.CropViewportModel)
	crop.SetBlock(0, 0, WIDTH, HEIGHT)
	al, sl := cropviewport.ProcessContent(CONTENT)
	al.SetStyle([]byte(fmt.Sprintf("%s%sm", termenv.CSI, termenv.RGBColor("#ff00ff").Sequence(true))), 7, 4*6-2)
	crop.SetContentGivenData(al, sl)
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
	// NewTestModel().View()
	if _, err := tea.NewProgram(NewTestModel()).Run(); err != nil {
		panic(err)
	}
}
