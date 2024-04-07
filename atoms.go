package bolt

func Div(children ...*Element) *Element {
	return NewElement("div").Children(children...)
}
func Img(src string) *Element {
	return NewElement("img").Src(src)
}
func Button(text string) *Element {
	return NewElement("button").Text(text)
}
func Label(text string) *Element {
	return NewElement("label").Text(text)
}
func Input(name string) *Element {
	return NewElement("input").Name(name)
}
func Form() *Element {
	return NewElement("form")
}
func Span(text string) *Element {
	return NewElement("span").Text(text)
}
func P(text string) *Element {
	return NewElement("p").Text(text)
}

func H1(text string) *Element {
	return NewElement("h1").Text(text)
}
func H2(text string) *Element {
	return NewElement("h2").Text(text)
}
func H3(text string) *Element {
	return NewElement("h3").Text(text)
}
func H4(text string) *Element {
	return NewElement("h4").Text(text)
}
func H5(text string) *Element {
	return NewElement("h5").Text(text)
}
func H6(text string) *Element {
	return NewElement("h6").Text(text)
}
func A(href string, children ...*Element) *Element {
	return NewElement("a").Children(children...).Href(href)
}
func None() *Element {
	return NewElement("")
}
func Fragment(children ...*Element) *Element {
	return NewElement("").Children(children...)
}
func Html(content string) *Element {
	return NewElement("").Text(content)
}
func Section(children ...*Element) *Element {
	return NewElement("section").Children(children...)
}
