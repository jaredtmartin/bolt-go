package bolt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// for a valid return value.
func TestNewTextField(t *testing.T) {
	e := TextField("name", "Hello", "")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestNewFieldWithoutLabel(t *testing.T) {
	e := TextField("name", "", "")
	result := e.Render()
	assert.Equalf(t, `<div><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
}

func TestFieldAttr(t *testing.T) {
	e := TextField("name", "Hello", "")
	e.Input.Attr("hx-post", "/url")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-post="/url" id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestFieldText(t *testing.T) {
	e := TextField("name", "Hello", "").Text("hello")
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text"><div id="name-field-error"></div>hello</div>`, result, "should match")
}
func TestFieldWithNoLabel(t *testing.T) {
	e := TextField("name", "", "")
	result := e.Render()
	assert.Equalf(t, `<div><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
}

// func TestFieldRange(t *testing.T) {
// 	e := TextField("range", "How much?").Range(1, 10)
// 	result := e.Render()
// 	assert.Equalf(t, `<div><label for="range-field">How much?</label><input id="range-field" max="10" min="1" name="range" type="range"/><div id="range-field-error"></div></div>`, result, "should match")
// }

// func TestFieldElementAndRender(t *testing.T) {
// 	e := TextField("name", "Hello").Element()
// 	result := e.Render()
// 	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
// }
// func TestFieldLabelClass(t *testing.T) {
// 	e := TextField("name", "Hello").LabelClass("hello")
// 	result := e.Render()
// 	assert.Equalf(t, `<div><label for="name-field" class="hello">Hello</label><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
// }

// func TestFieldRequired(t *testing.T) {
// 	e := TextField("name", "Hello").Required()
// 	result := e.Render()
// 	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" onblur="doFieldValidation(event)" required="required" type="text"><div id="name-field-error"></div></div>`, result, "should match")
// }
// func TestFieldRequiredFalse(t *testing.T) {
// 	e := TextField("name", "Hello").Required(true).Required(false)
// 	result := e.Render()
// 	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
// }

// func TestFieldCurrency(t *testing.T) {
// 	e := TextField("name", "Hello").Currency()
// 	result := e.Render()
// 	assert.Equalf(t, `<div><label for="name-field">Hello</label><input currency="true" id="name-field" name="name" onblur="doFieldValidation(event)" type="text"><div id="name-field-error"></div></div>`, result, "should match")
// }
// func TestFieldCurrencyFalse(t *testing.T) {
// 	e := TextField("name", "Hello").Currency(true).Currency(false)
// 	result := e.Render()
// 	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
// }

func TestFieldEmail(t *testing.T) {
	e := TextField("name", "Hello", "").Email()
	result := e.Render()

	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="email"><div id="name-field-error"></div></div>`, result, "should match")
}

func TestFieldEmailFalse(t *testing.T) {
	e := TextField("name", "Hello", "").Email(true).Email(false)
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
}

func TestFieldPassword(t *testing.T) {
	e := TextField("name", "Hello", "").Password()
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="password"><div id="name-field-error"></div></div>`, result, "should match")
}

// func TestFieldPost(t *testing.T) {
// 	e := TextField("name", "Hello", "").Post("/url")
// 	result := e.Render()
// 	assert.Equalf(t, `<div><label for="name-field">Hello</label><input hx-post="/url" id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
// }

// func TestFieldLabel(t *testing.T) {
// 	e := TextField("name", "Hello").Label("hello")
// 	result := e.Render()
// 	assert.Equalf(t, `<div><label for="name-field">hello</label><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
// }

// func TestSelectWithDefaultOptionRenderer(t *testing.T) {
// 	opts := []Option{
// 		{Label: "Red", Value: "red"},
// 		{Label: "Green", Value: "green"},
// 		{Label: "Blue", Value: "blue"},
// 	}
// 	field := Select("color", "Color", "red", opts)
// 	result := field.Render()
// 	assert.Equalf(t, `<div><label for="color-field">Color</label><select id="color-field" name="color"><option selected="true" value="red">Red</option><option value="green">Green</option><option value="blue">Blue</option></select><div id="color-field-error"></div></div>`, result, "should match")
// }

//	func TestSelectWithCustomOptionRenderer(t *testing.T) {
//		opts := []Option{
//			{Label: "Red", Value: "red"},
//			{Label: "Green", Value: "green"},
//			{Label: "Blue", Value: "blue"},
//		}
//		renderOpt := func(opt Option, value string) Element {
//			return Div("").Text(opt.Label)
//		}
//		field := Select("color", "Color", "red", opts, renderOpt)
//		result := field.Render()
//		assert.Equalf(t, `<div><label for="color-field">Color</label><select id="color-field" name="color"><div>Red</div><div>Green</div><div>Blue</div></select><div id="color-field-error"></div></div>`, result, "should match")
//	}
// func TestSelectWithNoValue(t *testing.T) {
// 	opts := []Option{
// 		{Label: "Red", Value: "red"},
// 		{Label: "Green", Value: "green"},
// 		{Label: "Blue", Value: "blue"},
// 	}
// 	field := Select("color", "Color", "", opts)
// 	result := field.Render()
// 	assert.Equalf(t, `<div><label for="color-field">Color</label><select id="color-field" name="color"><option value="red">Red</option><option value="green">Green</option><option value="blue">Blue</option></select><div id="color-field-error"></div></div>`, result, "should match")
// }

// func TestCustomFieldRendering(t *testing.T) {
// 	e := TextField("name", "Hello", "")
// 	e.Renderer = func(e *FieldElement) string {
// 		renderedLabel := e.GetLabel()
// 		renderedInput := e.Input
// 		return e.RenderWithContent("blah", renderedInput, renderedLabel)
// 	}
// 	result := e.Render()
// 	assert.Equalf(t, `<div><input id="name-field" name="name" type="text"><label for="name-field">Hello</label>blah</div>`, result, "should match")
// }

//	func TestFieldInputClass(t *testing.T) {
//		e := TextField("name", "Hello").InputClass("hello")
//		result := e.Render()
//		assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text" class="hello"/><div id="name-field-error"></div></div>`, result, "should match")
//	}
//
//	func TestFieldErrorClass(t *testing.T) {
//		e := TextField("name", "Hello").ErrorClass("hello")
//		result := e.Render()
//		assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text"><div id="name-field-error" class="hello"></div></div>`, result, "should match")
//	}
//
//	func TestFieldWrapperClass(t *testing.T) {
//		e := TextField("name", "Hello").WrapperClass("hello")
//		result := e.Render()
//		assert.Equalf(t, `<div class="hello"><label for="name-field">Hello</label><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
//	}
//
//	func TestFieldLabelStyle(t *testing.T) {
//		e := TextField("name", "Hello").LabelStyle("color: red")
//		result := e.Render()
//		assert.Equalf(t, `<div><label for="name-field" style="color: red;">Hello</label><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
//	}
//
//	func TestFieldInputStyle(t *testing.T) {
//		e := TextField("name", "Hello").InputStyle("color: red")
//		result := e.Render()
//		assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text" style="color: red;"/><div id="name-field-error"></div></div>`, result, "should match")
//	}
//
//	func TestFieldErrorStyle(t *testing.T) {
//		e := TextField("name", "Hello").ErrorStyle("color: red")
//		result := e.Render()
//		assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text"><div id="name-field-error" style="color: red;"></div></div>`, result, "should match")
//	}
//
//	func TestFieldWrapperStyle(t *testing.T) {
//		e := TextField("name", "Hello").WrapperStyle("color: red")
//		result := e.Render()
//		assert.Equalf(t, `<div style="color: red;"><label for="name-field">Hello</label><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
//	}
//
//	func TestFieldTextarea(t *testing.T) {
//		e := TextField("name", "Hello").Value("red\ngreen").Textarea()
//		result := e.Render()
//		assert.Equalf(t, `<div><label for="name-field">Hello</label><textarea id="name-field" name="name" type="text">red`+"\n"+`green</textarea><div id="name-field-error"></div></div>`, result, "should match")
//	}

func TestNoLabel(t *testing.T) {
	e := Label("")
	result := e.Render()
	assert.Equalf(t, `<label></label>`, result, "should match")
}

func TestRequired(t *testing.T) {
	e := TextField("name", "Hello", "")
	e.Required()
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" onblur="doFieldValidation(event)" required="required" type="text"><div id="name-field-error"></div></div>`, result, "should match")

	e.Required(false)
	result = e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
}
func TestPassword(t *testing.T) {
	e := TextField("name", "Hello", "")
	e.Password()
	result := e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="password"><div id="name-field-error"></div></div>`, result, "should match")

	e.Password(false)
	result = e.Render()
	assert.Equalf(t, `<div><label for="name-field">Hello</label><input id="name-field" name="name" type="text"><div id="name-field-error"></div></div>`, result, "should match")
}

// func TestGetValue(t *testing.T) {
// 	e := TextField("hello", "Hello", "123")
// 	result := e.GetValue()
// 	assert.Equalf(t, "123", result, "should match")

// }
