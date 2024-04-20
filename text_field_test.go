package bolt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestNewTextField(t *testing.T) {
	e := TextField("name", "Hello")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text"/><div id="name-field-error"></div></div>`, result, "should match")
}
func TestTextFieldAttr(t *testing.T) {
	e := TextField("name", "Hello").Attr("hx-post", "/url")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-post="/url" id="name-field" name="name" type="text"/><div id="name-field-error"></div></div>`, result, "should match")
}
func TestTextFieldValue(t *testing.T) {
	e := TextField("name", "Hello").Value("hello")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text" value="hello"/><div id="name-field-error"></div></div>`, result, "should match")
}
func TestTextFieldTarget(t *testing.T) {
	e := TextField("name", "Hello").Target("hello")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-target="hello" id="name-field" name="name" type="text"/><div id="name-field-error"></div></div>`, result, "should match")
}
func TestTextFieldSwap(t *testing.T) {
	e := TextField("name", "Hello").Swap("hello")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-swap="hello" id="name-field" name="name" type="text"/><div id="name-field-error"></div></div>`, result, "should match")
}
func TestTextFieldType(t *testing.T) {
	e := TextField("name", "Hello").Type("email")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="email"/><div id="name-field-error"></div></div>`, result, "should match")
}

func TestTextFieldName(t *testing.T) {
	e := TextField("name", "Hello").Name("name")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text"/><div id="name-field-error"></div></div>`, result, "should match")
}
func TestTextFieldNameExistingId(t *testing.T) {
	e := TextField("name", "Hello").Id("sup").Name("name")
	result := e.Render()
	assert.Equalf(t, `<div><label for="sup">Hello</label><input id="sup" name="name" type="text"/><div id="sup-error"></div></div>`, result, "should match")
}

func TestTextFieldId(t *testing.T) {
	e := TextField("name", "Hello").Id("name")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name">Hello</label><input id="name" name="name" type="text"/><div id="name-error"></div></div>`, result, "should match")
}

func TestTextFieldElementAndRender(t *testing.T) {
	e := TextField("name", "Hello").Element()
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text"/><div id="name-field-error"></div></div>`, result, "should match")
}
func TestTextFieldLabelClass(t *testing.T) {
	e := TextField("name", "Hello").LabelClass("hello")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field" class="hello">Hello</label><input id="name-field" name="name" type="text"/><div id="name-field-error"></div></div>`, result, "should match")
}

func TestTextFieldRequired(t *testing.T) {
	e := TextField("name", "Hello").Required()
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" onblur="doFieldValidation(event)" required="required" type="text"/><div id="name-field-error"></div></div>`, result, "should match")
}
func TestTextFieldRequiredFalse(t *testing.T) {
	e := TextField("name", "Hello").Required(true).Required(false)
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text"/><div id="name-field-error"></div></div>`, result, "should match")
}
func TestTextFieldCurrency(t *testing.T) {
	e := TextField("name", "Hello").Currency()
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input currency="true" id="name-field" name="name" onblur="doFieldValidation(event)" type="text"/><div id="name-field-error"></div></div>`, result, "should match")
}
func TestTextFieldCurrencyFalse(t *testing.T) {
	e := TextField("name", "Hello").Currency(true).Currency(false)
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text"/><div id="name-field-error"></div></div>`, result, "should match")
}

func TestTextFieldEmail(t *testing.T) {
	e := TextField("name", "Hello").Email()
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input email="true" id="name-field" name="name" onblur="doFieldValidation(event)" type="text"/><div id="name-field-error"></div></div>`, result, "should match")
}
func TestTextFieldEmailFalse(t *testing.T) {
	e := TextField("name", "Hello").Email(true).Email(false)
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text"/><div id="name-field-error"></div></div>`, result, "should match")
}

func TestTextFieldPassword(t *testing.T) {
	e := TextField("name", "Hello").Password()
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="password"/><div id="name-field-error"></div></div>`, result, "should match")
}

func TestTextFieldError(t *testing.T) {
	e := TextField("name", "Hello").Error("hello")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text"/><div id="name-field-error">hello</div></div>`, result, "should match")
}

func TestTextFieldPost(t *testing.T) {
	e := TextField("name", "Hello").Post("/url")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-post="/url" id="name-field" name="name" type="text"/><div id="name-field-error"></div></div>`, result, "should match")
}
func TestTextFieldLabel(t *testing.T) {
	e := TextField("name", "Hello").Label("hello")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">hello</label><input id="name-field" name="name" type="text"/><div id="name-field-error"></div></div>`, result, "should match")
}

func TestTextFieldGet(t *testing.T) {
	e := TextField("name", "Hello").Get("/url")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-get="/url" id="name-field" name="name" type="text"/><div id="name-field-error"></div></div>`, result, "should match")
}

func TestTextFieldPut(t *testing.T) {
	e := TextField("name", "Hello").Put("/url")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-put="/url" id="name-field" name="name" type="text"/><div id="name-field-error"></div></div>`, result, "should match")
}

func TestTextFieldPatch(t *testing.T) {
	e := TextField("name", "Hello").Patch("/url")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-patch="/url" id="name-field" name="name" type="text"/><div id="name-field-error"></div></div>`, result, "should match")
}

func TestTextFieldDelete(t *testing.T) {
	e := TextField("name", "Hello").Delete("/url")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-delete="/url" id="name-field" name="name" type="text"/><div id="name-field-error"></div></div>`, result, "should match")
}
func TestTextFieldInputClass(t *testing.T) {
	e := TextField("name", "Hello").InputClass("hello")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text" class="hello"/><div id="name-field-error"></div></div>`, result, "should match")
}
func TestTextFieldErrorClass(t *testing.T) {
	e := TextField("name", "Hello").ErrorClass("hello")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text"/><div id="name-field-error" class="hello"></div></div>`, result, "should match")
}
func TestTextFieldWrapperClass(t *testing.T) {
	e := TextField("name", "Hello").WrapperClass("hello")
	result := e.Render()
	assert.Equalf(t, `<div class="hello"><label for="name-field">Hello</label><input id="name-field" name="name" type="text"/><div id="name-field-error"></div></div>`, result, "should match")
}
func TestTextFieldLabelStyle(t *testing.T) {
	e := TextField("name", "Hello").LabelStyle("color: red")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field" style="color: red;">Hello</label><input id="name-field" name="name" type="text"/><div id="name-field-error"></div></div>`, result, "should match")
}
func TestTextFieldInputStyle(t *testing.T) {
	e := TextField("name", "Hello").InputStyle("color: red")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text" style="color: red;"/><div id="name-field-error"></div></div>`, result, "should match")
}
func TestTextFieldErrorStyle(t *testing.T) {
	e := TextField("name", "Hello").ErrorStyle("color: red")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text"/><div id="name-field-error" style="color: red;"></div></div>`, result, "should match")
}
func TestTextFieldWrapperStyle(t *testing.T) {
	e := TextField("name", "Hello").WrapperStyle("color: red")
	result := e.Render()
	assert.Equalf(t, `<div style="color: red;"><label for="name-field">Hello</label><input id="name-field" name="name" type="text"/><div id="name-field-error"></div></div>`, result, "should match")
}
