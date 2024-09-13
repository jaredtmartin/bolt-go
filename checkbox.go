package bolt

const (
	CheckBoxLabelWrapsInput  = "CheckBoxLabelWrapsInput"
	CheckBoxLabelBeforeInput = "CheckBoxLabelBeforeInput"
	CheckBoxLabelAfterInput  = "CheckBoxLabelAfterInput"
)

type CheckboxMolecule struct {
	_id     string
	label   Element
	input   Element
	error   Element
	check   Element
	wrapper Element
	layout  string
}

func Checkbox(name string, label string) *CheckboxMolecule {
	cb := &CheckboxMolecule{
		_id:     "",
		label:   *Label(label),
		input:   *Input(name),
		error:   *Div(),
		check:   *NewElement("span"),
		wrapper: *Div(),
		layout:  CheckBoxLabelAfterInput,
	}
	return cb.Type("checkbox").setId(name + "-field")
}
func (f *CheckboxMolecule) Type(tipe string) *CheckboxMolecule {
	f.input = *f.input.Type(tipe)
	return f
}
func (f *CheckboxMolecule) Label(label string) *CheckboxMolecule {
	f.label = *f.label.Text(label)
	return f
}

func (f *CheckboxMolecule) setId(id string) *CheckboxMolecule {
	f.input = *f.input.Id(id)
	f.label = *f.label.Attr("for", id)
	f.error = *f.error.Id(id + "-error")
	return f
}
func (f *CheckboxMolecule) Name(name string) *CheckboxMolecule {
	f.input = *f.input.Name(name)
	if f._id == "" {
		return f.setId(name + "-field")
	}
	return f
}
func (f *CheckboxMolecule) Id(id string) *CheckboxMolecule {
	f._id = id
	return f.setId(id)
}
func (f *CheckboxMolecule) Value(value string) *CheckboxMolecule {
	f.input = *f.input.Value(value)
	return f
}
func (f *CheckboxMolecule) Checked(value ...bool) *CheckboxMolecule {
	if len(value) > 0 && !value[0] {
		f.input = *f.input.RemoveAttributes("checked")
	} else {
		f.input = *f.input.Attr("checked", "checked")
	}
	return f
}
func (f *CheckboxMolecule) Element() *Element {
	switch f.layout {
	case CheckBoxLabelWrapsInput:
		f.label.Children(&f.input)
		f.wrapper.Children(&f.label, &f.error)
	case CheckBoxLabelBeforeInput:
		f.wrapper.Children(&f.label, &f.input, &f.error)
	default:
		f.wrapper.Children(&f.input, &f.label, &f.error)
	}
	return &f.wrapper
}
func (f *CheckboxMolecule) Render() string {
	return f.Element().Render()
}
func (f *CheckboxMolecule) LabelClass(class string) *CheckboxMolecule {
	f.label = *f.label.Class(class)
	return f
}
func (f *CheckboxMolecule) InputClass(class string) *CheckboxMolecule {
	f.input = *f.input.Class(class)
	return f
}
func (f *CheckboxMolecule) ErrorClass(class string) *CheckboxMolecule {
	f.error = *f.error.Class(class)
	return f
}
func (f *CheckboxMolecule) WrapperClass(class string) *CheckboxMolecule {
	f.wrapper = *f.wrapper.Class(class)
	return f
}
func (f *CheckboxMolecule) LabelStyle(style string) *CheckboxMolecule {
	f.label = *f.label.Style(style)
	return f
}
func (f *CheckboxMolecule) InputStyle(style string) *CheckboxMolecule {
	f.input = *f.input.Style(style)
	return f
}
func (f *CheckboxMolecule) ErrorStyle(style string) *CheckboxMolecule {
	f.error = *f.error.Style(style)
	return f
}
func (f *CheckboxMolecule) WrapperStyle(style string) *CheckboxMolecule {
	f.wrapper = *f.wrapper.Style(style)
	return f
}
func (f *CheckboxMolecule) Error(err string) *CheckboxMolecule {
	f.error = *f.error.Text(err)
	return f
}
func (f *CheckboxMolecule) Target(value string) *CheckboxMolecule {
	f.input = *f.input.Target(value)
	return f
}
func (f *CheckboxMolecule) Attr(key string, value string) *CheckboxMolecule {
	f.input = *f.input.Attr(key, value)
	return f
}
func (f *CheckboxMolecule) Post(value string) *CheckboxMolecule {
	f.input = *f.input.Post(value)
	return f
}
func (f *CheckboxMolecule) Get(value string) *CheckboxMolecule {
	f.input = *f.input.Get(value)
	return f
}
func (f *CheckboxMolecule) Put(value string) *CheckboxMolecule {
	f.input = *f.input.Put(value)
	return f
}
func (f *CheckboxMolecule) Patch(value string) *CheckboxMolecule {
	f.input = *f.input.Patch(value)
	return f
}
func (f *CheckboxMolecule) Delete(value string) *CheckboxMolecule {
	f.input = *f.input.Delete(value)
	return f
}
func (f *CheckboxMolecule) Swap(value string) *CheckboxMolecule {
	f.input = *f.input.Swap(value)
	return f
}
func (f *CheckboxMolecule) LabelBeforeInput() *CheckboxMolecule {
	f.layout = CheckBoxLabelBeforeInput
	return f
}
func (f *CheckboxMolecule) LabelWrapsInput() *CheckboxMolecule {
	f.layout = CheckBoxLabelWrapsInput
	return f
}
func (f *CheckboxMolecule) LabelAfterInput() *CheckboxMolecule {
	f.layout = CheckBoxLabelAfterInput
	return f
}
