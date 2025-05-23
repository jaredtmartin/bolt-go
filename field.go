package bolt

//	type FieldElement struct {
//		_id     string
//		label   Element
//		input   Element
//		error   Element
//		wrapper Element
//	}
type fieldElementIndex int

const (
	FieldLabel fieldElementIndex = iota
	FieldInput
	FieldError
	FieldCheck
	FieldOptions
)

// Returns the label for a given value in a list of options
func GetOptionLabel(value string, options []Option) string {
	for _, v := range options {
		if v.Value == value {
			return v.Label
		}
	}
	return ""
}

// Option is a struct that represents an option in a select element.
type Option struct {
	Label string
	Value string
}

// FieldElement is a struct that represents a field element in a form.
type FieldElement struct {
	DefaultElement
	Renderer func(f *FieldElement) string
}

// Returns a label element with the given label and a for attribute set to the id of the input element
func createLabelElement(label string, id string) Element {
	return Label(label).Attr("for", id)
}

// Returns a FieldElement with label, input, and error(span) elements
func Field(name string, label string, value string) *FieldElement {
	id := name + "-field"
	field := &FieldElement{
		DefaultElement: NewDefaultElement("div"),
		Renderer:       defaultFieldRenderer,
	}
	var labelElement Element
	// Only used for checkboxes
	var checkElement Element
	if label != "" {
		labelElement = createLabelElement(label, id)
	}
	inputElement := Input(name).Type("text").Id(id)
	if value != "" {
		inputElement.Value(value)
	}
	field.Children(
		labelElement,
		inputElement,
		Div("").Id(id+"-error"),
		checkElement,
	)
	return field
}

// Returns a checkbox field with the given name and label
func Checkbox(name string, label string) *FieldElement {
	field := Field(name, label, "")
	field.GetInput().Type("checkbox")
	field.children[FieldCheck] = Span("")
	return field
}

// Default function to render an option in a select element
func defaultRenderOption(opt Option, value string) Element {
	item := NewElement("option").Value(opt.Value).Text(opt.Label)
	if value == opt.Value {
		item.Attr("selected", "true")
	}
	return item
}

// Returns a select field with the given name, label, value, and options
func Select(name string, label string, value string, options []Option, renderOption ...func(Option, string) Element) *FieldElement {
	renderOpt := defaultRenderOption
	if len(renderOption) > 0 {
		renderOpt = renderOption[0]
	}
	field := Field(name, label, "")
	input := field.GetInput().Tag("select").RemoveAttr("type")
	for i := 0; i < len(options); i++ {
		input.AddChild(renderOpt(options[i], value))
	}
	return field
}

// Sets the checked value of a checkbox
func (f *FieldElement) Checked(value ...bool) *FieldElement {
	if len(value) > 0 && !value[0] {
		f.GetInput().RemoveAttr("checked")
	} else {
		f.GetInput().Attr("checked", "checked")
	}
	return f
}

// Returns the input element of the field
func (f *FieldElement) GetInput() Element {
	return f.children[FieldInput]
}

// Sets the input element of the field
func (f *FieldElement) SetInput(e Element) *FieldElement {
	f.children[FieldInput] = e
	return f
}

// Sets the label text of the field
func (f *FieldElement) Label(label string) *FieldElement {
	l := f.GetLabel()
	// log.Printf("l: %v\n", l)
	if l == nil {
		// log.Printf("creating label")
		if label == "" {
			return f
		}
		f.children[FieldLabel] = createLabelElement(label, f.GetInput().GetAttr("id"))
	} else {
		if label == "" {
			f.children[FieldLabel] = nil
			return f
		} else {
			f.children[FieldLabel].Text(label)
		}
	}
	// log.Printf("l: %v\n", f.children)
	return f
}

// Returns the label element of the field
func (f *FieldElement) GetLabel() Element {
	return f.children[FieldLabel]
}

// Sets the label element of the field
func (f *FieldElement) SetLabel(l Element) *FieldElement {
	if l == nil {
		return f
	}
	f.children[FieldLabel] = l
	return f
}

// Returns the error element of the field
func (f *FieldElement) GetError() Element {
	return f.children[FieldError]
}

// Sets the error element of the field
func (f *FieldElement) SetError(e Element) *FieldElement {
	f.children[FieldError] = e
	return f
}

// Sets the type attribute on the input element
func (f *FieldElement) Type(tipe string) Element {
	f.GetInput().Type(tipe)
	return f
}

// Returns the value of the input field
func (f *FieldElement) GetValue() string {
	return f.GetInput().GetAttr("value")
}

// Returns the wrapper element of the field
// This is useful because it returns an Element instead of a FieldElement
func (f *FieldElement) GetWrapper() Element {
	return f
}

// Renders the field to an HTML string
func (f *FieldElement) Render() string {
	return f.Renderer(f)
}

// The default Renderer which renders the field to an HTML string
func defaultFieldRenderer(f *FieldElement) string {
	return f.DefaultElement.Render()
}

// func (f *FieldElement) Class

// func (f *FieldElement) Label(label string) Element {
// 	f.label = *f.label.Text(label)
// 	return f
// }

// Sets the id attribute of the input element, updates the for attribute on the label, and the id on the error span
func (f *FieldElement) Id(id string) Element {
	f.GetInput().Id(id)
	label := f.GetLabel()
	if label != nil {
		label.Attr("for", id)
	}
	f.GetError().Id(id + "-error")
	return f
}

// Sets the name attribute of the input element. Also updates the id on the field to name + "-field"
func (f *FieldElement) Name(name string) Element {
	input := f.GetInput()
	oldName := input.GetAttr("name")
	// log.Printf("oldName: %v\n", oldName)
	oldId := input.GetId()
	// log.Printf("oldId: %v\n", oldId)
	// log.Printf("oldId == oldName: %v\n", oldId == oldName+"-field")
	input.Name(name)
	if oldId == oldName+"-field" {
		// log.Printf("setting id")
		return f.Id(name + "-field")
	}
	return f
}

// Sets the value attribute of the input element
func (f *FieldElement) Value(value string) Element {
	f.GetInput().Value(value)
	return f
}

// Sets the required and onblur attributes of the input element
func (f *FieldElement) Required(required ...bool) *FieldElement {
	if len(required) > 0 && !required[0] {
		f.GetInput().RemoveAttr("required").RemoveAttr("onblur")
	} else {
		f.GetInput().Attr("required", "required").Attr("onblur", "doFieldValidation(event)")
	}
	return f
}

// Sets the type of the input element to "email" if is_email is not specified or true
// Sets type to "text" if is_email is false
func (f *FieldElement) Email(is_email ...bool) *FieldElement {
	if len(is_email) > 0 && !is_email[0] {
		f.GetInput().Type("text")
	} else {
		f.GetInput().Type("email")
	}
	return f
}

// Sets the type of the input element to "password" if is_password is not specified or true
// Sets type to "text" if is_password is false
func (f *FieldElement) Password(is_password ...bool) *FieldElement {
	if len(is_password) > 0 && !is_password[0] {
		f.GetInput().Type("text")
	} else {
		f.GetInput().Type("password")
	}
	return f
}

// func (e *DefaultElement) Debug(prefix ...string) *FieldElement {
// 	s := ""
// 	if len(prefix) > 0 {
// 		s = prefix[0]
// 	}
// 	log.Printf("%stag: %s, text: %s\n", s, e.tag, e.text)
// 	log.Printf("%sclasses: %v\n", s, e.classes)
// 	log.Printf("%sstyles: %v\n", s, e.styles)
// 	log.Printf("%sattributes: %v\n", s, e.attributes)
// 	log.Printf("%schildren: \n", s)
// 	for _, child := range e.children {
// 		if child != nil {
// 			child.Debug(s + "  ")
// 		} else {
// 			log.Printf("%s  nil\n", s)
// 		}

// 	}
// 	return e
// }

//	func (f *FieldElement) Pattern(pattern string) Element {
//		f.input = *f.input.Attr("pattern", pattern)
//		return f
//	}
//
//	func (f *FieldElement) Accept(accept string) Element {
//		f.input = *f.input.Attr("accept", accept)
//		return f
//	}
//
//	func (f *FieldElement) Autocapitalize(autocapitalize string) Element {
//		f.input = *f.input.Attr("autocapitalize", autocapitalize)
//		return f
//	}
//
//	func (f *FieldElement) Autocomplete(autocomplete string) Element {
//		f.input = *f.input.Attr("autocomplete", autocomplete)
//		return f
//	}
//
//	func (f *FieldElement) Max(max string) Element {
//		f.input = *f.input.Attr("max", max)
//		return f
//	}
//
//	func (f *FieldElement) Min(min string) Element {
//		f.input = *f.input.Attr("min", min)
//		return f
//	}
//
//	func (f *FieldElement) MaxLength(maxlength string) Element {
//		f.input = *f.input.Attr("maxlength", maxlength)
//		return f
//	}
//
//	func (f *FieldElement) MinLength(minlength string) Element {
//		f.input = *f.input.Attr("minlength", minlength)
//		return f
//	}
//
//	func (f *FieldElement) Color(is_color ...bool) Element {
//		if len(is_color) > 0 && !is_color[0] {
//			f.input = *f.input.Type("text")
//		} else {
//			f.input = *f.input.Type("color")
//		}
//		return f
//	}
//
//	func (f *FieldElement) Date(is_date ...bool) Element {
//		if len(is_date) > 0 && !is_date[0] {
//			f.input = *f.input.Type("text")
//		} else {
//			f.input = *f.input.Type("date")
//		}
//		return f
//	}
//
//	func (f *FieldElement) Datetime(is_datetime ...bool) Element {
//		if len(is_datetime) > 0 && !is_datetime[0] {
//			f.input = *f.input.Type("text")
//		} else {
//			f.input = *f.input.Type("datetime-local") // datetime is deprecated
//		}
//		return f
//	}
//
//	func (f *FieldElement) File(is_file ...bool) Element {
//		if len(is_file) > 0 && !is_file[0] {
//			f.input = *f.input.Type("text")
//		} else {
//			f.input = *f.input.Type("file")
//		}
//		return f
//	}
//
//	func (f *FieldElement) Hidden(is_hidden ...bool) Element {
//		if len(is_hidden) > 0 && !is_hidden[0] {
//			f.input = *f.input.Type("text")
//		} else {
//			f.input = *f.input.Type("hidden")
//		}
//		return f
//	}
//
//	func (f *FieldElement) Month(is_month ...bool) Element {
//		if len(is_month) > 0 && !is_month[0] {
//			f.input = *f.input.Type("text")
//		} else {
//			f.input = *f.input.Type("month")
//		}
//		return f
//	}
//
//	func (f *FieldElement) Number(is_number ...bool) Element {
//		if len(is_number) > 0 && !is_number[0] {
//			f.input = *f.input.Type("text")
//		} else {
//			f.input = *f.input.Type("number")
//		}
//		return f
//	}
//
//	func (f *FieldElement) Tel(is_tel ...bool) Element {
//		if len(is_tel) > 0 && !is_tel[0] {
//			f.input = *f.input.Type("text")
//		} else {
//			f.input = *f.input.Type("tel")
//		}
//		return f
//	}
//
//	func (f *FieldElement) Time(is_time ...bool) Element {
//		if len(is_time) > 0 && !is_time[0] {
//			f.input = *f.input.Type("text")
//		} else {
//			f.input = *f.input.Type("time")
//		}
//		return f
//	}
//
//	func (f *FieldElement) Url(is_url ...bool) Element {
//		if len(is_url) > 0 && !is_url[0] {
//			f.input = *f.input.Type("text")
//		} else {
//			f.input = *f.input.Type("url")
//		}
//		return f
//	}
//
//	func (f *FieldElement) Week(is_week ...bool) Element {
//		if len(is_week) > 0 && !is_week[0] {
//			f.input = *f.input.Type("text")
//		} else {
//			f.input = *f.input.Type("week")
//		}
//		return f
//	}
//
//	func (f *FieldElement) Search(is_search ...bool) Element {
//		if len(is_search) > 0 && !is_search[0] {
//			f.input = *f.input.Type("text")
//		} else {
//			f.input = *f.input.Type("search")
//		}
//		return f
//	}
//
//	func (f *FieldElement) Password(is_password ...bool) Element {
//		if len(is_password) > 0 && !is_password[0] {
//			f.input = *f.input.Type("text")
//		} else {
//			f.input = *f.input.Type("password")
//		}
//		return f
//	}
//
//	func (f *FieldElement) Range(min int, max int) Element {
//		f.input = *f.input.Type("range").Attr("min", strconv.Itoa(min)).Attr("max", strconv.Itoa(max))
//		return f
//	}
//
//	func (f *FieldElement) Placeholder(placeholder string) Element {
//		f.input = *f.input.Attr("placeholder", placeholder)
//		return f
//	}
//
//	func (f *FieldElement) Step(step string) Element {
//		f.input = *f.input.Attr("step", step)
//		return f
//	}

// Sets the Error message of the element
func (f *FieldElement) Error(err string) *FieldElement {
	f.GetError().Text(err)
	return f
}

// Sets the hx-target attribute on the input element
func (f *FieldElement) HXTarget(value string) Element {
	f.GetInput().HXTarget(value)
	return f
}

//	func (f *FieldElement) Indicator(selector string) Element {
//		f.input = *f.input.Indicator(selector)
//		return f
//	}
//
//	func (f *FieldElement) Class(classes ...string) Element {
//		f.wrapper = *f.wrapper.Class(classes)
//		return f
//	}
//
//	func (f *FieldElement) Trigger(value string) Element {
//		f.input = *f.input.Trigger(value)
//		return f
//	}
//
//	func (f *FieldElement) PushUrl(pushUrl ...bool) Element {
//		f.input = *f.input.PushUrl(pushUrl...)
//		return f
//	}
// func (f *FieldElement) Attr(key string, value string) *FieldElement {
// 	f.GetInput().Attr(key, value)
// 	return f
// }

// Sets the hx-post attribute on the input element
func (f *FieldElement) HXPost(value string) Element {
	f.GetInput().HXPost(value)
	return f
}

// Sets the hx-get attribute on the input element
func (f *FieldElement) HXGet(value string) Element {
	f.GetInput().HXGet(value)
	return f
}

// Sets the hx-put attribute on the input element
func (f *FieldElement) HXPut(value string) Element {
	f.GetInput().HXPut(value)
	return f
}

// Sets the hx-patch attribute on the input element
func (f *FieldElement) HXPatch(value string) Element {
	f.GetInput().HXPatch(value)
	return f
}

// Sets the hx-delete attribute on the input element
func (f *FieldElement) HXDelete(value string) Element {
	f.GetInput().HXDelete(value)
	return f
}

// Sets the hx-swap attribute on the input element
func (f *FieldElement) HXSwap(value string) Element {
	f.GetInput().HXSwap(value)
	return f
}

// func (f *FieldElement) Textarea() Element {
// 	f.input = *f.input.Tag("textarea")
// 	return f
// }
