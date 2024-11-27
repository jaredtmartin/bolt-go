package bolt

import "fmt"

func defaultRow[T any](item T) Element {
	return P(fmt.Sprintf("%v", item))
}
func List[T any](data []T, rowFunc ...func(T) Element) Element {
	var row func(T) Element
	if len(rowFunc) == 0 {
		row = defaultRow[T]
	} else {
		row = rowFunc[0]
	}
	list := Div()
	for i := 0; i < len(data); i++ {
		list.AddChild(row(data[i]))
	}
	return list
}
