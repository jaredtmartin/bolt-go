package bolt

import (
	"fmt"
	"strconv"
)

func Svg() *SvgElement {
	svg := &(SvgElement{
		DefaultElement: NewDefaultElement("svg"),
	})
	svg.Version("1.1").
		Xmlns("http://www.w3.org/2000/svg").
		XmlnsXlink("http://www.w3.org/1999/xlink")
	return svg
}

type SvgElement struct {
	DefaultElement
}

func (s *SvgElement) ViewBox(x, y, width, height int) *SvgElement {
	s.Attr("viewBox", fmt.Sprintf("%d %d %d %d", x, y, width, height))
	return s
}

func (s *SvgElement) Version(version string) *SvgElement {
	s.Attr("version", version)
	return s
}
func (s *SvgElement) Width(width int) *SvgElement {
	s.Attr("width", strconv.Itoa(width))
	return s
}
func (s *SvgElement) Height(height int) *SvgElement {
	s.Attr("height", strconv.Itoa(height))
	return s
}
func (s *SvgElement) Xmlns(xmlns string) *SvgElement {
	s.Attr("xmlns", xmlns)
	return s
}
func (s *SvgElement) XmlnsXlink(xmlnsXlink string) *SvgElement {
	s.Attr("xmlns:xlink", xmlnsXlink)
	return s
}
func (s *SvgElement) D(d string) *SvgElement {
	s.Attr("d", d)
	return s
}
