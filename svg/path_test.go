package bolt_test

import (
	"testing"

	bolt "github.com/jaredtmartin/bolt-go/svg"
	"github.com/stretchr/testify/assert"
)

func TestPath(t *testing.T) {
	path := bolt.Path().D("M10 10 L20 20 L30 10 Z")
	result := path.Render()
	expected := `<path d="M10 10 L20 20 L30 10 Z">`
	assert.Equalf(t, expected, result, "should match")
}

func TestPathStrokeLinecap(t *testing.T) {
	path := bolt.Path().D("M10 10 L20 20 L30 10 Z").StrokeLinecap("round")
	result := path.Render()
	expected := `<path d="M10 10 L20 20 L30 10 Z" stroke-linecap="round">`
	assert.Equalf(t, expected, result, "should match")
}

func TestPathStrokeLinejoin(t *testing.T) {
	path := bolt.Path().D("M10 10 L20 20 L30 10 Z").StrokeLinejoin("round")
	result := path.Render()
	expected := `<path d="M10 10 L20 20 L30 10 Z" stroke-linejoin="round">`
	assert.Equalf(t, expected, result, "should match")
}

func TestPathStrokeMiterlimit(t *testing.T) {
	path := bolt.Path().D("M10 10 L20 20 L30 10 Z").StrokeMiterlimit(2)
	result := path.Render()
	expected := `<path d="M10 10 L20 20 L30 10 Z" stroke-miterlimit="2">`
	assert.Equalf(t, expected, result, "should match")
}

func TestPathStrokeDasharray(t *testing.T) {
	path := bolt.Path().D("M10 10 L20 20 L30 10 Z").StrokeDasharray("10 5")

	result := path.Render()
	expected := `<path d="M10 10 L20 20 L30 10 Z" stroke-dasharray="10 5">`
	assert.Equalf(t, expected, result, "should match")
}

func TestPathStrokeDashoffset(t *testing.T) {
	path := bolt.Path().D("M10 10 L20 20 L30 10 Z").StrokeDashoffset(10)
	result := path.Render()
	expected := `<path d="M10 10 L20 20 L30 10 Z" stroke-dashoffset="10">`
	assert.Equalf(t, expected, result, "should match")
}

func TestPathOpacity(t *testing.T) {
	path := bolt.Path().D("M10 10 L20 20 L30 10 Z").Opacity(0.5)
	result := path.Render()
	expected := `<path d="M10 10 L20 20 L30 10 Z" opacity="0.5">`
	assert.Equalf(t, expected, result, "should match")
}

func TestPathPathLength(t *testing.T) {
	path := bolt.Path().D("M10 10 L20 20 L30 10 Z").PathLength(100)
	result := path.Render()
	expected := `<path d="M10 10 L20 20 L30 10 Z" pathLength="100">`
	assert.Equalf(t, expected, result, "should match")
}

func TestPathTransform(t *testing.T) {
	path := bolt.Path().D("M10 10 L20 20 L30 10 Z").Transform("translate(10, 10)")
	result := path.Render()
	expected := `<path d="M10 10 L20 20 L30 10 Z" transform="translate(10, 10)">`
	assert.Equalf(t, expected, result, "should match")
}

func TestPathFilter(t *testing.T) {
	path := bolt.Path().D("M10 10 L20 20 L30 10 Z").Filter("url(#filter)")
	result := path.Render()
	expected := `<path d="M10 10 L20 20 L30 10 Z" filter="url(#filter)">`
	assert.Equalf(t, expected, result, "should match")
}

func TestPathClipPath(t *testing.T) {
	path := bolt.Path().D("M10 10 L20 20 L30 10 Z").ClipPath("url(#clip)")
	result := path.Render()
	expected := `<path clip-path="url(#clip)" d="M10 10 L20 20 L30 10 Z">`
	assert.Equalf(t, expected, result, "should match")
}
