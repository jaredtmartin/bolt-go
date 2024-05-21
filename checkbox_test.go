package bolt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestNewCheckbox(t *testing.T) {
	e := Checkbox("name", "Hello")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="checkbox"/><div id="name-field-error"></div></div>`, result, "should match")
}
func TestCheckboxAttr(t *testing.T) {
	e := Checkbox("name", "Hello").Attr("hx-post", "/url")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-post="/url" id="name-field" name="name" type="checkbox"/><div id="name-field-error"></div></div>`, result, "should match")
}
func TestCheckboxLabel(t *testing.T) {
	e := Checkbox("name", "Hello").Label("hello")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">hello</label><input id="name-field" name="name" type="checkbox"/><div id="name-field-error"></div></div>`, result, "should match")
}
func TestCheckboxName(t *testing.T) {
	e := Checkbox("name", "Hello").Name("hello")
	result := e.Render()
	assert.Equalf(t, `<div><label for="hello-field">Hello</label><input id="hello-field" name="hello" type="checkbox"/><div id="hello-field-error"></div></div>`, result, "should match")
}
func TestCheckboxNameWithExistingId(t *testing.T) {
	e := Checkbox("name", "Hello").Id("blue").Name("hello")
	result := e.Render()
	assert.Equalf(t, `<div><label for="blue">Hello</label><input id="blue" name="hello" type="checkbox"/><div id="blue-error"></div></div>`, result, "should match")
}
func TestCheckboxId(t *testing.T) {
	e := Checkbox("name", "Hello").Id("hello")
	result := e.Render()
	assert.Equalf(t, `<div><label for="hello">Hello</label><input id="hello" name="name" type="checkbox"/><div id="hello-error"></div></div>`, result, "should match")
}

func TestCheckboxValue(t *testing.T) {
	e := Checkbox("name", "Hello").Value("ispublic")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="checkbox" value="ispublic"/><div id="name-field-error"></div></div>`, result, "should match")
}
func TestCheckboxChecked(t *testing.T) {
	e := Checkbox("name", "Hello").Checked()
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input checked="checked" id="name-field" name="name" type="checkbox"/><div id="name-field-error"></div></div>`, result, "should match")
}
func TestCheckboxCheckedTrue(t *testing.T) {
	e := Checkbox("name", "Hello").Checked(true)
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input checked="checked" id="name-field" name="name" type="checkbox"/><div id="name-field-error"></div></div>`, result, "should match")
}
func TestCheckboxCheckedFalse(t *testing.T) {
	e := Checkbox("name", "Hello").Checked(false)
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="checkbox"/><div id="name-field-error"></div></div>`, result, "should match")
}
func TestCheckboxTarget(t *testing.T) {
	e := Checkbox("name", "Hello").Target("hello")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-target="hello" id="name-field" name="name" type="checkbox"/><div id="name-field-error"></div></div>`, result, "should match")
}
func TestCheckboxSwap(t *testing.T) {
	e := Checkbox("name", "Hello").Swap("hello")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-swap="hello" id="name-field" name="name" type="checkbox"/><div id="name-field-error"></div></div>`, result, "should match")
}
func TestCheckboxLabelClass(t *testing.T) {
	e := Checkbox("name", "Hello").LabelClass("hello")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field" class="hello">Hello</label><input id="name-field" name="name" type="checkbox"/><div id="name-field-error"></div></div>`, result, "should match")
}
func TestCheckboxInputClass(t *testing.T) {
	e := Checkbox("name", "Hello").InputClass("hello")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="checkbox" class="hello"/><div id="name-field-error"></div></div>`, result, "should match")
}
func TestCheckboxErrorClass(t *testing.T) {
	e := Checkbox("name", "Hello").ErrorClass("hello")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="checkbox"/><div id="name-field-error" class="hello"></div></div>`, result, "should match")
}
func TestCheckboxWrapperClass(t *testing.T) {
	e := Checkbox("name", "Hello").WrapperClass("hello")
	result := e.Render()
	assert.Equalf(t, `<div class="hello"><label for="name-field">Hello</label><input id="name-field" name="name" type="checkbox"/><div id="name-field-error"></div></div>`, result, "should match")
}
func TestCheckboxLabelStyle(t *testing.T) {
	e := Checkbox("name", "Hello").LabelStyle("color: red")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field" style="color: red;">Hello</label><input id="name-field" name="name" type="checkbox"/><div id="name-field-error"></div></div>`, result, "should match")
}
func TestCheckboxInputStyle(t *testing.T) {
	e := Checkbox("name", "Hello").InputStyle("color: red")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="checkbox" style="color: red;"/><div id="name-field-error"></div></div>`, result, "should match")
}
func TestCheckboxErrorStyle(t *testing.T) {
	e := Checkbox("name", "Hello").ErrorStyle("color: red")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="checkbox"/><div id="name-field-error" style="color: red;"></div></div>`, result, "should match")
}
func TestCheckboxWrapperStyle(t *testing.T) {
	e := Checkbox("name", "Hello").WrapperStyle("color: red")
	result := e.Render()
	assert.Equalf(t, `<div style="color: red;"><label for="name-field">Hello</label><input id="name-field" name="name" type="checkbox"/><div id="name-field-error"></div></div>`, result, "should match")
}
func TestCheckboxText(t *testing.T) {
	e := Checkbox("name", "Hello").Error("Can't do that.")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="checkbox"/><div id="name-field-error">Can't do that.</div></div>`, result, "should match")
}

// TestCheckboxPost
func TestCheckboxPost(t *testing.T) {
	e := Checkbox("name", "Hello").Post("/url")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-post="/url" id="name-field" name="name" type="checkbox"/><div id="name-field-error"></div></div>`, result, "should match")
}

// TestCheckboxGet
func TestCheckboxGet(t *testing.T) {
	e := Checkbox("name", "Hello").Get("/url")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-get="/url" id="name-field" name="name" type="checkbox"/><div id="name-field-error"></div></div>`, result, "should match")
}

// TestCheckboxPut
func TestCheckboxPut(t *testing.T) {
	e := Checkbox("name", "Hello").Put("/url")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-put="/url" id="name-field" name="name" type="checkbox"/><div id="name-field-error"></div></div>`, result, "should match")
}

// TestCheckboxPatch
func TestCheckboxPatch(t *testing.T) {
	e := Checkbox("name", "Hello").Patch("/url")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-patch="/url" id="name-field" name="name" type="checkbox"/><div id="name-field-error"></div></div>`, result, "should match")
}
func TestCheckboxDelete(t *testing.T) {
	e := Checkbox("name", "Hello").Delete("/url")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-delete="/url" id="name-field" name="name" type="checkbox"/><div id="name-field-error"></div></div>`, result, "should match")
}
