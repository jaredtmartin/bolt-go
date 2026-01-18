package bolt

import (
	"strconv"

	"github.com/jaredtmartin/bolt-go"
)

func Image(x, y, width, height int, href string) *SvgImage {
	image := &SvgImage{
		DefaultElement: bolt.NewDefaultElement("image"),
	}
	image.At(x, y).Size(width, height).Href(href)
	return image
}

type SvgImage struct {
	bolt.DefaultElement
}

func (i *SvgImage) At(x, y int) *SvgImage {
	i.X(x).Y(y)
	return i
}
func (i *SvgImage) Size(width, height int) *SvgImage {
	i.Width(width).Height(height)
	return i
}
func (i *SvgImage) X(x int) *SvgImage {
	i.Attr("x", strconv.Itoa(x))
	return i
}
func (i *SvgImage) Y(y int) *SvgImage {
	i.Attr("y", strconv.Itoa(y))
	return i
}
func (i *SvgImage) Href(href string) *SvgImage {
	i.Attr("href", href)
	return i
}
func (i *SvgImage) Width(width int) *SvgImage {
	i.Attr("width", strconv.Itoa(width))
	return i
}
func (i *SvgImage) Height(height int) *SvgImage {
	i.Attr("height", strconv.Itoa(height))
	return i
}
