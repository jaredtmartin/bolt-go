package bolt

import "fmt"

type ListMolecule struct {
	wrapper   Element
	renderRow func(interface{}) *Element
}

func defaultRow(item interface{}) *Element {
	return P(fmt.Sprintf("%v", item))
}
func List() *ListMolecule {
	return &ListMolecule{
		wrapper:   *Div(),
		renderRow: defaultRow,
	}
}
func (f *ListMolecule) Element(data ...interface{}) *Element {
	items := []*Element{}
	for i := 0; i < len(data); i++ {
		items = append(items, f.renderRow(data[i]))
	}
	f.wrapper.Children(items...)
	return &f.wrapper
}
func (f *ListMolecule) Render(data ...interface{}) string {
	return f.Element(data...).Render()
}
