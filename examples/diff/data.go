package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/alecthomas/chroma/v2/quick"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	process "github.com/ogios/ansisgr-process"
	"github.com/ogios/cropviewport"
	"github.com/ogios/go-diffcontext"
	"github.com/sergi/go-diff/diffmatchpatch"
)

const (
	HEIGHT = 20
	WIDTH  = HEIGHT * 2
)

var BorderStyle = lipgloss.NewStyle().
	Width(WIDTH).Height(HEIGHT).
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("#b31a66"))

var code1, code2 string

// var (
// 	code1 = `11111
// 11111
// 11111
// 11111`
// 	code2 = `22222
// 11111
// 22222
// 11111`
// )

func init() {
	// _c1, err := os.ReadFile("./code1.txt")
	// _c2, err := os.ReadFile("./code2.txt")
	// _c1, _ := os.ReadFile("./test1")
	// _c2, _ := os.ReadFile("./test2")
	_c1, _ := os.ReadFile("./layout21")
	_c2, _ := os.ReadFile("./layout22")
	code1 = string(_c1)
	code2 = string(_c2)
}

func diffContent() (*process.ANSITableList, []*cropviewport.SubLine, error) {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(string(code1), string(code2), true)
	diffs = dmp.DiffCleanupSemantic(diffs)
	diffs = dmp.DiffCleanupEfficiency(diffs)
	dc := diffcontext.New()
	dc.AddDiffs(diffs)
	// mls, records := dc.GetMixedLinesAndStateRecord()
	_, records := dc.GetMixedLinesAndStateRecord()
	err := highlightCodes(dc)
	if err != nil {
		return nil, nil, err
	}

	at, sl := highlightDiffLines(dc.GetMixed(), records)
	return at, sl, err
}

func highlightCodes(dc *diffcontext.DiffConstractor) error {
	// c1, err := highlight(code1)
	// if err != nil {
	// 	return err
	// }
	// linesC1 := strings.Split(c1, "\n")
	linesC1 := strings.Split(code1, "\n")

	// c2, err := highlight(code2)
	// if err != nil {
	// 	return err
	// }
	// linesC2 := strings.Split(c2, "\n")
	linesC2 := strings.Split(code2, "\n")
	i1 := 0
	i2 := 0
	for _, dl := range dc.Lines {
		switch dl.State {
		case diffmatchpatch.DiffEqual:
			be := []byte(linesC1[i1])
			dl.Before, dl.After = be, be
			i1++
			i2++
		default:
			switch dl.State {
			case diffcontext.DiffChanged:
				be := []byte(linesC1[i1])
				af := []byte(linesC2[i2])
				dl.Before, dl.After = be, af
				i1++
				i2++
			case diffmatchpatch.DiffInsert:
				af := []byte(linesC2[i2])
				dl.After = af
				i2++
			case diffmatchpatch.DiffDelete:
				be := []byte(linesC1[i1])
				dl.Before = be
				i1++
			}
		}
	}
	return nil
}

var (
	redBG   = []byte(fmt.Sprintf("%s%sm", termenv.CSI, termenv.RGBColor("#991a1a").Sequence(true)))
	greenBG = []byte(fmt.Sprintf("%s%sm", termenv.CSI, termenv.RGBColor("#008033").Sequence(true)))
)

func highlightDiffLines(content string, records [][3]int) (*process.ANSITableList, []*cropviewport.SubLine) {
	at, sl := cropviewport.ProcessContent(content)
	for _, v := range records {
		var color []byte
		switch v[0] {
		case int(diffmatchpatch.DiffDelete):
			color = redBG
		case int(diffmatchpatch.DiffInsert):
			color = greenBG
		}
		at.SetStyle(color, v[1], v[2])
	}
	return at, sl
}

func highlight(c string) (string, error) {
	buf := new(bytes.Buffer)
	err := quick.Highlight(buf, string(c), "sum", "terminal16m", "catppuccin-mocha")
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
