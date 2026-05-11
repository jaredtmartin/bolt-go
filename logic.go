package bolt

func If(condition bool, children ...Element) Element {
	if condition {
		return Fragment(children...)
	}
	return None()
}
func IfElse(condition bool, ifElement Element, elseElement ...Element) Element {
	if condition {
		return ifElement
	}
	if len(elseElement) > 0 {
		return Fragment(elseElement...)
	}
	return None()
}
func For[T any](objs []T, element func(T, int) Element) Element {
	result := Div("")
	for i, obj := range objs {
		result = result.Add(element(obj, i))
	}
	return result
}

// type SwitchCaseType[T comparable] func(T, T, Element) Element

// func Switch[T comparable](cases ...SwitchCaseType[T]) Element {
// 	result := Fragment()
// 	for _, caseValue := range cases {
// 		result.Add(caseValue(value))
// 	}
// 	return result
// }
// func Case[T comparable](value, condition T, element func(T) Element) Element {
// 	if value == condition {
// 		return element(value)
// 	}
// 	return None()
// }

type switchCase[T comparable] struct {
	match   T
	element Element
}

type defaultCase struct {
	element Element
}

func Case[T comparable](match T, element Element) any {
	return switchCase[T]{
		match:   match,
		element: element,
	}
}

func Default(element Element) any {
	return defaultCase{
		element: element,
	}
}

func Switch[T comparable](value T, items ...any) Element {
	var fallback Element = None()

	for _, item := range items {
		switch v := item.(type) {
		case switchCase[T]:
			if value == v.match {
				return v.element
			}
		case defaultCase:
			fallback = v.element
		}
	}

	return fallback
}
