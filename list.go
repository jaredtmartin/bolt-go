package bolt

import "fmt"

// Default function to render each item in a list
func defaultRow[T any](item T, idx int) Element {
	return P(fmt.Sprintf("%v", item))
}

// Returns a list of elements from the given data wrapped in a div using an optional function to render each item.
func List[T any](data []T, rowFunc ...func(T, int) Element) Element {
	var row func(T, int) Element
	if len(rowFunc) == 0 {
		row = defaultRow[T]
	} else {
		row = rowFunc[0]
	}
	list := Div("")
	for i := range data {
		list.AddChild(row(data[i], i))
	}
	return list
}
