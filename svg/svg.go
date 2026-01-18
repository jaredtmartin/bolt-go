package bolt

import (
	"fmt"
	"strconv"

	"github.com/jaredtmartin/bolt-go"
)

func NewSvg(width, height int) *Svg {
	svg := &(Svg{
		DefaultElement: bolt.NewDefaultElement("svg"),
	})
	svg.ViewBox(0, 0, width, height).
		Version("1.1").
		Xmlns("http://www.w3.org/2000/svg").
		XmlnsXlink("http://www.w3.org/1999/xlink").
		Width(width).
		Height(height).
		Class("fill-current w-full")
	return svg
}

type Svg struct {
	bolt.DefaultElement
}

func (s *Svg) ViewBox(x, y, width, height int) *Svg {
	s.Attr("viewBox", fmt.Sprintf("%d %d %d %d", x, y, width, height))
	return s
}

func (s *Svg) Version(version string) *Svg {
	s.Attr("version", version)
	return s
}
func (s *Svg) Width(width int) *Svg {
	s.Attr("width", strconv.Itoa(width))
	return s
}
func (s *Svg) Height(height int) *Svg {
	s.Attr("height", strconv.Itoa(height))
	return s
}
func (s *Svg) Xmlns(xmlns string) *Svg {
	s.Attr("xmlns", xmlns)
	return s
}
func (s *Svg) XmlnsXlink(xmlnsXlink string) *Svg {
	s.Attr("xmlns:xlink", xmlnsXlink)
	return s
}
