package bolt

import "strconv"

type TextFieldMolecule struct {
	_id     string
	label   Element
	input   Element
	error   Element
	wrapper Element
}

func TextField(name string, label string) *TextFieldMolecule {
	tf := TextFieldMolecule{
		_id:     "",
		label:   *Label(label),
		input:   *Input(name),
		error:   *Div(),
		wrapper: *Div(),
	}
	return tf.Type("text").setId(name + "-field")
}
func (f *TextFieldMolecule) Type(tipe string) *TextFieldMolecule {
	f.input = *f.input.Type(tipe)
	return f
}
func (f *TextFieldMolecule) Label(label string) *TextFieldMolecule {
	f.label = *f.label.Text(label)
	return f
}

func (f *TextFieldMolecule) setId(id string) *TextFieldMolecule {
	f.input = *f.input.Id(id)
	f.label = *f.label.Attr("for", id)
	f.error = *f.error.Id(id + "-error")
	return f
}
func (f *TextFieldMolecule) Name(name string) *TextFieldMolecule {
	f.input = *f.input.Name(name)
	if f._id == "" {
		return f.setId(name + "-field")
	}
	return f
}
func (f *TextFieldMolecule) Id(id string) *TextFieldMolecule {
	f._id = id
	return f.setId(id)
}
func (f *TextFieldMolecule) Value(value string) *TextFieldMolecule {
	f.input = *f.input.Value(value)
	return f
}
func (f *TextFieldMolecule) Element() *Element {
	f.wrapper.Children(&f.label, &f.input, &f.error)
	return &f.wrapper
}
func (f *TextFieldMolecule) Render() string {
	return f.Element().Render()
}
func (f *TextFieldMolecule) LabelClass(class string) *TextFieldMolecule {
	f.label = *f.label.Class(class)
	return f
}
func (f *TextFieldMolecule) InputClass(class string) *TextFieldMolecule {
	f.input = *f.input.Class(class)
	return f
}
func (f *TextFieldMolecule) ErrorClass(class string) *TextFieldMolecule {
	f.error = *f.error.Class(class)
	return f
}
func (f *TextFieldMolecule) WrapperClass(class string) *TextFieldMolecule {
	f.wrapper = *f.wrapper.Class(class)
	return f
}
func (f *TextFieldMolecule) LabelStyle(style string) *TextFieldMolecule {
	f.label = *f.label.Style(style)
	return f
}
func (f *TextFieldMolecule) InputStyle(style string) *TextFieldMolecule {
	f.input = *f.input.Style(style)
	return f
}
func (f *TextFieldMolecule) ErrorStyle(style string) *TextFieldMolecule {
	f.error = *f.error.Style(style)
	return f
}
func (f *TextFieldMolecule) WrapperStyle(style string) *TextFieldMolecule {
	f.wrapper = *f.wrapper.Style(style)
	return f
}
func (f *TextFieldMolecule) Required(required ...bool) *TextFieldMolecule {
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

func (f *TextFieldMolecule) Email(is_email ...bool) *TextFieldMolecule {
	if len(is_email) > 0 && !is_email[0] {
		f.input = *f.input.Type("text")
	} else {
		f.input = *f.input.Type("email")
	}
	return f
}
func (f *TextFieldMolecule) Pattern(pattern string) *TextFieldMolecule {
	f.input = *f.input.Attr("pattern", pattern)
	return f
}
func (f *TextFieldMolecule) Accept(accept string) *TextFieldMolecule {
	f.input = *f.input.Attr("accept", accept)
	return f
}
func (f *TextFieldMolecule) Autocapitalize(autocapitalize string) *TextFieldMolecule {
	f.input = *f.input.Attr("autocapitalize", autocapitalize)
	return f
}
func (f *TextFieldMolecule) Autocomplete(autocomplete string) *TextFieldMolecule {
	f.input = *f.input.Attr("autocomplete", autocomplete)
	return f
}
func (f *TextFieldMolecule) Max(max string) *TextFieldMolecule {
	f.input = *f.input.Attr("max", max)
	return f
}
func (f *TextFieldMolecule) Min(min string) *TextFieldMolecule {
	f.input = *f.input.Attr("min", min)
	return f
}
func (f *TextFieldMolecule) MaxLength(maxlength string) *TextFieldMolecule {
	f.input = *f.input.Attr("maxlength", maxlength)
	return f
}
func (f *TextFieldMolecule) MinLength(minlength string) *TextFieldMolecule {
	f.input = *f.input.Attr("minlength", minlength)
	return f
}
func (f *TextFieldMolecule) Color(is_color ...bool) *TextFieldMolecule {
	if len(is_color) > 0 && !is_color[0] {
		f.input = *f.input.Type("text")
	} else {
		f.input = *f.input.Type("color")
	}
	return f
}
func (f *TextFieldMolecule) Date(is_date ...bool) *TextFieldMolecule {
	if len(is_date) > 0 && !is_date[0] {
		f.input = *f.input.Type("text")
	} else {
		f.input = *f.input.Type("date")
	}
	return f
}
func (f *TextFieldMolecule) Datetime(is_datetime ...bool) *TextFieldMolecule {
	if len(is_datetime) > 0 && !is_datetime[0] {
		f.input = *f.input.Type("text")
	} else {
		f.input = *f.input.Type("datetime-local") // datetime is deprecated
	}
	return f
}
func (f *TextFieldMolecule) File(is_file ...bool) *TextFieldMolecule {
	if len(is_file) > 0 && !is_file[0] {
		f.input = *f.input.Type("text")
	} else {
		f.input = *f.input.Type("file")
	}
	return f
}
func (f *TextFieldMolecule) Hidden(is_hidden ...bool) *TextFieldMolecule {
	if len(is_hidden) > 0 && !is_hidden[0] {
		f.input = *f.input.Type("text")
	} else {
		f.input = *f.input.Type("hidden")
	}
	return f
}
func (f *TextFieldMolecule) Month(is_month ...bool) *TextFieldMolecule {
	if len(is_month) > 0 && !is_month[0] {
		f.input = *f.input.Type("text")
	} else {
		f.input = *f.input.Type("month")
	}
	return f
}
func (f *TextFieldMolecule) Number(is_number ...bool) *TextFieldMolecule {
	if len(is_number) > 0 && !is_number[0] {
		f.input = *f.input.Type("text")
	} else {
		f.input = *f.input.Type("number")
	}
	return f
}
func (f *TextFieldMolecule) Tel(is_tel ...bool) *TextFieldMolecule {
	if len(is_tel) > 0 && !is_tel[0] {
		f.input = *f.input.Type("text")
	} else {
		f.input = *f.input.Type("tel")
	}
	return f
}
func (f *TextFieldMolecule) Time(is_time ...bool) *TextFieldMolecule {
	if len(is_time) > 0 && !is_time[0] {
		f.input = *f.input.Type("text")
	} else {
		f.input = *f.input.Type("time")
	}
	return f
}
func (f *TextFieldMolecule) Url(is_url ...bool) *TextFieldMolecule {
	if len(is_url) > 0 && !is_url[0] {
		f.input = *f.input.Type("text")
	} else {
		f.input = *f.input.Type("url")
	}
	return f
}
func (f *TextFieldMolecule) Week(is_week ...bool) *TextFieldMolecule {
	if len(is_week) > 0 && !is_week[0] {
		f.input = *f.input.Type("text")
	} else {
		f.input = *f.input.Type("week")
	}
	return f
}
func (f *TextFieldMolecule) Search(is_search ...bool) *TextFieldMolecule {
	if len(is_search) > 0 && !is_search[0] {
		f.input = *f.input.Type("text")
	} else {
		f.input = *f.input.Type("search")
	}
	return f
}
func (f *TextFieldMolecule) Password(is_password ...bool) *TextFieldMolecule {
	if len(is_password) > 0 && !is_password[0] {
		f.input = *f.input.Type("text")
	} else {
		f.input = *f.input.Type("password")
	}
	return f
}
func (f *TextFieldMolecule) Range(min int, max int) *TextFieldMolecule {
	f.input = *f.input.Attr("min", strconv.Itoa(min)).Attr("max", strconv.Itoa(max))
	return f
}
func (f *TextFieldMolecule) Placeholder(placeholder string) *TextFieldMolecule {
	f.input = *f.input.Attr("placeholder", placeholder)
	return f
}
func (f *TextFieldMolecule) Step(step string) *TextFieldMolecule {
	f.input = *f.input.Attr("step", step)
	return f
}
func (f *TextFieldMolecule) Error(err string) *TextFieldMolecule {
	f.error = *f.error.Text(err)
	return f
}
func (f *TextFieldMolecule) Target(value string) *TextFieldMolecule {
	f.input = *f.input.Target(value)
	return f
}
func (f *TextFieldMolecule) Indicator(selector string) *TextFieldMolecule {
	f.input = *f.input.Indicator(selector)
	return f
}
func (f *TextFieldMolecule) Class(classes string) *TextFieldMolecule {
	f.wrapper = *f.wrapper.Class(classes)
	return f
}
func (f *TextFieldMolecule) Trigger(value string) *TextFieldMolecule {
	f.input = *f.input.Trigger(value)
	return f
}
func (f *TextFieldMolecule) PushUrl(pushUrl ...bool) *TextFieldMolecule {
	f.input = *f.input.PushUrl(pushUrl...)
	return f
}
func (f *TextFieldMolecule) Attr(key string, value string) *TextFieldMolecule {
	f.input = *f.input.Attr(key, value)
	return f
}
func (f *TextFieldMolecule) Post(value string) *TextFieldMolecule {
	f.input = *f.input.Post(value)
	return f
}
func (f *TextFieldMolecule) Get(value string) *TextFieldMolecule {
	f.input = *f.input.Get(value)
	return f
}
func (f *TextFieldMolecule) Put(value string) *TextFieldMolecule {
	f.input = *f.input.Put(value)
	return f
}
func (f *TextFieldMolecule) Patch(value string) *TextFieldMolecule {
	f.input = *f.input.Patch(value)
	return f
}
func (f *TextFieldMolecule) Delete(value string) *TextFieldMolecule {
	f.input = *f.input.Delete(value)
	return f
}
func (f *TextFieldMolecule) Swap(value string) *TextFieldMolecule {
	f.input = *f.input.Swap(value)
	return f
}
