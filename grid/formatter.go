package grid

import (
	"bytes"
	"fmt"
)

type Formatter struct {
	Separator string
	// 0 = left, 1 = right
	Align    int
	Stringer func(interface{}) string
}

type alignStr func(w int) string

var align = []alignStr{
	func(w int) string {
		return fmt.Sprintf("%%-%ds", w)
	},
	func(w int) string {
		return fmt.Sprintf("%%%ds", w)
	},
}

func defaultStringer(i interface{}) string {
	return fmt.Sprint(i)
}

func (f Formatter) Format(g *Grid) string {
	stringer := f.Stringer
	if stringer == nil {
		stringer = defaultStringer
	}
	widths := make([]int, g.Size.X)
	strs := make([]string, g.Size.Area())
	for iter, pt, ok := g.Size.FromOrigin(); ok; pt, ok = iter.Next() {
		s := stringer(g.Get(pt))
		strs[iter.Idx()] = s
		if l := len([]rune(s)); l > widths[pt.X] {
			widths[pt.X] = l
		}
	}
	widthFmt := make([]string, len(widths))
	a := align[f.Align]
	for i, w := range widths {
		widthFmt[i] = a(w)
	}
	var buf bytes.Buffer
	for iter, pt, ok := g.Size.FromOrigin(); ok; pt, ok = iter.Next() {
		if pt.X == 0 {
			buf.WriteString("\n")
		} else {
			buf.WriteString(f.Separator)
		}
		buf.WriteString(fmt.Sprintf(widthFmt[pt.X], strs[iter.Idx()]))
	}
	return buf.String()
}
