package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ogios/clipviewport"
)

const (
	HEIGHT = 6
	WIDTH  = HEIGHT * 3
)

var (
	Blue        = "#0066ff"
	BgBlueStyle = lipgloss.NewStyle().Background(lipgloss.Color(Blue))
	FgBlueStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(Blue))
	BorderStyle = lipgloss.NewStyle().
			Width(WIDTH).Height(HEIGHT).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#b31a66"))

	CONTENTS = func() []string {
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
		lines := make([]string, len(ss))
		for i, v := range ss {
			s := v + BgBlueStyle.Render(v) + FgBlueStyle.Render(v) + v
			if i < len(ss)-1 {
				s += "\n"
			}
			lines[i] = s
		}
		return lines
	}()
)

type TestViewModel struct {
	ClipViewModel tea.Model
	Cacher        []*ContentData
	CurrentIndex  int
}

func NewTestModel() tea.Model {
	t := &TestViewModel{
		Cacher: make([]*ContentData, len(CONTENTS)),
	}
	for i, v := range CONTENTS {
		t.Cacher[i] = &ContentData{
			Raw: v,
		}
	}
	clip := clipviewport.NewClipViewportModel().(*clipviewport.ClipViewportModel)
	clip.SetBlock(0, 0, WIDTH, HEIGHT)
	t.ClipViewModel = clip
	t.SetContent(0)
	return t
}

func (t *TestViewModel) SetContent(index int) {
	t.CurrentIndex = index
	view := t.ClipViewModel.(*clipviewport.ClipViewportModel)
	c := t.Cacher[index]
	if c.Lines == nil || c.Table == nil {
		c.Table, c.Lines = view.SetContent(c.Raw)
	} else {
		view.SetContentGivenData(c.Table, c.Lines)
	}
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
		case "J":
			i := (t.CurrentIndex + 1 + len(t.Cacher)) % len(t.Cacher)
			t.SetContent(i)
		case "K":
			i := (t.CurrentIndex - 1 + len(t.Cacher)) % len(t.Cacher)
			t.SetContent(i)
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
	if _, err := tea.NewProgram(NewTestModel()).Run(); err != nil {
		panic(err)
	}
}
