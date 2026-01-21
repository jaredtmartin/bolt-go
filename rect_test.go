package bolt_test

import (
	"testing"

	"github.com/jaredtmartin/bolt-go"
	"github.com/stretchr/testify/assert"
)

func TestRect(t *testing.T) {
	rect := bolt.Rect(100, 100, 50, 50)
	result := rect.Render()
	expected := `<rect height="50" width="50" x="100" y="100"></rect>`
	assert.Equalf(t, expected, result, "should match")
}

func TestRectFill(t *testing.T) {
	rect := bolt.Rect(100, 100, 50, 50).Fill("red")
	result := rect.Render()
	expected := `<rect fill="red" height="50" width="50" x="100" y="100"></rect>`
	assert.Equalf(t, expected, result, "should match")
}

func TestRectStroke(t *testing.T) {
	rect := bolt.Rect(100, 100, 50, 50).Stroke("blue")
	result := rect.Render()
	expected := `<rect height="50" stroke="blue" width="50" x="100" y="100"></rect>`
	assert.Equalf(t, expected, result, "should match")
}

func TestRectStrokeWidth(t *testing.T) {
	rect := bolt.Rect(100, 100, 50, 50).Stroke("blue").StrokeWidth(2)
	result := rect.Render()
	expected := `<rect height="50" stroke-width="2" stroke="blue" width="50" x="100" y="100"></rect>`
	assert.Equalf(t, expected, result, "should match")
}

func TestRectOpacity(t *testing.T) {
	rect := bolt.Rect(100, 100, 50, 50).Opacity(0.5)
	result := rect.Render()
	expected := `<rect height="50" opacity="0.5" width="50" x="100" y="100"></rect>`
	assert.Equalf(t, expected, result, "should match")
}

func TestRectFillOpacity(t *testing.T) {
	rect := bolt.Rect(100, 100, 50, 50).FillOpacity(0.5)
	result := rect.Render()
	expected := `<rect fill-opacity="0.5" height="50" width="50" x="100" y="100"></rect>`
	assert.Equalf(t, expected, result, "should match")
}

func TestRectStrokeOpacity(t *testing.T) {
	rect := bolt.Rect(100, 100, 50, 50).StrokeOpacity(0.5)
	result := rect.Render()
	expected := `<rect height="50" stroke-opacity="0.5" width="50" x="100" y="100"></rect>`
	assert.Equalf(t, expected, result, "should match")
}

func TestRectFilter(t *testing.T) {
	rect := bolt.Rect(100, 100, 50, 50).Filter("url(#filter)")
	result := rect.Render()
	expected := `<rect filter="url(#filter)" height="50" width="50" x="100" y="100"></rect>`
	assert.Equalf(t, expected, result, "should match")
}

func TestRectTransform(t *testing.T) {
	rect := bolt.Rect(100, 100, 50, 50).Transform("translate(10, 10)")
	result := rect.Render()
	expected := `<rect height="50" transform="translate(10, 10)" width="50" x="100" y="100"></rect>`
	assert.Equalf(t, expected, result, "should match")
}

func TestRectClipPath(t *testing.T) {
	rect := bolt.Rect(100, 100, 50, 50).ClipPath("url(#clipPath)")
	result := rect.Render()
	expected := `<rect clip-path="url(#clipPath)" height="50" width="50" x="100" y="100"></rect>`
	assert.Equalf(t, expected, result, "should match")
}

func TestRectStrokeDasharray(t *testing.T) {
	rect := bolt.Rect(100, 100, 50, 50).StrokeDasharray("5, 5")
	result := rect.Render()
	expected := `<rect height="50" stroke-dasharray="5, 5" width="50" x="100" y="100"></rect>`
	assert.Equalf(t, expected, result, "should match")
}

func TestRectStrokeLinecap(t *testing.T) {
	rect := bolt.Rect(100, 100, 50, 50).StrokeLinecap("round")
	result := rect.Render()
	expected := `<rect height="50" stroke-linecap="round" width="50" x="100" y="100"></rect>`
	assert.Equalf(t, expected, result, "should match")
}

func TestRectStrokeLinejoin(t *testing.T) {
	rect := bolt.Rect(100, 100, 50, 50).StrokeLinejoin("round")
	result := rect.Render()
	expected := `<rect height="50" stroke-linejoin="round" width="50" x="100" y="100"></rect>`
	assert.Equalf(t, expected, result, "should match")
}

func TestRectStrokeMiterlimit(t *testing.T) {
	rect := bolt.Rect(100, 100, 50, 50).StrokeMiterlimit(2)
	result := rect.Render()
	expected := `<rect height="50" stroke-miterlimit="2" width="50" x="100" y="100"></rect>`
	assert.Equalf(t, expected, result, "should match")
}
