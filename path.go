package bolt

import (
	"strconv"
)

func Path() *SvgPath {
	return &SvgPath{
		DefaultElement: NewDefaultElement("path"),
	}
}

type SvgPath struct {
	DefaultElement
}

func (p *SvgPath) D(d string) *SvgPath {
	p.Attr("d", d)
	return p
}

func (p *SvgPath) Fill(fill string) *SvgPath {
	p.Attr("fill", fill)
	return p
}

func (p *SvgPath) Stroke(stroke string) *SvgPath {
	p.Attr("stroke", stroke)
	return p
}

func (p *SvgPath) StrokeWidth(width float64) *SvgPath {
	p.Attr("stroke-width", strconv.FormatFloat(width, 'f', -1, 64))
	return p
}

func (p *SvgPath) StrokeLinecap(linecap string) *SvgPath {
	p.Attr("stroke-linecap", linecap)
	return p
}

func (p *SvgPath) StrokeLinejoin(linejoin string) *SvgPath {
	p.Attr("stroke-linejoin", linejoin)
	return p
}

func (p *SvgPath) StrokeMiterlimit(miterlimit float64) *SvgPath {
	p.Attr("stroke-miterlimit", strconv.FormatFloat(miterlimit, 'f', -1, 64))
	return p
}

func (p *SvgPath) StrokeDasharray(dasharray string) *SvgPath {
	p.Attr("stroke-dasharray", dasharray)
	return p
}

func (p *SvgPath) StrokeDashoffset(offset float64) *SvgPath {
	p.Attr("stroke-dashoffset", strconv.FormatFloat(offset, 'f', -1, 64))
	return p
}
func (p *SvgPath) PathLength(length float64) *SvgPath {
	p.Attr("pathLength", strconv.FormatFloat(length, 'f', -1, 64))
	return p
}
func (p *SvgPath) Opacity(opacity float64) *SvgPath {
	p.Attr("opacity", strconv.FormatFloat(opacity, 'f', -1, 64))
	return p
}
func (p *SvgPath) Transform(transform string) *SvgPath {
	p.Attr("transform", transform)
	return p
}
func (p *SvgPath) Filter(filter string) *SvgPath {
	p.Attr("filter", filter)
	return p
}

func (p *SvgPath) ClipPath(clipPath string) *SvgPath {
	p.Attr("clip-path", clipPath)
	return p
}
