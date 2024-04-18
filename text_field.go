package bolt

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
func (f *TextFieldMolecule) Currency(is_currency ...bool) *TextFieldMolecule {
	if len(is_currency) > 0 && !is_currency[0] {
		f.input = *f.input.RemoveAttributes("currency").RemoveAttributes("onblur")
	} else {
		f.input = *f.input.Attr("currency", "true").Attr("onblur", "doFieldValidation(event)")
	}
	return f
}
func (f *TextFieldMolecule) Email(is_email ...bool) *TextFieldMolecule {
	if len(is_email) > 0 && !is_email[0] {
		f.input = *f.input.RemoveAttributes("email").RemoveAttributes("onblur")
	} else {
		f.input = *f.input.Attr("email", "true").Attr("onblur", "doFieldValidation(event)")
	}
	return f
}
func (f *TextFieldMolecule) Password() *TextFieldMolecule {
	f.input = *f.input.Type("password")
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
