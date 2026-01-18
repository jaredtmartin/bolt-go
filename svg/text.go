package bolt

import (
	"strconv"

	"github.com/jaredtmartin/bolt-go"
)

func Text(x, y int, text string) *SvgText {
	textElement := &SvgText{
		DefaultElement: bolt.NewDefaultElement("text"),
	}
	textElement.X(x).Y(y).Text(text)
	return textElement
}

type SvgText struct {
	bolt.DefaultElement
}

func (t *SvgText) X(x int) *SvgText {
	t.Attr("x", strconv.Itoa(x))
	return t
}
func (t *SvgText) Y(y int) *SvgText {
	t.Attr("y", strconv.Itoa(y))
	return t
}
func (t *SvgText) Dx(dx float64) *SvgText {
	t.Attr("dx", strconv.FormatFloat(dx, 'f', -1, 64))
	return t
}
func (t *SvgText) Rotate(angle float64) *SvgText {
	t.Attr("rotate", strconv.FormatFloat(angle, 'f', -1, 64))
	return t
}
func (t *SvgText) Dy(dy float64) *SvgText {
	t.Attr("dy", strconv.FormatFloat(dy, 'f', -1, 64))
	return t
}
func (t *SvgText) Text(text string) *SvgText {
	t.DefaultElement.Text(text)
	return t
}
func (t *SvgText) TextAnchor(anchor string) *SvgText {
	t.Attr("text-anchor", anchor)
	return t
}
func (t *SvgText) DominantBaseline(baseline string) *SvgText {
	t.Attr("dominant-baseline", baseline)
	return t
}
func (t *SvgText) FontSize(size int) *SvgText {
	t.Attr("font-size", strconv.Itoa(size))
	return t
}
func (t *SvgText) Fill(color string) *SvgText {
	t.Attr("fill", color)
	return t
}
func (t *SvgText) StrokeDasharray(dasharray string) *SvgText {
	t.Attr("stroke-dasharray", dasharray)
	return t
}
func (t *SvgText) StrokeLinecap(linecap string) *SvgText {
	t.Attr("stroke-linecap", linecap)
	return t
}
func (t *SvgText) StrokeLinejoin(linejoin string) *SvgText {
	t.Attr("stroke-linejoin", linejoin)
	return t
}
func (t *SvgText) StrokeMiterlimit(miterlimit float64) *SvgText {
	t.Attr("stroke-miterlimit", strconv.FormatFloat(miterlimit, 'f', -1, 64))
	return t
}
func (t *SvgText) Transform(transform string) *SvgText {
	t.Attr("transform", transform)
	return t
}
func (t *SvgText) StrokeDashoffset(offset float64) *SvgText {
	t.Attr("stroke-dashoffset", strconv.FormatFloat(offset, 'f', -1, 64))
	return t
}
func (t *SvgText) StrokeWidth(width float64) *SvgText {
	t.Attr("stroke-width", strconv.FormatFloat(width, 'f', -1, 64))
	return t
}
func (t *SvgText) Stroke(color string) *SvgText {
	t.Attr("stroke", color)
	return t
}

func (t *SvgText) AlignmentBaseline(baseline string) *SvgText {
	t.Attr("alignment-baseline", baseline)
	return t
}
func (t *SvgText) TextLength(length float64) *SvgText {
	t.Attr("textLength", strconv.FormatFloat(length, 'f', -1, 64))
	return t
}
func (t *SvgText) LengthAdjust(lengthAdjust string) *SvgText {
	t.Attr("lengthAdjust", lengthAdjust)
	return t
}
func (t *SvgText) FontFamily(fontFamily string) *SvgText {
	t.Attr("font-family", fontFamily)
	return t
}
func (t *SvgText) WritingMode(mode string) *SvgText {
	t.Attr("writing-mode", mode)
	return t
}
func (t *SvgText) FontWeight(weight string) *SvgText {
	t.Attr("font-weight", weight)
	return t
}
func (t *SvgText) FontSizeAdjust(sizeAdjust string) *SvgText {
	t.Attr("font-size-adjust", sizeAdjust)
	return t
}
func (t *SvgText) LetterSpacing(spacing float64) *SvgText {
	t.Attr("letter-spacing", strconv.FormatFloat(spacing, 'f', -1, 64))
	return t
}
func (t *SvgText) WordSpacing(spacing float64) *SvgText {
	t.Attr("word-spacing", strconv.FormatFloat(spacing, 'f', -1, 64))
	return t
}
func (t *SvgText) TextDecoration(decoration string) *SvgText {
	t.Attr("text-decoration", decoration)
	return t
}
func (t *SvgText) TextPath(path string) *SvgText {
	t.Attr("textPath", path)
	return t
}
func (t *SvgText) FontStyle(style string) *SvgText {
	t.Attr("font-style", style)
	return t
}
