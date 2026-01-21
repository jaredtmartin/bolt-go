package bolt_test

import (
	"testing"

	"github.com/jaredtmartin/bolt-go"
	"github.com/stretchr/testify/assert"
)

func TestLine(t *testing.T) {
	line := bolt.Line(100, 100, 200, 200)
	result := line.Render()
	expected := `<line x1="100" x2="200" y1="100" y2="200"></line>`
	assert.Equalf(t, expected, result, "should match")
}

func TestLineStroke(t *testing.T) {
	line := bolt.Line(100, 100, 200, 200).Stroke("red")
	result := line.Render()
	expected := `<line stroke="red" x1="100" x2="200" y1="100" y2="200"></line>`
	assert.Equalf(t, expected, result, "should match")
}

func TestLineStrokeWidth(t *testing.T) {
	line := bolt.Line(100, 100, 200, 200).Stroke("red").StrokeWidth(2)
	result := line.Render()
	expected := `<line stroke-width="2" stroke="red" x1="100" x2="200" y1="100" y2="200"></line>`
	assert.Equalf(t, expected, result, "should match")
}

func TestLineOpacity(t *testing.T) {
	line := bolt.Line(100, 100, 200, 200).Opacity(0.5)
	result := line.Render()
	expected := `<line opacity="0.5" x1="100" x2="200" y1="100" y2="200"></line>`
	assert.Equalf(t, expected, result, "should match")
}

func TestLineFill(t *testing.T) {
	line := bolt.Line(100, 100, 200, 200).Fill("blue")
	result := line.Render()
	expected := `<line fill="blue" x1="100" x2="200" y1="100" y2="200"></line>`
	assert.Equalf(t, expected, result, "should match")
}

func TestLineFillOpacity(t *testing.T) {
	line := bolt.Line(100, 100, 200, 200).Fill("blue").FillOpacity(0.5)
	result := line.Render()
	expected := `<line fill-opacity="0.5" fill="blue" x1="100" x2="200" y1="100" y2="200"></line>`
	assert.Equalf(t, expected, result, "should match")
}

func TestLineStrokeDasharray(t *testing.T) {
	line := bolt.Line(100, 100, 200, 200).Stroke("red").StrokeDasharray("10,5")
	result := line.Render()
	expected := `<line stroke-dasharray="10,5" stroke="red" x1="100" x2="200" y1="100" y2="200"></line>`
	assert.Equalf(t, expected, result, "should match")
}

func TestLineStrokeDashoffset(t *testing.T) {
	line := bolt.Line(100, 100, 200, 200).Stroke("red").StrokeDashoffset(10)
	result := line.Render()
	expected := `<line stroke-dashoffset="10" stroke="red" x1="100" x2="200" y1="100" y2="200"></line>`
	assert.Equalf(t, expected, result, "should match")
}

func TestLineStrokeLinecap(t *testing.T) {
	line := bolt.Line(100, 100, 200, 200).Stroke("red").StrokeLinecap("round")
	result := line.Render()
	expected := `<line stroke-linecap="round" stroke="red" x1="100" x2="200" y1="100" y2="200"></line>`
	assert.Equalf(t, expected, result, "should match")
}
