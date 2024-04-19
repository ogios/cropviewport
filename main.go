package cropviewport

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ogios/cropviewport/process"
)

type CropViewportModel struct {
	ANSITableList *process.ANSITableList
	KeyMap        map[string]func() tea.Cmd
	Sublines      []*process.SubLine
	Block         [4]int
}

func NewCropViewportModel() tea.Model {
	c := &CropViewportModel{}
	c.KeyMap = map[string]func() tea.Cmd{
		"j": func() tea.Cmd {
			c.NextLine(1)
			return nil
		},
		"k": func() tea.Cmd {
			c.PrevLine(1)
			return nil
		},
		"ctrl+d": func() tea.Cmd {
			c.NextLine(c.Block[3] / 2)
			return nil
		},
		"ctrl+u": func() tea.Cmd {
			c.PrevLine(c.Block[3] / 2)
			return nil
		},
		"h": func() tea.Cmd {
			c.PrevCol(1)
			return nil
		},
		"l": func() tea.Cmd {
			c.NextCol(1)
			return nil
		},
		"H": func() tea.Cmd {
			c.PrevCol(c.Block[2] / 2)
			return nil
		},
		"L": func() tea.Cmd {
			c.NextCol(c.Block[2] / 2)
			return nil
		},
	}
	return c
}

func (c *CropViewportModel) SetBlock(x, y, width, height int) {
	c.Block = [4]int{
		x, y, width, height,
	}
}

func (c *CropViewportModel) SetContent(s string) (*process.ANSITableList, []*process.SubLine) {
	t, l := process.ProcessContent(s)
	c.SetContentGivenData(t, l)
	return t, l
}

func (c *CropViewportModel) SetContentGivenData(tableList *process.ANSITableList, lines []*process.SubLine) {
	c.ANSITableList = tableList
	c.Sublines = lines
}

func (c *CropViewportModel) BackToTop() {
	c.Block[1] = 0
}

func (c *CropViewportModel) BackToLeft() {
	c.Block[0] = 0
}

func (c *CropViewportModel) PrevLine(step int) {
	c.Block[1] = max(c.Block[1]-step, 0)
}

func (c *CropViewportModel) NextLine(step int) {
	c.Block[1] = min(c.Block[1]+step, max(len(c.Sublines)-c.Block[3], 0))
}

func (c *CropViewportModel) PrevCol(step int) {
	c.Block[0] = max(c.Block[0]-step, 0)
}

func (c *CropViewportModel) NextCol(step int) {
	c.Block[0] += step
}

// * returns nil
func (c *CropViewportModel) Init() tea.Cmd {
	return nil
}

func (c *CropViewportModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := []tea.Cmd{}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		handler, ok := c.KeyMap[msg.String()]
		if ok {
			cmds = append(cmds, handler())
		}
	}
	return c, tea.Batch(cmds...)
}

// const NO_CONTENT = "No content available"
// var NO_CONTENT_TABLE, NO_CONTENT_SUBLINES = process.ProcessContent("No content available")
var NO_CONTENT_TABLE, NO_CONTENT_SUBLINES = process.ProcessContent("No content available")

func (c *CropViewportModel) View() string {
	if c.ANSITableList == nil || c.Sublines == nil {
		return process.CropView(NO_CONTENT_TABLE, NO_CONTENT_SUBLINES, c.Block[0], c.Block[1], c.Block[2], c.Block[3])
	}
	return process.CropView(c.ANSITableList, c.Sublines, c.Block[0], c.Block[1], c.Block[2], c.Block[3])
}
