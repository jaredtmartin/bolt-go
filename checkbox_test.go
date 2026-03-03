package bolt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestNewCheckbox(t *testing.T) {
	e := Checkbox("name", "Hello", "name")
	// log.Printf("e: %v\n", e)
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="checkbox" value="name"><div id="name-field-error"></div><span></span></div>`, result, "should match")
}
func TestCheckboxAttr(t *testing.T) {
	e := Checkbox("name", "Hello", "name")
	e.Input.Attr("hx-post", "/url")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-post="/url" id="name-field" name="name" type="checkbox" value="name"><div id="name-field-error"></div><span></span></div>`, result, "should match")
}

func TestCheckboxChecked(t *testing.T) {
	e := Checkbox("name", "Hello", "name").Checked()
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input checked="checked" id="name-field" name="name" type="checkbox" value="name"><div id="name-field-error"></div><span></span></div>`, result, "should match")
}
func TestCheckboxCheckedTrue(t *testing.T) {
	e := Checkbox("name", "Hello", "name").Checked(true)
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input checked="checked" id="name-field" name="name" type="checkbox" value="name"><div id="name-field-error"></div><span></span></div>`, result, "should match")
}
func TestCheckboxCheckedFalse(t *testing.T) {
	e := Checkbox("name", "Hello", "name").Checked(false)
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="checkbox" value="name"><div id="name-field-error"></div><span></span></div>`, result, "should match")
}
func TestCheckboxLabelClass(t *testing.T) {
	checkbox := Checkbox("name", "Hello", "name")
	checkbox.Label.Class("hello")
	result := checkbox.Render()
	assert.Equalf(t, `<div><label for="name-field" class="hello">Hello</label><input id="name-field" name="name" type="checkbox" value="name"><div id="name-field-error"></div><span></span></div>`, result, "should match")
}
func TestCheckboxInputClass(t *testing.T) {
	checkbox := Checkbox("name", "Hello", "name")
	checkbox.Input.Class("hello")
	result := checkbox.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="checkbox" value="name" class="hello"><div id="name-field-error"></div><span></span></div>`, result, "should match")
}
func TestCheckboxErrorClass(t *testing.T) {
	e := Checkbox("name", "Hello", "name")
	e.Error.Class("hello")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="checkbox" value="name"><div id="name-field-error" class="hello"></div><span></span></div>`, result, "should match")
}
func TestCheckboxWrapperClass(t *testing.T) {
	e := Checkbox("name", "Hello", "name").Class("hello")
	result := e.Render()
	assert.Equalf(t, `<div class="hello"><label for="name-field">Hello</label><input id="name-field" name="name" type="checkbox" value="name"><div id="name-field-error"></div><span></span></div>`, result, "should match")
}
func TestCheckboxLabelStyle(t *testing.T) {
	e := Checkbox("name", "Hello", "name")
	e.Label.Style("color: red")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field" style="color: red;">Hello</label><input id="name-field" name="name" type="checkbox" value="name"><div id="name-field-error"></div><span></span></div>`, result, "should match")
}
func TestCheckboxInputStyle(t *testing.T) {
	e := Checkbox("name", "Hello", "name")
	e.Input.Style("color: red")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="checkbox" value="name" style="color: red;"><div id="name-field-error"></div><span></span></div>`, result, "should match")
}
func TestCheckboxErrorStyle(t *testing.T) {
	e := Checkbox("name", "Hello", "name")
	e.Error.Style("color: red")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="checkbox" value="name"><div id="name-field-error" style="color: red;"></div><span></span></div>`, result, "should match")
}
func TestCheckboxWrapperStyle(t *testing.T) {
	e := Checkbox("name", "Hello", "name").Style("color: red")
	result := e.Render()
	assert.Equalf(t, `<div style="color: red;"><label for="name-field">Hello</label><input id="name-field" name="name" type="checkbox" value="name"><div id="name-field-error"></div><span></span></div>`, result, "should match")
}

func TestCheckboxGroup(t *testing.T) {
	flavors := []string{"chocolate", "vanilla", "strawberry"}
	e := CheckboxGroup("flavors", "Flavors", OptionStrings(flavors)).Style("color: red")
	result := e.Render()
	assert.Equalf(t, `<div style="color: red;"><div><label for="chocolate-field">flavors</label><input id="chocolate-field" name="chocolate" type="checkbox" value="chocolate"><div id="chocolate-field-error"></div><span></span></div><div><label for="vanilla-field">flavors</label><input id="vanilla-field" name="vanilla" type="checkbox" value="vanilla"><div id="vanilla-field-error"></div><span></span></div><div><label for="strawberry-field">flavors</label><input id="strawberry-field" name="strawberry" type="checkbox" value="strawberry"><div id="strawberry-field-error"></div><span></span></div></div>`, result, "should match")
}
