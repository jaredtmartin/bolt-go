package bolt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type option struct {
	Label string
	Value string
}

func TestSelectWithDefaultOptionRenderer(t *testing.T) {
	opts := []string{"red", "green", "blue"}
	field := Select("color", "Color", "red", OptionStrings(opts))
	result := field.Render()
	expected := `<div><label for="color-field">Color</label><select id="color-field" name="color" value="red"><option selected="true" value="red">red</option><option value="green">green</option><option value="blue">blue</option></select><div id="color-field-error"></div></div>`
	assert.Equalf(t, expected, result, "should match")
}

func TestSelectWithCustomOptionRenderer(t *testing.T) {
	opts := []Option{
		{Label: "Red", Value: "red"},
		{Label: "Green", Value: "green"},
		{Label: "Blue", Value: "blue"},
	}
	renderOpt := func(opt Option, value string) Element {
		return Div("").Text(opt.Label)
	}
	field := Select("color", "Color", "red", opts, renderOpt)
	result := field.Render()
	assert.Equalf(t, `<div><label for="color-field">Color</label><select id="color-field" name="color" value="red"><div>Red</div><div>Green</div><div>Blue</div></select><div id="color-field-error"></div></div>`, result, "should match")
}

func TestSelectWithNoValue(t *testing.T) {
	opts := []string{"red", "green", "blue"}
	field := Select("color", "Color", "", OptionStrings(opts))
	result := field.Render()
	expected := `<div><label for="color-field">Color</label><select id="color-field" name="color"><option value="red">red</option><option value="green">green</option><option value="blue">blue</option></select><div id="color-field-error"></div></div>`
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}
