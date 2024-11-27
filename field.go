package bolt

import "log"

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
)

type FieldElement struct {
	DefaultElement
}

func createLabelElement(label string, id string) Element {
	return Label(label).Attr("for", id)
}
func Field(name string, label string) *FieldElement {
	id := name + "-field"
	field := &FieldElement{DefaultElement: NewDefaultElement("div")}
	var labelElement Element
	// Only used for checkboxes
	var checkElement Element
	if label != "" {
		labelElement = createLabelElement(label, id)
	}
	field.Children(
		labelElement,
		Input(name).Type("text").Id(id),
		Div().Id(id+"-error"),
		checkElement,
	)
	return field
}
func Checkbox(name string, label string) *FieldElement {
	field := Field(name, label)
	field.GetInput().Type("checkbox")
	field.children[FieldCheck] = Span("")
	return field
}
func (f *FieldElement) Checked(value ...bool) *FieldElement {
	if len(value) > 0 && !value[0] {
		f.GetInput().RemoveAttr("checked")
	} else {
		f.GetInput().Attr("checked", "checked")
	}
	return f
}
func (f *FieldElement) GetInput() Element {
	return f.children[FieldInput]
}
func (f *FieldElement) SetInput(e Element) *FieldElement {
	f.children[FieldInput] = e
	return f
}
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
func (f *FieldElement) GetLabel() Element {
	return f.children[FieldLabel]
}
func (f *FieldElement) SetLabel(l Element) *FieldElement {
	if l == nil {
		return f
	}
	f.children[FieldLabel] = l
	return f
}
func (f *FieldElement) GetError() Element {
	return f.children[FieldError]
}
func (f *FieldElement) SetError(e Element) *FieldElement {
	f.children[FieldError] = e
	return f
}
func (f *FieldElement) Type(tipe string) Element {
	f.GetInput().Type(tipe)
	return f
}

// func (f *FieldElement) Label(label string) Element {
// 	f.label = *f.label.Text(label)
// 	return f
// }

func (f *FieldElement) Id(id string) Element {
	f.GetInput().Id(id)
	label := f.GetLabel()
	if label != nil {
		label.Attr("for", id)
	}
	f.GetError().Id(id + "-error")
	return f
}
func (f *FieldElement) Name(name string) Element {
	input := f.GetInput()
	oldName := input.GetAttr("name")
	log.Printf("oldName: %v\n", oldName)
	oldId := input.GetAttr("id")
	log.Printf("oldId: %v\n", oldId)
	log.Printf("oldId == oldName: %v\n", oldId == oldName+"-field")
	input.Name(name)
	if oldId == oldName+"-field" {
		log.Printf("setting id")
		return f.Id(name + "-field")
	}
	return f
}
func (f *FieldElement) Value(value string) Element {
	f.GetInput().Value(value)
	return f
}
func (f *FieldElement) Required(required ...bool) Element {
	if len(required) > 0 && !required[0] {
		f.GetInput().RemoveAttr("required").RemoveAttr("onblur")
	} else {
		f.GetInput().Attr("required", "required").Attr("onblur", "doFieldValidation(event)")
	}
	return f
}
func (f *FieldElement) Email(is_email ...bool) Element {
	if len(is_email) > 0 && !is_email[0] {
		f.GetInput().Type("text")
	} else {
		f.GetInput().Type("email")
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
func (f *FieldElement) Error(err string) Element {
	f.GetError().Text(err)
	return f
}
func (f *FieldElement) Target(value string) Element {
	f.GetInput().Target(value)
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
func (f *FieldElement) Attr(key string, value string) Element {
	f.GetInput().Attr(key, value)
	return f
}

func (f *FieldElement) Post(value string) Element {
	f.GetInput().Post(value)
	return f
}
func (f *FieldElement) Get(value string) Element {
	f.GetInput().Get(value)
	return f
}

func (f *FieldElement) Put(value string) Element {
	f.GetInput().Put(value)
	return f
}

func (f *FieldElement) Patch(value string) Element {
	f.GetInput().Patch(value)
	return f
}

func (f *FieldElement) Delete(value string) Element {
	f.GetInput().Delete(value)
	return f
}
func (f *FieldElement) Swap(value string) Element {
	f.GetInput().Swap(value)
	return f
}

// func (f *FieldElement) Textarea() Element {
// 	f.input = *f.input.Tag("textarea")
// 	return f
// }
