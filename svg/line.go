package bolt

import (
	"strconv"

	"github.com/jaredtmartin/bolt-go"
)

func Line(x1, y1, x2, y2 int) *SvgLine {
	line := &SvgLine{
		DefaultElement: bolt.NewDefaultElement("line"),
	}
	line.From(x1, y1).To(x2, y2)
	return line
}

type SvgLine struct {
	bolt.DefaultElement
}

func (t *SvgLine) From(x, y int) *SvgLine {
	t.Attr("x1", strconv.Itoa(x))
	t.Attr("y1", strconv.Itoa(y))
	return t
}
func (t *SvgLine) To(x, y int) *SvgLine {
	t.Attr("x2", strconv.Itoa(x))
	t.Attr("y2", strconv.Itoa(y))
	return t
}

func (t *SvgLine) Stroke(color string) *SvgLine {
	t.Attr("stroke", color)
	return t
}
func (t *SvgLine) StrokeWidth(width float64) *SvgLine {
	t.Attr("stroke-width", strconv.FormatFloat(width, 'f', -1, 64))
	return t
}

func (t *SvgLine) Opacity(opacity float64) *SvgLine {
	t.Attr("opacity", strconv.FormatFloat(opacity, 'f', -1, 64))
	return t
}

func (t *SvgLine) Fill(color string) *SvgLine {
	t.Attr("fill", color)
	return t
}
func (t *SvgLine) FillOpacity(opacity float64) *SvgLine {
	t.Attr("fill-opacity", strconv.FormatFloat(opacity, 'f', -1, 64))
	return t
}

func (t *SvgLine) StrokeDasharray(dasharray string) *SvgLine {
	t.Attr("stroke-dasharray", dasharray)
	return t
}

func (t *SvgLine) StrokeDashoffset(offset float64) *SvgLine {
	t.Attr("stroke-dashoffset", strconv.FormatFloat(offset, 'f', -1, 64))
	return t
}

func (t *SvgLine) StrokeLinecap(linecap string) *SvgLine {
	t.Attr("stroke-linecap", linecap)
	return t
}

func (t *SvgLine) StrokeLinejoin(linejoin string) *SvgLine {
	t.Attr("stroke-linejoin", linejoin)
	return t
}

func (t *SvgLine) StrokeMiterlimit(miterlimit float64) *SvgLine {
	t.Attr("stroke-miterlimit", strconv.FormatFloat(miterlimit, 'f', -1, 64))
	return t
}
