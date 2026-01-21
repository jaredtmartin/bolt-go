package bolt_test

import (
	"testing"

	"github.com/jaredtmartin/bolt-go"
	"github.com/stretchr/testify/assert"
)

func TestCircle(t *testing.T) {
	circle := bolt.Circle(100, 100, 50)
	result := circle.Render()
	expected := `<circle cx="100" cy="100" r="50"></circle>`
	assert.Equalf(t, expected, result, "should match")
}

func TestCircleFill(t *testing.T) {
	circle := bolt.Circle(100, 100, 50).Fill("red")
	result := circle.Render()
	expected := `<circle cx="100" cy="100" fill="red" r="50"></circle>`
	assert.Equalf(t, expected, result, "should match")
}

func TestCircleStroke(t *testing.T) {
	circle := bolt.Circle(100, 100, 50).Stroke("blue")
	result := circle.Render()
	expected := `<circle cx="100" cy="100" r="50" stroke="blue"></circle>`
	assert.Equalf(t, expected, result, "should match")
}

func TestCircleStrokeWidth(t *testing.T) {
	circle := bolt.Circle(100, 100, 50).Stroke("blue").StrokeWidth(2)
	result := circle.Render()
	expected := `<circle cx="100" cy="100" r="50" stroke-width="2" stroke="blue"></circle>`
	assert.Equalf(t, expected, result, "should match")
}

func TestCircleOpacity(t *testing.T) {
	circle := bolt.Circle(100, 100, 50).Opacity(0.5)
	result := circle.Render()
	expected := `<circle cx="100" cy="100" opacity="0.5" r="50"></circle>`
	assert.Equalf(t, expected, result, "should match")
}

func TestCircleFillOpacity(t *testing.T) {
	circle := bolt.Circle(100, 100, 50).FillOpacity(0.5)
	result := circle.Render()
	expected := `<circle cx="100" cy="100" fill-opacity="0.5" r="50"></circle>`
	assert.Equalf(t, expected, result, "should match")
}

func TestCircleStrokeOpacity(t *testing.T) {
	circle := bolt.Circle(100, 100, 50).StrokeOpacity(0.5)
	result := circle.Render()
	expected := `<circle cx="100" cy="100" r="50" stroke-opacity="0.5"></circle>`
	assert.Equalf(t, expected, result, "should match")
}

func TestCircleFilter(t *testing.T) {
	circle := bolt.Circle(100, 100, 50).Filter("url(#filter)")
	result := circle.Render()
	expected := `<circle cx="100" cy="100" filter="url(#filter)" r="50"></circle>`
	assert.Equalf(t, expected, result, "should match")
}

func TestCircleStrokeDasharray(t *testing.T) {
	circle := bolt.Circle(100, 100, 50).StrokeDasharray("5, 5")
	result := circle.Render()
	expected := `<circle cx="100" cy="100" r="50" stroke-dasharray="5, 5"></circle>`
	assert.Equalf(t, expected, result, "should match")
}

func TestCircleStrokeLinecap(t *testing.T) {
	circle := bolt.Circle(100, 100, 50).StrokeLinecap("round")
	result := circle.Render()
	expected := `<circle cx="100" cy="100" r="50" stroke-linecap="round"></circle>`
	assert.Equalf(t, expected, result, "should match")
}
