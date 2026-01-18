package bolt_test

import (
	"testing"

	bolt "github.com/jaredtmartin/bolt-go/svg"
	"github.com/stretchr/testify/assert"
)

func TestText(t *testing.T) {
	text := bolt.Text(100, 100, "Hello, World!")
	result := text.Render()
	expected := `<text x="100" y="100">Hello, World!</text>`
	assert.Equalf(t, expected, result, "should match")
}

func TestTextStrokeDasharray(t *testing.T) {
	text := bolt.Text(100, 100, "Hello, World!").StrokeDasharray("5, 5")
	result := text.Render()
	expected := `<text stroke-dasharray="5, 5" x="100" y="100">Hello, World!</text>`
	assert.Equalf(t, expected, result, "should match")
}

func TestTextStrokeLinecap(t *testing.T) {
	text := bolt.Text(100, 100, "Hello, World!").StrokeLinecap("round")
	result := text.Render()
	expected := `<text stroke-linecap="round" x="100" y="100">Hello, World!</text>`
	assert.Equalf(t, expected, result, "should match")
}

func TestTextStrokeLinejoin(t *testing.T) {
	text := bolt.Text(100, 100, "Hello, World!").StrokeLinejoin("round")
	result := text.Render()
	expected := `<text stroke-linejoin="round" x="100" y="100">Hello, World!</text>`
	assert.Equalf(t, expected, result, "should match")
}

func TestTextStrokeMiterlimit(t *testing.T) {
	text := bolt.Text(100, 100, "Hello, World!").StrokeMiterlimit(10)
	result := text.Render()
	expected := `<text stroke-miterlimit="10" x="100" y="100">Hello, World!</text>`
	assert.Equalf(t, expected, result, "should match")
}

func TestTextTransform(t *testing.T) {
	text := bolt.Text(100, 100, "Hello, World!").Transform("translate(10, 10)")
	result := text.Render()
	expected := `<text transform="translate(10, 10)" x="100" y="100">Hello, World!</text>`
	assert.Equalf(t, expected, result, "should match")
}

func TestTextFill(t *testing.T) {
	text := bolt.Text(100, 100, "Hello, World!").Fill("red")
	result := text.Render()
	expected := `<text fill="red" x="100" y="100">Hello, World!</text>`
	assert.Equalf(t, expected, result, "should match")
}

func TestTextStroke(t *testing.T) {
	text := bolt.Text(100, 100, "Hello, World!").Stroke("blue")
	result := text.Render()
	expected := `<text stroke="blue" x="100" y="100">Hello, World!</text>`
	assert.Equalf(t, expected, result, "should match")
}

func TestTextStrokeWidth(t *testing.T) {
	text := bolt.Text(100, 100, "Hello, World!").StrokeWidth(2)
	result := text.Render()
	expected := `<text stroke-width="2" x="100" y="100">Hello, World!</text>`
	assert.Equalf(t, expected, result, "should match")
}

func TestTextStrokeDashoffset(t *testing.T) {
	text := bolt.Text(100, 100, "Hello, World!").StrokeDashoffset(10)
	result := text.Render()
	expected := `<text stroke-dashoffset="10" x="100" y="100">Hello, World!</text>`
	assert.Equalf(t, expected, result, "should match")
}

func TestTextDx(t *testing.T) {
	text := bolt.Text(100, 100, "Hello, World!").Dx(10)
	result := text.Render()
	expected := `<text dx="10" x="100" y="100">Hello, World!</text>`
	assert.Equalf(t, expected, result, "should match")
}

func TestTextDy(t *testing.T) {
	text := bolt.Text(100, 100, "Hello, World!").Dy(10)
	result := text.Render()
	expected := `<text dy="10" x="100" y="100">Hello, World!</text>`
	assert.Equalf(t, expected, result, "should match")
}

func TestTextRotate(t *testing.T) {
	text := bolt.Text(100, 100, "Hello, World!").Rotate(45)
	result := text.Render()
	expected := `<text rotate="45" x="100" y="100">Hello, World!</text>`
	assert.Equalf(t, expected, result, "should match")
}

func TestTextAnchor(t *testing.T) {
	text := bolt.Text(100, 100, "Hello, World!").TextAnchor("middle")
	result := text.Render()
	expected := `<text text-anchor="middle" x="100" y="100">Hello, World!</text>`
	assert.Equalf(t, expected, result, "should match")
}

func TestTextFontWeight(t *testing.T) {
	text := bolt.Text(100, 100, "Hello, World!").FontWeight("bold")
	result := text.Render()
	expected := `<text font-weight="bold" x="100" y="100">Hello, World!</text>`
	assert.Equalf(t, expected, result, "should match")
}

func TestTextLetterSpacing(t *testing.T) {
	text := bolt.Text(100, 100, "Hello, World!").LetterSpacing(2)
	result := text.Render()
	expected := `<text letter-spacing="2" x="100" y="100">Hello, World!</text>`
	assert.Equalf(t, expected, result, "should match")
}

func TestTextWordSpacing(t *testing.T) {
	text := bolt.Text(100, 100, "Hello, World!").WordSpacing(2)
	result := text.Render()
	expected := `<text word-spacing="2" x="100" y="100">Hello, World!</text>`
	assert.Equalf(t, expected, result, "should match")
}

func TestTextTextDecoration(t *testing.T) {
	text := bolt.Text(100, 100, "Hello, World!").TextDecoration("underline")
	result := text.Render()
	expected := `<text text-decoration="underline" x="100" y="100">Hello, World!</text>`
	assert.Equalf(t, expected, result, "should match")
}

func TestTextTextPath(t *testing.T) {
	text := bolt.Text(100, 100, "Hello, World!").TextPath("path")
	result := text.Render()
	expected := `<text textPath="path" x="100" y="100">Hello, World!</text>`
	assert.Equalf(t, expected, result, "should match")
}

func TestTextWritingMode(t *testing.T) {
	text := bolt.Text(100, 100, "Hello, World!").WritingMode("vertical-rl")
	result := text.Render()
	expected := `<text writing-mode="vertical-rl" x="100" y="100">Hello, World!</text>`
	assert.Equalf(t, expected, result, "should match")
}

func TestTextFontFamily(t *testing.T) {
	text := bolt.Text(100, 100, "Hello, World!").FontFamily("Arial")
	result := text.Render()
	expected := `<text font-family="Arial" x="100" y="100">Hello, World!</text>`
	assert.Equalf(t, expected, result, "should match")
}

func TestTextFontSize(t *testing.T) {
	text := bolt.Text(100, 100, "Hello, World!").FontSize(12)
	result := text.Render()
	expected := `<text font-size="12" x="100" y="100">Hello, World!</text>`
	assert.Equalf(t, expected, result, "should match")
}

func TestTextFontStyle(t *testing.T) {
	text := bolt.Text(100, 100, "Hello, World!").FontStyle("italic")
	result := text.Render()
	expected := `<text font-style="italic" x="100" y="100">Hello, World!</text>`
	assert.Equalf(t, expected, result, "should match")
}
