package bolt

import (
	"fmt"
)

type Field struct {
	DefaultElement
	Label Element
	Input Element
	Error Element
	Check Element
}

func NewField(name, label, value, tipe string) *Field {
	labelEl, inputEl, errorEl := initializeElement(name, label, value)
	field := Field{
		DefaultElement: NewDefaultElement("div"),
		Label:          labelEl,
		Input:          inputEl.Type(tipe),
		Error:          errorEl,
	}
	field.Children(labelEl, inputEl, errorEl)
	return &field
}
func TextField(name, label, value string) *Field {
	return NewField(name, label, value, "text")
}
func Radio(name, label, value string) *Field {
	field := NewField(name, label, value, "radio")
	return field
}
func Checkbox(name, label, value string) *Field {
	field := NewField(name, label, value, "checkbox")
	field.Check = Span("")
	field.Add(field.Check)
	return field
}
func Textarea(name, label, value string) *Field {
	field := NewField(name, label, "", "")
	field.Input.Text(value).Tag("textarea").RemoveAttr("type")
	return field
}
func Select[T any](name, label, value string, options []T, renderOption ...func(T, string) Element) *Field {
	field := NewField(name, label, value, "select")
	renderOpt := defaultRenderOption[T]
	if len(renderOption) > 0 {
		renderOpt = renderOption[0]
	}
	for i := range options {
		field.Input.Add(renderOpt(options[i], value))
	}
	field.Input.Tag("select").RemoveAttr("type")
	return field
}

// func (t *Field) Value(value string) *Field {
// 	t.Input.Value(value)
// 	return t
// }

// Sets the required and onblur attributes of the input element
func (t *Field) Required(required ...bool) *Field {
	if len(required) > 0 && !required[0] {
		t.Input.RemoveAttr("required").RemoveAttr("onblur")
	} else {
		t.Input.Attr("required", "required").Attr("onblur", "doFieldValidation(event)")
	}
	return t
}

// Sets the type of the input element to "email" if is_email is not specified or true
// Sets type to "text" if is_email is false
func (t *Field) Email(is_email ...bool) *Field {
	if len(is_email) > 0 && !is_email[0] {
		t.Input.Type("text")
	} else {
		t.Input.Type("email")
	}
	return t
}

// Sets the type of the input element to "password" if is_password is not specified or true
// Sets type to "text" if is_password is false
func (t *Field) Password(is_password ...bool) *Field {
	if len(is_password) > 0 && !is_password[0] {
		t.Input.Type("text")
	} else {
		t.Input.Type("password")
	}
	return t
}

// Sets the checked attribute of the input element
func (t *Field) Checked(value ...bool) *Field {
	if len(value) > 0 && !value[0] {
		t.Input.RemoveAttr("checked")
	} else {
		t.Input.Attr("checked", "checked")
	}
	return t
}

func initializeElement(name, label, value string) (Element, Element, Element) {
	id := name + "-field"
	var labelEl, inputEl, errorEl Element
	if label != "" {
		labelEl = Label(label).For(id)
	} else {
		labelEl = None()
	}
	inputEl = Input(name).Id(id)
	if value != "" {
		inputEl.Value(value)
	}
	errorEl = Div("").Id(id + "-error")
	return labelEl, inputEl, errorEl
}

// Default function to render an option in a select element
func defaultRenderOption[T any](opt T, value string) Element {
	str := fmt.Sprintf("%v", opt)
	item := NewElement("option").Value(str).Text(str)
	if value == str {
		item.Attr("selected", "true")
	}
	return item
}
