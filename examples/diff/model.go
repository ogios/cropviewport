package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ogios/cropviewport"
)

type TestViewModel struct {
	CropViewModel tea.Model
}

func NewTestModel() tea.Model {
	t := &TestViewModel{}
	crop := cropviewport.NewCropViewportModel().(*cropviewport.CropViewportModel)
	crop.SetBlock(0, 0, WIDTH, HEIGHT)
	at, sl, err := diffContent()
	if err != nil {
		panic(err)
	}
	crop.SetContentGivenData(at, sl)
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
