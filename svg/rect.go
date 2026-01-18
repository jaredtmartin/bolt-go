package bolt

import (
	"strconv"

	"github.com/jaredtmartin/bolt-go"
)

func Rect(x, y, width, height int) *SvgRect {
	element := &SvgRect{
		DefaultElement: bolt.NewDefaultElement("rect"),
	}
	element.Attr("x", strconv.Itoa(x))
	element.Attr("y", strconv.Itoa(y))
	element.Attr("width", strconv.Itoa(width))
	element.Attr("height", strconv.Itoa(height))
	return element
}

type SvgRect struct {
	bolt.DefaultElement
}

func (t *SvgRect) X(x int) *SvgRect {
	t.Attr("x", strconv.Itoa(x))
	return t
}
func (t *SvgRect) Y(y int) *SvgRect {
	t.Attr("y", strconv.Itoa(y))
	return t
}
func (t *SvgRect) Width(width int) *SvgRect {
	t.Attr("width", strconv.Itoa(width))
	return t
}
func (t *SvgRect) Height(height int) *SvgRect {
	t.Attr("height", strconv.Itoa(height))
	return t
}
func (t *SvgRect) RX(rx int) *SvgRect {
	t.Attr("rx", strconv.Itoa(rx))
	return t
}
func (t *SvgRect) RY(ry int) *SvgRect {
	t.Attr("ry", strconv.Itoa(ry))
	return t
}
func (t *SvgRect) Fill(color string) *SvgRect {
	t.Attr("fill", color)
	return t
}
func (t *SvgRect) Stroke(color string) *SvgRect {
	t.Attr("stroke", color)
	return t
}
func (t *SvgRect) StrokeWidth(width int) *SvgRect {
	t.Attr("stroke-width", strconv.Itoa(width))
	return t
}
func (t *SvgRect) Opacity(opacity float64) *SvgRect {
	t.Attr("opacity", strconv.FormatFloat(opacity, 'f', -1, 64))
	return t
}
func (t *SvgRect) FillOpacity(opacity float64) *SvgRect {
	t.Attr("fill-opacity", strconv.FormatFloat(opacity, 'f', -1, 64))
	return t
}
func (t *SvgRect) StrokeOpacity(opacity float64) *SvgRect {
	t.Attr("stroke-opacity", strconv.FormatFloat(opacity, 'f', -1, 64))
	return t
}
func (t *SvgRect) StrokeDasharray(dasharray string) *SvgRect {
	t.Attr("stroke-dasharray", dasharray)
	return t
}
func (t *SvgRect) StrokeLinecap(linecap string) *SvgRect {
	t.Attr("stroke-linecap", linecap)
	return t
}
func (t *SvgRect) Transform(transform string) *SvgRect {
	t.Attr("transform", transform)
	return t
}
func (t *SvgRect) ClipPath(clipPath string) *SvgRect {
	t.Attr("clip-path", clipPath)
	return t
}
func (t *SvgRect) Filter(filter string) *SvgRect {
	t.Attr("filter", filter)
	return t
}
func (t *SvgRect) StrokeLinejoin(linejoin string) *SvgRect {
	t.Attr("stroke-linejoin", linejoin)
	return t
}

func (t *SvgRect) StrokeMiterlimit(miterlimit float64) *SvgRect {
	t.Attr("stroke-miterlimit", strconv.FormatFloat(miterlimit, 'f', -1, 64))
	return t
}
