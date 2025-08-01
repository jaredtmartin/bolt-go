package bolt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestNewField(t *testing.T) {
	e := Field("name", "Hello", "")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestNewFieldWithoutLabel(t *testing.T) {
	e := Field("name", "", "")
	result := e.Render()
	assert.Equalf(t, `<div><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
}

func TestAddLabelToFieldLater(t *testing.T) {
	e := Field("name", "", "")
	result := e.Render()
	assert.Equalf(t, `<div><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
	e.Label("Blue")
	// assert.Equalf(t, `<div><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
	// e.Label("Blue")
	result = e.Render()
	assert.Equalf(t, `<div><label for="name-field">Blue</label><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestRemoveLabelFromFieldLater(t *testing.T) {
	e := Field("name", "Blue", "")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Blue</label><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
	e.Label("")
	// assert.Equalf(t, `<div><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
	// e.Label("Blue")
	result = e.Render()
	assert.Equalf(t, `<div><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestFieldAttr(t *testing.T) {
	e := Field("name", "Hello", "")
	e.GetInput().Attr("hx-post", "/url")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-post="/url" id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestFieldValue(t *testing.T) {
	e := Field("name", "Hello", "first")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text" value="first"><div id="name-field-error"></div></div>`, result, "should match")
	e.Value("hello")
	result = e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text" value="hello"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestFieldTarget(t *testing.T) {
	e := Field("name", "Hello", "").HXTarget("hello")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-target="hello" id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestFieldSwap(t *testing.T) {
	e := Field("name", "Hello", "").HXSwap("hello")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-swap="hello" id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestFieldType(t *testing.T) {
	e := Field("name", "Hello", "").Type("email")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="email"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestFieldText(t *testing.T) {
	e := Field("name", "Hello", "").Text("hello")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text"><div id="name-field-error"></div>hello</div>`, result, "should match")
}
func TestFieldName(t *testing.T) {
	e := Field("name", "Hello", "").Name("name")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestFieldNameExistingId(t *testing.T) {
	e := Field("name", "Hello", "").Id("sup").Name("name")
	result := e.Render()
	assert.Equalf(t, `<div><label for="sup">Hello</label><input id="sup" name="name" type="text"><div id="sup-error"></div></div>`, result, "should match")
}
func TestFieldLabel(t *testing.T) {
	e := Field("name", "Hello", "").Label("Green")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Green</label><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestFieldWithNoLabel(t *testing.T) {
	e := Field("name", "", "")
	result := e.Render()
	assert.Equalf(t, `<div><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestFieldId(t *testing.T) {
	e := Field("name", "Hello", "").Id("name")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name">Hello</label><input id="name" name="name" type="text"><div id="name-error"></div></div>`, result, "should match")
}

// func TestFieldRange(t *testing.T) {
// 	e := Field("range", "How much?").Range(1, 10)
// 	result := e.Render()
// 	assert.Equalf(t, `<div><label for="range-field">How much?</label><input id="range-field" max="10" min="1" name="range" type="range"/><div id="range-field-error"></div></div>`, result, "should match")
// }

// func TestFieldElementAndRender(t *testing.T) {
// 	e := Field("name", "Hello").Element()
// 	result := e.Render()
// 	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
// }
// func TestFieldLabelClass(t *testing.T) {
// 	e := Field("name", "Hello").LabelClass("hello")
// 	result := e.Render()
// 	assert.Equalf(t, `<div><label for="name-field" class="hello">Hello</label><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
// }

// func TestFieldRequired(t *testing.T) {
// 	e := Field("name", "Hello").Required()
// 	result := e.Render()
// 	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" onblur="doFieldValidation(event)" required="required" type="text"><div id="name-field-error"></div></div>`, result, "should match")
// }
// func TestFieldRequiredFalse(t *testing.T) {
// 	e := Field("name", "Hello").Required(true).Required(false)
// 	result := e.Render()
// 	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
// }

// func TestFieldCurrency(t *testing.T) {
// 	e := Field("name", "Hello").Currency()
// 	result := e.Render()
// 	assert.Equalf(t, `<div><label for="name-field">Hello</label><input currency="true" id="name-field" name="name" onblur="doFieldValidation(event)" type="text"><div id="name-field-error"></div></div>`, result, "should match")
// }
// func TestFieldCurrencyFalse(t *testing.T) {
// 	e := Field("name", "Hello").Currency(true).Currency(false)
// 	result := e.Render()
// 	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
// }

func TestFieldEmail(t *testing.T) {
	e := Field("name", "Hello", "").Email()
	result := e.Render()

	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="email"><div id="name-field-error"></div></div>`, result, "should match")
}

func TestFieldEmailFalse(t *testing.T) {
	e := Field("name", "Hello", "").Email(true).Email(false)
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
}

func TestFieldPassword(t *testing.T) {
	e := Field("name", "Hello", "").Password()
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="password"><div id="name-field-error"></div></div>`, result, "should match")
}

func TestFieldError(t *testing.T) {
	e := Field("name", "Hello", "").Error("hello")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text"><div id="name-field-error">hello</div></div>`, result, "should match")
}

// func TestFieldPost(t *testing.T) {
// 	e := Field("name", "Hello", "").Post("/url")
// 	result := e.Render()
// 	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-post="/url" id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
// }

// func TestFieldLabel(t *testing.T) {
// 	e := Field("name", "Hello").Label("hello")
// 	result := e.Render()
// 	assert.Equalf(t, `<div><label for="name-field">hello</label><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
// }

func TestFieldGet(t *testing.T) {
	e := Field("name", "Hello", "").HXGet("/url")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-get="/url" id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
}

func TestFieldPut(t *testing.T) {
	e := Field("name", "Hello", "").HXPut("/url")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-put="/url" id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
}

func TestFieldPatch(t *testing.T) {
	e := Field("name", "Hello", "").HXPatch("/url")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-patch="/url" id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
}

func TestFieldDelete(t *testing.T) {
	e := Field("name", "Hello", "").HXDelete("/url")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-delete="/url" id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestSelectWithDefaultOptionRenderer(t *testing.T) {
	opts := []Option{
		{Label: "Red", Value: "red"},
		{Label: "Green", Value: "green"},
		{Label: "Blue", Value: "blue"},
	}
	field := Select("color", "Color", "red", opts)
	result := field.Render()
	assert.Equalf(t, `<div><label for="color-field">Color</label><select id="color-field" name="color"><option selected="true" value="red">Red</option><option value="green">Green</option><option value="blue">Blue</option></select><div id="color-field-error"></div></div>`, result, "should match")
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
	assert.Equalf(t, `<div><label for="color-field">Color</label><select id="color-field" name="color"><div>Red</div><div>Green</div><div>Blue</div></select><div id="color-field-error"></div></div>`, result, "should match")
}
func TestSelectWithNoValue(t *testing.T) {
	opts := []Option{
		{Label: "Red", Value: "red"},
		{Label: "Green", Value: "green"},
		{Label: "Blue", Value: "blue"},
	}
	field := Select("color", "Color", "", opts)
	result := field.Render()
	assert.Equalf(t, `<div><label for="color-field">Color</label><select id="color-field" name="color"><option value="red">Red</option><option value="green">Green</option><option value="blue">Blue</option></select><div id="color-field-error"></div></div>`, result, "should match")
}

func TestCustomFieldRendering(t *testing.T) {
	e := Field("name", "Hello", "")
	e.Renderer = func(e *FieldElement) string {
		renderedLabel := e.GetLabel()
		renderedInput := e.GetInput()
		return e.RenderWithContent("blah", renderedInput, renderedLabel)
	}
	result := e.Render()
	assert.Equalf(t, `<div><input id="name-field" name="name" type="text"><label for="name-field">Hello</label>blah</div>`, result, "should match")
}

//	func TestFieldInputClass(t *testing.T) {
//		e := Field("name", "Hello").InputClass("hello")
//		result := e.Render()
//		assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text" class="hello"/><div id="name-field-error"></div></div>`, result, "should match")
//	}
//
//	func TestFieldErrorClass(t *testing.T) {
//		e := Field("name", "Hello").ErrorClass("hello")
//		result := e.Render()
//		assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text"><div id="name-field-error" class="hello"></div></div>`, result, "should match")
//	}
//
//	func TestFieldWrapperClass(t *testing.T) {
//		e := Field("name", "Hello").WrapperClass("hello")
//		result := e.Render()
//		assert.Equalf(t, `<div class="hello"><label for="name-field">Hello</label><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
//	}
//
//	func TestFieldLabelStyle(t *testing.T) {
//		e := Field("name", "Hello").LabelStyle("color: red")
//		result := e.Render()
//		assert.Equalf(t, `<div><label for="name-field" style="color: red;">Hello</label><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
//	}
//
//	func TestFieldInputStyle(t *testing.T) {
//		e := Field("name", "Hello").InputStyle("color: red")
//		result := e.Render()
//		assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text" style="color: red;"/><div id="name-field-error"></div></div>`, result, "should match")
//	}
//
//	func TestFieldErrorStyle(t *testing.T) {
//		e := Field("name", "Hello").ErrorStyle("color: red")
//		result := e.Render()
//		assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text"><div id="name-field-error" style="color: red;"></div></div>`, result, "should match")
//	}
//
//	func TestFieldWrapperStyle(t *testing.T) {
//		e := Field("name", "Hello").WrapperStyle("color: red")
//		result := e.Render()
//		assert.Equalf(t, `<div style="color: red;"><label for="name-field">Hello</label><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
//	}
//
//	func TestFieldTextarea(t *testing.T) {
//		e := Field("name", "Hello").Value("red\ngreen").Textarea()
//		result := e.Render()
//		assert.Equalf(t, `<div><label for="name-field">Hello</label><textarea id="name-field" name="name" type="text">red`+"\n"+`green</textarea><div id="name-field-error"></div></div>`, result, "should match")
//	}
func TestGetOptionLabel(t *testing.T) {
	opts := []Option{
		{Label: "Red", Value: "red"},
		{Label: "Green", Value: "green"},
		{Label: "Blue", Value: "blue"},
	}

	result := GetOptionLabel("green", opts)
	assert.Equal(t, "Green", result)
}

func TestGetOptionLabelEmptyOptions(t *testing.T) {
	var opts []Option
	result := GetOptionLabel("anything", opts)
	assert.Equal(t, "", result)
}

func TestGetOptionLabelValueNotFound(t *testing.T) {
	opts := []Option{
		{Label: "Red", Value: "red"},
		{Label: "Green", Value: "green"},
	}
	result := GetOptionLabel("yellow", opts)
	assert.Equal(t, "", result)
}

func TestGetOptionLabelWithEmptyValue(t *testing.T) {
	opts := []Option{
		{Label: "Red", Value: "red"},
		{Label: "Empty", Value: ""},
		{Label: "Blue", Value: "blue"},
	}
	result := GetOptionLabel("", opts)
	assert.Equal(t, "Empty", result)
}

func TestGetOptionLabelWithDuplicateValues(t *testing.T) {
	opts := []Option{
		{Label: "First Red", Value: "red"},
		{Label: "Second Red", Value: "red"},
	}
	result := GetOptionLabel("red", opts)
	assert.Equal(t, "First Red", result)
}
func TestSetInput(t *testing.T) {
	e := Field("name", "Hello", "")
	e.SetInput(Input("name").Type("number"))
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input name="name" type="number"><div id="name-field-error"></div></div>`, result, "should match")
}

func TestSetInputNil(t *testing.T) {
	e := Field("name", "Hello", "")
	e.SetInput(nil)
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><div id="name-field-error"></div></div>`, result, "should match")
}
func TestSetLabel(t *testing.T) {
	e := Field("name", "Hello", "")
	e.SetLabel(Label("name-field"))
	result := e.Render()
	assert.Equalf(t, `<div><label>name-field</label><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
}

func TestSetLabelNil(t *testing.T) {
	e := Field("name", "Hello", "")
	originalRender := e.Render()
	e.SetLabel(nil)
	result := e.Render()
	assert.Equalf(t, originalRender, result, "should remain unchanged when setting nil label")
}
func TestNoLabel(t *testing.T) {
	e := Label("")
	result := e.Render()
	assert.Equalf(t, `<label></label>`, result, "should match")
}

func TestSetLabelCustomElement(t *testing.T) {
	e := Field("name", "Hello", "")
	customLabel := Div("").Text("Custom Label").Class("custom-label")
	e.SetLabel(customLabel)
	result := e.Render()
	assert.Equalf(t, `<div><div class="custom-label">Custom Label</div><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestSetError(t *testing.T) {
	e := Field("name", "Hello", "")
	e.SetError(P("Something isnt right!"))
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text"><p>Something isnt right!</p></div>`, result, "should match")
}

func TestSetErrorNil(t *testing.T) {
	e := Field("name", "Hello", "")
	e.SetError(nil)
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text"></div>`, result, "should match")
}
func TestRequired(t *testing.T) {
	e := Field("name", "Hello", "")
	e.Required()
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" onblur="doFieldValidation(event)" required="required" type="text"><div id="name-field-error"></div></div>`, result, "should match")

	e.Required(false)
	result = e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestPassword(t *testing.T) {
	e := Field("name", "Hello", "")
	e.Password()
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="password"><div id="name-field-error"></div></div>`, result, "should match")

	e.Password(false)
	result = e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestGetWrapper(t *testing.T) {
	e := Field("name", "Hello", "")
	wrapper := e.GetWrapper()

	// Test that wrapper is the same instance as the field
	assert.Equal(t, e, wrapper)

	// Verify wrapper modifications affect the original field
	wrapper.Class("test-class")
	result := e.Render()
	assert.Equalf(t, `<div class="test-class"><label for="name-field">Hello</label><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")

}
func TestRemoveLabel(t *testing.T) {
	e := Field("hello", "", "").Label("")
	result := e.Render()
	assert.Equalf(t, `<div><input id="hello-field" name="hello" type="text"><div id="hello-field-error"></div></div>`, result, "should match")
}

func TestGetValue(t *testing.T) {
	e := Field("hello", "Hello", "123")
	result := e.GetValue()
	assert.Equalf(t, "123", result, "should match")

}
