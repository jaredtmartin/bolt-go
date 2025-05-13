package bolt

import (
	"fmt"
	"os"
	"strconv"
)

func Div(class string, children ...Element) Element {
	e := NewElement("div").Children(children...).Class(class)
	return e
}
func Img(src string) Element {
	return NewElement("img").Src(src)
}
func Button(text string) Element {
	return NewElement("button").Type("button").Text(text)
}
func Label(text string) Element {
	return NewElement("label").Text(text)
}
func Input(name string) Element {
	return NewElement("input").Type("text").Name(name)
}
func Form() Element {
	return NewElement("form")
}
func Span(text string) Element {
	return NewElement("span").Text(text)
}
func P(text string) Element {
	return NewElement("p").Text(text)
}
func H1(text string) Element {
	return NewElement("h1").Text(text)
}
func H2(text string) Element {
	return NewElement("h2").Text(text)
}
func H3(text string) Element {
	return NewElement("h3").Text(text)
}
func H4(text string) Element {
	return NewElement("h4").Text(text)
}
func H5(text string) Element {
	return NewElement("h5").Text(text)
}
func H6(text string) Element {
	return NewElement("h6").Text(text)
}
func A(href string, children ...Element) Element {
	return NewElement("a").Children(children...).Href(href)
}
func None() Element {
	return NewElement("")
}
func Fragment(children ...Element) Element {
	return NewElement("").Children(children...)
}
func Html(content string) Element {
	return NewElement("").UnsafeHtml(content)
}
func Section(children ...Element) Element {
	return NewElement("section").Children(children...)
}
func Template(filename string) Element {
	html, err := os.ReadFile(filename + ".html")
	if err != nil {
		html, err = os.ReadFile(filename + "tmpl")
	}
	if err != nil {
		html, err = os.ReadFile(filename)
	}
	if err != nil {
		panic(err)
	}
	return NewElement("").UnsafeHtml(string(html))
}

// returns a script tag with the given js code without escaping
func Script(js string) Element {
	return NewElement("script").UnsafeHtml(js) //.Attr("defer", "")
}
func Stylesheet(url string) Element {
	return NewElement("link").Href(url).Tag("link").Attr("rel", "stylesheet")
}
func Svg(path string, width int, height int, vbX int, vbY int) Element {
	viewBox := fmt.Sprintf("0 0 %d %d", vbX, vbY)
	return NewElement("svg").
		Attr("viewBox", viewBox).
		Attr("version", "1.1").
		Attr("xmlns", "http://www.w3.org/2000/svg").
		Attr("xmlns:xlink", "http://www.w3.org/1999/xlink").
		Attr("version", "1.1").
		Attr("width", strconv.Itoa(width)).
		Attr("height", strconv.Itoa(height)).
		Class("fill-current").
		UnsafeHtml(path)
}
func VideoIframe(title string, width int, height int, src string) Element {
	return NewElement("iframe").
		Attr("width", strconv.Itoa(width)).
		Attr("height", strconv.Itoa(height)).
		Attr("frameborder", "0").
		Attr("allow", "autoplay; fullscreen; picture-in-picture; clipboard-write").
		Attr("title", title).
		Attr("data-ready", "true").
		Class("w-full h-full").
		Src(src)
}
func Head(children ...Element) Element {
	return NewElement("head").Children(children...)
}
func Body(children ...Element) Element {
	return NewElement("body").Children(children...)
}
