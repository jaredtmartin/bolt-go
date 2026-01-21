package bolt_test

import (
	"testing"

	"github.com/jaredtmartin/bolt-go"
	"github.com/stretchr/testify/assert"
)

func TestImage(t *testing.T) {
	image := bolt.Image(0, 0, 100, 100, "image.png")
	result := image.Render()
	expected := `<image height="100" href="image.png" width="100" x="0" y="0"></image>`
	assert.Equalf(t, expected, result, "should match")
}

func TestImageHref(t *testing.T) {
	image := bolt.Image(0, 0, 100, 100, "image.png").Href("image.png")
	result := image.Render()
	expected := `<image height="100" href="image.png" width="100" x="0" y="0"></image>`
	assert.Equalf(t, expected, result, "should match")
}

func TestImageSize(t *testing.T) {
	image := bolt.Image(0, 0, 100, 100, "image.png").Size(100, 100)
	result := image.Render()
	expected := `<image height="100" href="image.png" width="100" x="0" y="0"></image>`
	assert.Equalf(t, expected, result, "should match")
}

func TestImageWidth(t *testing.T) {
	image := bolt.Image(0, 0, 100, 100, "image.png").Width(100)
	result := image.Render()
	expected := `<image height="100" href="image.png" width="100" x="0" y="0"></image>`
	assert.Equalf(t, expected, result, "should match")
}

func TestImageHeight(t *testing.T) {
	image := bolt.Image(0, 0, 100, 100, "image.png").Height(100)
	result := image.Render()
	expected := `<image height="100" href="image.png" width="100" x="0" y="0"></image>`
	assert.Equalf(t, expected, result, "should match")
}

func TestImageAt(t *testing.T) {
	image := bolt.Image(0, 0, 100, 100, "image.png").At(0, 0)
	result := image.Render()
	expected := `<image height="100" href="image.png" width="100" x="0" y="0"></image>`
	assert.Equalf(t, expected, result, "should match")
}

func TestImageX(t *testing.T) {
	image := bolt.Image(0, 0, 100, 100, "image.png").X(0)
	result := image.Render()
	expected := `<image height="100" href="image.png" width="100" x="0" y="0"></image>`
	assert.Equalf(t, expected, result, "should match")
}

func TestImageY(t *testing.T) {
	image := bolt.Image(0, 0, 100, 100, "image.png").Y(0)
	result := image.Render()
	expected := `<image height="100" href="image.png" width="100" x="0" y="0"></image>`
	assert.Equalf(t, expected, result, "should match")
}
