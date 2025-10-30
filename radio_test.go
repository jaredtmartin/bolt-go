package bolt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestNewRadio(t *testing.T) {
	e := Radio("name", "Hello")
	// log.Printf("e: %v\n", e)
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="radio"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestRadioAttr(t *testing.T) {
	e := Radio("name", "Hello")
	e.GetInput().Attr("hx-post", "/url")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-post="/url" id="name-field" name="name" type="radio"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestRadioLabel(t *testing.T) {
	e := Radio("name", "Hello").Label("goodbye")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">goodbye</label><input id="name-field" name="name" type="radio"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestRadioName(t *testing.T) {
	e := Radio("name", "Hello").Name("hello")
	result := e.Render()
	assert.Equalf(t, `<div><label for="hello-field">Hello</label><input id="hello-field" name="hello" type="radio"><div id="hello-field-error"></div></div>`, result, "should match")
}
func TestRadioNameWithExistingId(t *testing.T) {
	e := Radio("name", "Hello").Id("blue").Name("hello")
	result := e.Render()
	assert.Equalf(t, `<div><label for="blue">Hello</label><input id="blue" name="hello" type="radio"><div id="blue-error"></div></div>`, result, "should match")
}
func TestRadioId(t *testing.T) {
	e := Radio("name", "Hello").Id("hello")
	result := e.Render()
	assert.Equalf(t, `<div><label for="hello">Hello</label><input id="hello" name="name" type="radio"><div id="hello-error"></div></div>`, result, "should match")
}

func TestRadioValue(t *testing.T) {
	e := Radio("name", "Hello").Value("ispublic")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="radio" value="ispublic"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestRadioChecked(t *testing.T) {
	e := Radio("name", "Hello").Checked()
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input checked="checked" id="name-field" name="name" type="radio"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestRadioCheckedTrue(t *testing.T) {
	e := Radio("name", "Hello").Checked(true)
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input checked="checked" id="name-field" name="name" type="radio"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestRadioCheckedFalse(t *testing.T) {
	e := Radio("name", "Hello").Checked(false)
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="radio"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestRadioTarget(t *testing.T) {
	e := Radio("name", "Hello").HXTarget("hello")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-target="hello" id="name-field" name="name" type="radio"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestRadioSwap(t *testing.T) {
	e := Radio("name", "Hello").HXSwap("hello")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-swap="hello" id="name-field" name="name" type="radio"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestRadioLabelClass(t *testing.T) {
	Radio := Radio("name", "Hello")
	Radio.GetLabel().Class("hello")
	result := Radio.Render()
	assert.Equalf(t, `<div><label for="name-field" class="hello">Hello</label><input id="name-field" name="name" type="radio"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestRadioInputClass(t *testing.T) {
	Radio := Radio("name", "Hello")
	Radio.GetInput().Class("hello")
	result := Radio.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="radio" class="hello"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestRadioErrorClass(t *testing.T) {
	e := Radio("name", "Hello")
	e.GetError().Class("hello")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="radio"><div id="name-field-error" class="hello"></div></div>`, result, "should match")
}
func TestRadioWrapperClass(t *testing.T) {
	e := Radio("name", "Hello").Class("hello")
	result := e.Render()
	assert.Equalf(t, `<div class="hello"><label for="name-field">Hello</label><input id="name-field" name="name" type="radio"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestRadioLabelStyle(t *testing.T) {
	e := Radio("name", "Hello")
	e.GetLabel().Style("color: red")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field" style="color: red;">Hello</label><input id="name-field" name="name" type="radio"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestRadioInputStyle(t *testing.T) {
	e := Radio("name", "Hello")
	e.GetInput().Style("color: red")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="radio" style="color: red;"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestRadioErrorStyle(t *testing.T) {
	e := Radio("name", "Hello")
	e.GetError().Style("color: red")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="radio"><div id="name-field-error" style="color: red;"></div></div>`, result, "should match")
}
func TestRadioWrapperStyle(t *testing.T) {
	e := Radio("name", "Hello").Style("color: red")
	result := e.Render()
	assert.Equalf(t, `<div style="color: red;"><label for="name-field">Hello</label><input id="name-field" name="name" type="radio"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestRadioText(t *testing.T) {
	e := Radio("name", "Hello").Error("Can't do that.")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="radio"><div id="name-field-error">Can't do that.</div></div>`, result, "should match")
}

// TestRadioPost
func TestRadioPost(t *testing.T) {
	e := Radio("name", "Hello").HXPost("/url")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-post="/url" id="name-field" name="name" type="radio"><div id="name-field-error"></div></div>`, result, "should match")
}

// TestRadioGet
func TestRadioGet(t *testing.T) {
	e := Radio("name", "Hello").HXGet("/url")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-get="/url" id="name-field" name="name" type="radio"><div id="name-field-error"></div></div>`, result, "should match")
}

// TestRadioPut
func TestRadioPut(t *testing.T) {
	e := Radio("name", "Hello").HXPut("/url")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-put="/url" id="name-field" name="name" type="radio"><div id="name-field-error"></div></div>`, result, "should match")
}

// TestRadioPatch
func TestRadioPatch(t *testing.T) {
	e := Radio("name", "Hello").HXPatch("/url")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-patch="/url" id="name-field" name="name" type="radio"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestRadioDelete(t *testing.T) {
	e := Radio("name", "Hello").HXDelete("/url")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-delete="/url" id="name-field" name="name" type="radio"><div id="name-field-error"></div></div>`, result, "should match")
}
