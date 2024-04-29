package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/alecthomas/chroma/v2/quick"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	"github.com/ogios/cropviewport/process"
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

func init() {
	_c1, err := os.ReadFile("./code1.txt")
	if err != nil {
		panic(err)
	}
	code1 = string(_c1)
	_c2, err := os.ReadFile("./code2.txt")
	if err != nil {
		panic(err)
	}
	code2 = string(_c2)
}

func diffContent() (*process.ANSITableList, []*process.SubLine, error) {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(string(code1), string(code2), true)
	diffs = dmp.DiffCleanupSemantic(diffs)
	diffs = dmp.DiffCleanupEfficiency(diffs)
	dc := diffcontext.New()
	dc.AddDiffs(diffs)
	records := getLinesRecords(dc)
	err := highlightCodes(dc)
	if err != nil {
		return nil, nil, err
	}

	at, sl := highlightDiffLines(dc, records)
	return at, sl, err
}

const (
	line_break = '\n'
	tab        = "\t"
)

var empty_records = make([][3]int, 0)

func getLinesRecords(dc *diffcontext.DiffConstractor) [][3]int {
	mls := dc.GetMixedLines()
	if len(mls) == 0 {
		return empty_records
	}
	records := make([][3]int, 0)
	var n int
	for _, ml := range mls {
		length := len(ml.Data) + strings.Count(string(ml.Data), tab)*(4-1)
		switch ml.State {
		case diffmatchpatch.DiffInsert:
			records = append(records, [3]int{int(diffmatchpatch.DiffInsert), n, n + length})
		case diffmatchpatch.DiffDelete:
			records = append(records, [3]int{int(diffmatchpatch.DiffDelete), n, n + length})
		}
		n += length + 1
	}
	return records
}

func highlightCodes(dc *diffcontext.DiffConstractor) error {
	c1, err := highlight(code1)
	if err != nil {
		return err
	}
	linesC1 := strings.Split(c1, "\n")

	c2, err := highlight(code2)
	if err != nil {
		return err
	}
	linesC2 := strings.Split(c2, "\n")
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

func highlightDiffLines(dc *diffcontext.DiffConstractor, records [][3]int) (*process.ANSITableList, []*process.SubLine) {
	at, sl := process.ProcessContent(dc.GetMixed())
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
	err := quick.Highlight(buf, string(c), "go", "terminal16m", "catppuccin-mocha")
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
