package bolt

type Option struct {
	Label string
	Value string
}
type SelectMolecule struct {
	_id     string
	value   string
	label   Element
	input   Element
	error   Element
	wrapper Element
	options []Option
}

func Select(name string, label string, options []Option) *SelectMolecule {
	molecule := SelectMolecule{
		_id:     "",
		label:   *Label(label),
		input:   *NewElement("select").Name(name),
		error:   *Div(),
		wrapper: *Div(),
		options: options,
	}
	return molecule.setId(name + "-field")
}

//	func Selector(name string, value string, options []Option) *Element {
//		var opts []*Element
//		for _, o := range options {
//			newopt := NewElement("option").Value(o.Value).Text(o.Label)
//			if o.Value == value {
//				newopt.Attr("selected", "")
//			}
//			opts = append(opts, newopt)
//		}
//		return NewElement("select").Children(opts...)
//	}
func (f *SelectMolecule) Label(label string) *SelectMolecule {
	f.label = *f.label.Text(label)
	return f
}

func (f *SelectMolecule) setId(id string) *SelectMolecule {
	f.input = *f.input.Id(id)
	f.label = *f.label.Attr("for", id)
	f.error = *f.error.Id(id + "-error")
	return f
}
func (f *SelectMolecule) Name(name string) *SelectMolecule {
	f.input = *f.input.Name(name)
	if f._id == "" {
		return f.setId(name + "-field")
	}
	return f
}
func (f *SelectMolecule) Id(id string) *SelectMolecule {
	f._id = id
	return f.setId(id)
}
func (f *SelectMolecule) Value(value string) *SelectMolecule {
	f.value = value
	return f
}
func (f *SelectMolecule) Element() *Element {
	var opts []*Element
	for _, o := range f.options {
		newopt := NewElement("option").Value(o.Value).Text(o.Label)
		if o.Value == f.value {
			newopt.Attr("selected", "")
		}
		opts = append(opts, newopt)
	}
	input := f.input.Children(opts...)
	f.wrapper.Children(&f.label, input, &f.error)
	return &f.wrapper
}
func (f *SelectMolecule) Render() string {
	return f.Element().Render()
}
func (f *SelectMolecule) LabelClass(class string) *SelectMolecule {
	f.label = *f.label.Class(class)
	return f
}
func (f *SelectMolecule) InputClass(class string) *SelectMolecule {
	f.input = *f.input.Class(class)
	return f
}
func (f *SelectMolecule) ErrorClass(class string) *SelectMolecule {
	f.error = *f.error.Class(class)
	return f
}
func (f *SelectMolecule) WrapperClass(class string) *SelectMolecule {
	f.wrapper = *f.wrapper.Class(class)
	return f
}
func (f *SelectMolecule) LabelStyle(style string) *SelectMolecule {
	f.label = *f.label.Style(style)
	return f
}
func (f *SelectMolecule) InputStyle(style string) *SelectMolecule {
	f.input = *f.input.Style(style)
	return f
}
func (f *SelectMolecule) ErrorStyle(style string) *SelectMolecule {
	f.error = *f.error.Style(style)
	return f
}
func (f *SelectMolecule) WrapperStyle(style string) *SelectMolecule {
	f.wrapper = *f.wrapper.Style(style)
	return f
}
func (f *SelectMolecule) Required(required ...bool) *SelectMolecule {
	if len(required) > 0 && !required[0] {
		f.input = *f.input.RemoveAttributes("required").RemoveAttributes("onblur")
	} else {
		f.input = *f.input.Attr("required", "required").Attr("onblur", "doFieldValidation(event)")
	}
	return f
}
