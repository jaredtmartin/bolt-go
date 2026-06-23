package bolt_test

import (
	"testing"

	"github.com/jaredtmartin/bolt-go"
	"github.com/stretchr/testify/assert"
)

func TestPicture(t *testing.T) {
	path := bolt.Picture("Cat Pic",
		bolt.PictureSource{Src: "big-kitty.webp", Media: "(min-width: 1024px)"},
		bolt.PictureSource{Src: "nice-cat.webp", Media: "(min-width: 768px)"},
		bolt.PictureSource{Src: "little-kitty.webp"},
	)
	result := path.Render()
	expected := `<picture><source media="(min-width: 1024px)" srcset="big-kitty.webp"><source media="(min-width: 768px)" srcset="nice-cat.webp"><img alt="Cat Pic" src="little-kitty.webp"></picture>`
	assert.Equalf(t, expected, result, "should match")
}
