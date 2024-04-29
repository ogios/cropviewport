package cropviewport

import (
	"slices"
	"strings"
	"unicode/utf8"

	"github.com/mattn/go-runewidth"
	process "github.com/ogios/ansisgr-process"
)

type SubLine struct {
	Data  *RuneDataList
	Bound [2]int
}

type RuneDataList struct {
	L          []process.BoundsStruct
	TotalWidth int
}

// init RuneData list given runes
//
// RuneDataList can only be set with this function, no more process allowed afterwards
func (r *RuneDataList) Init(s []rune) *RuneDataList {
	r.L = make([]process.BoundsStruct, len(s))
	visibleIndex := 0
	// for every rune, get its width, start and end index refers to the visible line
	// and save rune data into bytes
	for i, v := range s {
		bs := []byte{}
		bs = utf8.AppendRune(bs, v)
		w := runewidth.RuneWidth(v)
		r.L[i] = &RuneData{
			Byte:  slices.Clip(bs),
			Bound: [2]int{visibleIndex, visibleIndex + w},
		}
		r.TotalWidth += w
		visibleIndex += w
	}
	return r
}

type RuneData struct {
	Byte  []byte
	Bound [2]int // refers to the visible width
}

func (r *RuneData) GetBounds() [2]int {
	return r.Bound
}

const (
	line_break_rune   = '\n'
	line_break_string = "\n"
)

func SplitLines(raw string) []*SubLine {
	// split lines
	rawlines := strings.Split(raw, line_break_string)
	// process every line into runedata list and record line's bounds
	sublines := make([]*SubLine, len(rawlines))
	index := 0
	for i, v := range rawlines {
		data := (&RuneDataList{}).Init([]rune(v))
		lastIndex := index + len(data.L)
		sublines[i] = &SubLine{
			Bound: [2]int{index, lastIndex},
			Data:  data,
		}
		index = lastIndex + 1
	}
	return sublines
}

// separate ansi and normal string, separate lines into RuneData list
func ProcessContent(s string) (*process.ANSITableList, []*SubLine) {
	// separate ansi and normal string
	atablelist, raw := process.Extract(s)
	sublines := SplitLines(raw)
	return atablelist, sublines
}

var (
	SPACE_HODLER   = []byte(" ")
	SPACE_RUNEDATA = &RuneData{
		Byte: SPACE_HODLER,
	}
)

// extract certain area of the given lines, and render ansi sequence
func CropView(atablelist *process.ANSITableList, lines []*SubLine, x, y, width, height int) string {
	// get visible lines
	lines = process.SliceFrom(lines, y, y+height)
	// clip every visible line and add ansi
	var buf strings.Builder
	// not sure if this is necessary, it only considers the worst case of raw string(no ansi sequence)
	// if there are ansi sequences, the buf cap may still needs to grow when calling `write`
	buf.Grow((width + 1) * height)

	// lines
	for lineIndex, sl := range lines {
		// if x is within the width of line
		if sl.Data.TotalWidth-1 >= x {
			// (x) for range every rune and count width
			// (✓) binary search for a range of rune
			var start, end int
			temp := process.Search(sl.Data.L, x)
			start = temp[0]
			temp = process.Search(sl.Data.L, x+width-1)
			if len(temp) > 0 {
				end = temp[0]
			}
			lineRunes := make([]process.BoundsStruct, end-start+1)
			copy(lineRunes, sl.Data.L[start:end+1])
			// check for first rune, if width over 1 (max 2), replace to SPACE_RUNEDATA
			if lineRunes[0].GetBounds()[0] < x {
				lineRunes[0] = SPACE_RUNEDATA
			}
			// check for last rune, if not changed to SPACE_RUNEDATA and width over 1 (max 2), replace to SPACE_RUNEDATA
			if endRune := lineRunes[len(lineRunes)-1]; endRune != SPACE_RUNEDATA && endRune.GetBounds()[1] > x+width {
				lineRunes[len(lineRunes)-1] = SPACE_RUNEDATA
			}
			var lineBuf strings.Builder
			lineBuf.Grow(len(lineRunes) * 3)
			for _, bs := range lineRunes {
				lineBuf.Write(bs.(*RuneData).Byte)
			}
			buf.Write(process.Render(atablelist, lineBuf.String(), sl.Bound[0]+start))

		}
		// line break
		if lineIndex < len(lines)-1 {
			buf.WriteRune(line_break_rune)
		}
	}
	return buf.String()
}
