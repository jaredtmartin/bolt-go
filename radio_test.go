package bolt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestNewRadio(t *testing.T) {
	e := Radio("name", "Hello", "")
	// log.Printf("e: %v\n", e)
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="radio"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestRadioAttr(t *testing.T) {
	e := Radio("name", "Hello", "")
	e.Input.Attr("hx-post", "/url")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-post="/url" id="name-field" name="name" type="radio"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestRadioLabel(t *testing.T) {
	radio := Radio("name", "Hello", "")
	radio.Label.Text("goodbye")
	result := radio.Render()
	assert.Equalf(t, `<div><label for="name-field">goodbye</label><input id="name-field" name="name" type="radio"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestRadioValue(t *testing.T) {
	e := Radio("name", "Hello", "").Value("ispublic")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="radio" value="ispublic"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestRadioChecked(t *testing.T) {
	e := Radio("name", "Hello", "").Checked()
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input checked="checked" id="name-field" name="name" type="radio"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestRadioCheckedTrue(t *testing.T) {
	e := Radio("name", "Hello", "").Checked(true)
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input checked="checked" id="name-field" name="name" type="radio"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestRadioCheckedFalse(t *testing.T) {
	e := Radio("name", "Hello", "").Checked(false)
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="radio"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestRadioLabelClass(t *testing.T) {
	Radio := Radio("name", "Hello", "")
	Radio.Label.Class("hello")
	result := Radio.Render()
	assert.Equalf(t, `<div><label for="name-field" class="hello">Hello</label><input id="name-field" name="name" type="radio"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestRadioInputClass(t *testing.T) {
	Radio := Radio("name", "Hello", "")
	Radio.Input.Class("hello")
	result := Radio.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="radio" class="hello"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestRadioErrorClass(t *testing.T) {
	e := Radio("name", "Hello", "")
	e.Error.Class("hello")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="radio"><div id="name-field-error" class="hello"></div></div>`, result, "should match")
}
func TestRadioWrapperClass(t *testing.T) {
	e := Radio("name", "Hello", "").Class("hello")
	result := e.Render()
	assert.Equalf(t, `<div class="hello"><label for="name-field">Hello</label><input id="name-field" name="name" type="radio"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestRadioLabelStyle(t *testing.T) {
	e := Radio("name", "Hello", "")
	e.Label.Style("color: red")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field" style="color: red;">Hello</label><input id="name-field" name="name" type="radio"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestRadioInputStyle(t *testing.T) {
	e := Radio("name", "Hello", "")
	e.Input.Style("color: red")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="radio" style="color: red;"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestRadioErrorStyle(t *testing.T) {
	e := Radio("name", "Hello", "")
	e.Error.Style("color: red")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="radio"><div id="name-field-error" style="color: red;"></div></div>`, result, "should match")
}
func TestRadioWrapperStyle(t *testing.T) {
	e := Radio("name", "Hello", "").Style("color: red")
	result := e.Render()
	assert.Equalf(t, `<div style="color: red;"><label for="name-field">Hello</label><input id="name-field" name="name" type="radio"><div id="name-field-error"></div></div>`, result, "should match")
}
