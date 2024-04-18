package main

import (
	"bytes"
	"fmt"

	"github.com/alecthomas/chroma/v2/quick"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ogios/clipviewport"
)

const (
	CONTENT = `package main

import (
	"errors"
	"fmt"
)

type meow string

type Aughhhhhhhh struct {
	Cat meow
}

func YesThisIsATestFunctionOrWhatElseItWouldBeAnyway(nothing Aughhhhhhhh) error {
	fmt.Println(nothing.Cat)
	return errors.New("Nah")
}

func main() {
	a := Aughhhhhhhh{
		Cat: ` + "`" + `I am the storm that is approaching
Provoking black clouds in isolation
I am reclaimer of my name
Born in flames, I have been blessed
My family crest is a demon of death

Forsakened, I am awakened
A phoenix's ash in dark divine
Descending misery
Destiny chasing time

Inherit the nightmare, surrounded by fate
Can't run away
Keep walking the line, between the light
Led astray

Through vacant halls I won't surrendеr
The truth revealеd in eyes of ember
We fight through fire and ice forever
Two souls once lost and now they remember

I am the storm that is approaching
Provoking black clouds in isolation
I am reclaimer of my name
Born in flames, I have been blessed
My family crest is a demon of death

Forsakened, I am awakened
A phoenix's ash in dark divine
Descending misery
Destiny chasing time

Disappear into the night
Lost shadows left behind
Obsession's pulling me
Fading, I've come to take what's mine

Lurking in the shadows under veil of night
Constellations of blood pirouette
Dancing through the graves of those who stand at my feet
Dreams of the black throne I keep on repeat

A derelict of dark summoned from the ashes
The puppet master congregates all the masses
Pulling strings, twisting minds as blades hit
You want this power? Then come try and take it

Beyond the tree
Fire burns
Secret love
Bloodline yearns

Dark minds embrace
Crimson joy
Does your dim heart
Heal or destroy?

Bury the light deep within
Cast aside, there's no coming home
We're burning chaos in the wind` + "`" + `,
	}
	err := YesThisIsATestFunctionOrWhatElseItWouldBeAnyway(a)
	if err != nil {
		panic(err)
	}
}
`
	HEIGHT = 20
	WIDTH  = HEIGHT * 2
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
	buf := &bytes.Buffer{}
	err := quick.Highlight(buf, CONTENT, "go", "terminal16m", "catppuccin-mocha")
	if err != nil {
		panic(err)
	}
	fmt.Println(buf.String())
	clip.SetContent(buf.String())
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
