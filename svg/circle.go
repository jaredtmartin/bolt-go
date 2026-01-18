package bolt

import (
	"strconv"

	"github.com/jaredtmartin/bolt-go"
)

func Circle(x, y, radius int) *SvgCircle {
	circle := &SvgCircle{
		DefaultElement: bolt.NewDefaultElement("circle"),
	}
	circle.At(x, y).Radius(radius)
	return circle
}

type SvgCircle struct {
	bolt.DefaultElement
}

func (t *SvgCircle) CenterX(x int) *SvgCircle {
	t.Attr("cx", strconv.Itoa(x))
	return t
}
func (t *SvgCircle) CenterY(y int) *SvgCircle {
	t.Attr("cy", strconv.Itoa(y))
	return t
}
func (t *SvgCircle) At(x, y int) *SvgCircle {
	t.Attr("cx", strconv.Itoa(x))
	t.Attr("cy", strconv.Itoa(y))
	return t
}
func (t *SvgCircle) Radius(radius int) *SvgCircle {
	t.Attr("r", strconv.Itoa(radius))
	return t
}
func (t *SvgCircle) Fill(color string) *SvgCircle {
	t.Attr("fill", color)
	return t
}
func (t *SvgCircle) Stroke(color string) *SvgCircle {
	t.Attr("stroke", color)
	return t
}
func (t *SvgCircle) StrokeWidth(width int) *SvgCircle {
	t.Attr("stroke-width", strconv.Itoa(width))
	return t
}
func (t *SvgCircle) Opacity(opacity float64) *SvgCircle {
	t.Attr("opacity", strconv.FormatFloat(opacity, 'f', -1, 64))
	return t
}
func (t *SvgCircle) FillOpacity(opacity float64) *SvgCircle {
	t.Attr("fill-opacity", strconv.FormatFloat(opacity, 'f', -1, 64))
	return t
}
func (t *SvgCircle) StrokeOpacity(opacity float64) *SvgCircle {
	t.Attr("stroke-opacity", strconv.FormatFloat(opacity, 'f', -1, 64))
	return t
}
func (t *SvgCircle) Filter(filter string) *SvgCircle {
	t.Attr("filter", filter)
	return t
}
func (t *SvgCircle) StrokeDasharray(dasharray string) *SvgCircle {
	t.Attr("stroke-dasharray", dasharray)
	return t
}
func (t *SvgCircle) StrokeLinecap(linecap string) *SvgCircle {
	t.Attr("stroke-linecap", linecap)
	return t
}
func (t *SvgCircle) StrokeLinejoin(linejoin string) *SvgCircle {
	t.Attr("stroke-linejoin", linejoin)
	return t
}
func (t *SvgCircle) StrokeMiterlimit(miterlimit float64) *SvgCircle {
	t.Attr("stroke-miterlimit", strconv.FormatFloat(miterlimit, 'f', -1, 64))
	return t
}
