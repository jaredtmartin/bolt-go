package bolt

import "strconv"

type FieldMolecule struct {
	_id     string
	label   Element
	input   Element
	error   Element
	wrapper Element
}

func Field(name string, label string) *FieldMolecule {
	tf := FieldMolecule{
		_id:     "",
		label:   *Label(label),
		input:   *Input(name),
		error:   *Div(),
		wrapper: *Div(),
	}
	return tf.Type("text").setId(name + "-field")
}
func (f *FieldMolecule) Type(tipe string) *FieldMolecule {
	f.input = *f.input.Type(tipe)
	return f
}
func (f *FieldMolecule) Label(label string) *FieldMolecule {
	f.label = *f.label.Text(label)
	return f
}

func (f *FieldMolecule) setId(id string) *FieldMolecule {
	f.input = *f.input.Id(id)
	f.label = *f.label.Attr("for", id)
	f.error = *f.error.Id(id + "-error")
	return f
}
func (f *FieldMolecule) Name(name string) *FieldMolecule {
	f.input = *f.input.Name(name)
	if f._id == "" {
		return f.setId(name + "-field")
	}
	return f
}
func (f *FieldMolecule) Id(id string) *FieldMolecule {
	f._id = id
	return f.setId(id)
}
func (f *FieldMolecule) Value(value string) *FieldMolecule {
	f.input = *f.input.Value(value)
	return f
}
func (f *FieldMolecule) Element() *Element {
	f.wrapper.Children(&f.label, &f.input, &f.error)
	return &f.wrapper
}
func (f *FieldMolecule) Render() string {
	return f.Element().Render()
}
func (f *FieldMolecule) LabelClass(class string) *FieldMolecule {
	f.label = *f.label.Class(class)
	return f
}
func (f *FieldMolecule) InputClass(class string) *FieldMolecule {
	f.input = *f.input.Class(class)
	return f
}
func (f *FieldMolecule) ErrorClass(class string) *FieldMolecule {
	f.error = *f.error.Class(class)
	return f
}
func (f *FieldMolecule) WrapperClass(class string) *FieldMolecule {
	f.wrapper = *f.wrapper.Class(class)
	return f
}
func (f *FieldMolecule) LabelStyle(style string) *FieldMolecule {
	f.label = *f.label.Style(style)
	return f
}
func (f *FieldMolecule) InputStyle(style string) *FieldMolecule {
	f.input = *f.input.Style(style)
	return f
}
func (f *FieldMolecule) ErrorStyle(style string) *FieldMolecule {
	f.error = *f.error.Style(style)
	return f
}
func (f *FieldMolecule) WrapperStyle(style string) *FieldMolecule {
	f.wrapper = *f.wrapper.Style(style)
	return f
}
func (f *FieldMolecule) Required(required ...bool) *FieldMolecule {
	if len(required) > 0 && !required[0] {
		f.input = *f.input.RemoveAttributes("required").RemoveAttributes("onblur")
	} else {
		f.input = *f.input.Attr("required", "required").Attr("onblur", "doFieldValidation(event)")
	}
	return f
}

// func (f *TextFieldMolecule) Currency(is_currency ...bool) *TextFieldMolecule {
// 	if len(is_currency) > 0 && !is_currency[0] {
// 		f.input = *f.input.RemoveAttributes("currency").RemoveAttributes("onblur")
// 	} else {
// 		f.input = *f.input.Attr("currency", "true").Attr("onblur", "doFieldValidation(event)")
// 	}
// 	return f
// }

func (f *FieldMolecule) Email(is_email ...bool) *FieldMolecule {
	if len(is_email) > 0 && !is_email[0] {
		f.input = *f.input.Type("text")
	} else {
		f.input = *f.input.Type("email")
	}
	return f
}
func (f *FieldMolecule) Pattern(pattern string) *FieldMolecule {
	f.input = *f.input.Attr("pattern", pattern)
	return f
}
func (f *FieldMolecule) Accept(accept string) *FieldMolecule {
	f.input = *f.input.Attr("accept", accept)
	return f
}
func (f *FieldMolecule) Autocapitalize(autocapitalize string) *FieldMolecule {
	f.input = *f.input.Attr("autocapitalize", autocapitalize)
	return f
}
func (f *FieldMolecule) Autocomplete(autocomplete string) *FieldMolecule {
	f.input = *f.input.Attr("autocomplete", autocomplete)
	return f
}
func (f *FieldMolecule) Max(max string) *FieldMolecule {
	f.input = *f.input.Attr("max", max)
	return f
}
func (f *FieldMolecule) Min(min string) *FieldMolecule {
	f.input = *f.input.Attr("min", min)
	return f
}
func (f *FieldMolecule) MaxLength(maxlength string) *FieldMolecule {
	f.input = *f.input.Attr("maxlength", maxlength)
	return f
}
func (f *FieldMolecule) MinLength(minlength string) *FieldMolecule {
	f.input = *f.input.Attr("minlength", minlength)
	return f
}
func (f *FieldMolecule) Color(is_color ...bool) *FieldMolecule {
	if len(is_color) > 0 && !is_color[0] {
		f.input = *f.input.Type("text")
	} else {
		f.input = *f.input.Type("color")
	}
	return f
}
func (f *FieldMolecule) Date(is_date ...bool) *FieldMolecule {
	if len(is_date) > 0 && !is_date[0] {
		f.input = *f.input.Type("text")
	} else {
		f.input = *f.input.Type("date")
	}
	return f
}
func (f *FieldMolecule) Datetime(is_datetime ...bool) *FieldMolecule {
	if len(is_datetime) > 0 && !is_datetime[0] {
		f.input = *f.input.Type("text")
	} else {
		f.input = *f.input.Type("datetime-local") // datetime is deprecated
	}
	return f
}
func (f *FieldMolecule) File(is_file ...bool) *FieldMolecule {
	if len(is_file) > 0 && !is_file[0] {
		f.input = *f.input.Type("text")
	} else {
		f.input = *f.input.Type("file")
	}
	return f
}
func (f *FieldMolecule) Hidden(is_hidden ...bool) *FieldMolecule {
	if len(is_hidden) > 0 && !is_hidden[0] {
		f.input = *f.input.Type("text")
	} else {
		f.input = *f.input.Type("hidden")
	}
	return f
}
func (f *FieldMolecule) Month(is_month ...bool) *FieldMolecule {
	if len(is_month) > 0 && !is_month[0] {
		f.input = *f.input.Type("text")
	} else {
		f.input = *f.input.Type("month")
	}
	return f
}
func (f *FieldMolecule) Number(is_number ...bool) *FieldMolecule {
	if len(is_number) > 0 && !is_number[0] {
		f.input = *f.input.Type("text")
	} else {
		f.input = *f.input.Type("number")
	}
	return f
}
func (f *FieldMolecule) Tel(is_tel ...bool) *FieldMolecule {
	if len(is_tel) > 0 && !is_tel[0] {
		f.input = *f.input.Type("text")
	} else {
		f.input = *f.input.Type("tel")
	}
	return f
}
func (f *FieldMolecule) Time(is_time ...bool) *FieldMolecule {
	if len(is_time) > 0 && !is_time[0] {
		f.input = *f.input.Type("text")
	} else {
		f.input = *f.input.Type("time")
	}
	return f
}
func (f *FieldMolecule) Url(is_url ...bool) *FieldMolecule {
	if len(is_url) > 0 && !is_url[0] {
		f.input = *f.input.Type("text")
	} else {
		f.input = *f.input.Type("url")
	}
	return f
}
func (f *FieldMolecule) Week(is_week ...bool) *FieldMolecule {
	if len(is_week) > 0 && !is_week[0] {
		f.input = *f.input.Type("text")
	} else {
		f.input = *f.input.Type("week")
	}
	return f
}
func (f *FieldMolecule) Search(is_search ...bool) *FieldMolecule {
	if len(is_search) > 0 && !is_search[0] {
		f.input = *f.input.Type("text")
	} else {
		f.input = *f.input.Type("search")
	}
	return f
}
func (f *FieldMolecule) Password(is_password ...bool) *FieldMolecule {
	if len(is_password) > 0 && !is_password[0] {
		f.input = *f.input.Type("text")
	} else {
		f.input = *f.input.Type("password")
	}
	return f
}
func (f *FieldMolecule) Range(min int, max int) *FieldMolecule {
	f.input = *f.input.Type("range").Attr("min", strconv.Itoa(min)).Attr("max", strconv.Itoa(max))
	return f
}
func (f *FieldMolecule) Placeholder(placeholder string) *FieldMolecule {
	f.input = *f.input.Attr("placeholder", placeholder)
	return f
}
func (f *FieldMolecule) Step(step string) *FieldMolecule {
	f.input = *f.input.Attr("step", step)
	return f
}
func (f *FieldMolecule) Error(err string) *FieldMolecule {
	f.error = *f.error.Text(err)
	return f
}
func (f *FieldMolecule) Target(value string) *FieldMolecule {
	f.input = *f.input.Target(value)
	return f
}
func (f *FieldMolecule) Indicator(selector string) *FieldMolecule {
	f.input = *f.input.Indicator(selector)
	return f
}
func (f *FieldMolecule) Class(classes string) *FieldMolecule {
	f.wrapper = *f.wrapper.Class(classes)
	return f
}
func (f *FieldMolecule) Trigger(value string) *FieldMolecule {
	f.input = *f.input.Trigger(value)
	return f
}
func (f *FieldMolecule) PushUrl(pushUrl ...bool) *FieldMolecule {
	f.input = *f.input.PushUrl(pushUrl...)
	return f
}
func (f *FieldMolecule) Attr(key string, value string) *FieldMolecule {
	f.input = *f.input.Attr(key, value)
	return f
}
func (f *FieldMolecule) Post(value string) *FieldMolecule {
	f.input = *f.input.Post(value)
	return f
}
func (f *FieldMolecule) Get(value string) *FieldMolecule {
	f.input = *f.input.Get(value)
	return f
}
func (f *FieldMolecule) Put(value string) *FieldMolecule {
	f.input = *f.input.Put(value)
	return f
}
func (f *FieldMolecule) Patch(value string) *FieldMolecule {
	f.input = *f.input.Patch(value)
	return f
}
func (f *FieldMolecule) Delete(value string) *FieldMolecule {
	f.input = *f.input.Delete(value)
	return f
}
func (f *FieldMolecule) Swap(value string) *FieldMolecule {
	f.input = *f.input.Swap(value)
	return f
}
func (f *FieldMolecule) Textarea() *FieldMolecule {
	f.input = *f.input.Tag("textarea")
	return f
}
