package bolt

import (
	"fmt"
	"strconv"
)

// Returns a new div element with the given class.
// If children are provided, they will be added as the content of the div.
func Div(class string, children ...Element) Element {
	e := NewElement("div").Children(children...).Class(class)
	return e
}

// Returns a new img element with the given src attribute.
func Img(src string) Element {
	return NewElement("img").Src(src)
}

// Returns a new button element with the given text.
func Button(text string) Element {
	return NewElement("button").Type("button").Text(text)
}

// Returns a new label element with the given text.
func Label(text string) Element {
	return NewElement("label").Text(text)
}

// Returns a new input element with the given name attribute.
func Input(name string) Element {
	return NewElement("input").Type("text").Name(name)
}

// Returns a new form element
func Form() Element {
	return NewElement("form")
}

// Returns a new span element
func Span(text string) Element {
	return NewElement("span").Text(text)
}

// Returns a new p element
func P(text string) Element {
	return NewElement("p").Text(text)
}

// Returns a new h1 element with the given text
func H1(text string) Element {
	return NewElement("h1").Text(text)
}

// Returns a new h2 element with the given text
func H2(text string) Element {
	return NewElement("h2").Text(text)
}

// Returns a new h3 element with the given text
func H3(text string) Element {
	return NewElement("h3").Text(text)
}

// Returns a new h4 element with the given text
func H4(text string) Element {
	return NewElement("h4").Text(text)
}

// Returns a new h5 element with the given text
func H5(text string) Element {
	return NewElement("h5").Text(text)
}

// Returns a new h6 element with the given text
func H6(text string) Element {
	return NewElement("h6").Text(text)
}

// Returns a new a element with the given href.
// If children are provided, they will be used as the content of the link.
func A(href string, children ...Element) Element {
	return NewElement("a").Children(children...).Href(href)
}

// Returns a new element with no tag
func None() Element {
	return NewElement("")
}

// Returns a new element with no tag
// If children are provided, they will be added as the content of the element.
func Fragment(children ...Element) Element {
	return NewElement("").Children(children...)
}

// Returns a new html element with the given UNESCAPED content
func Html(content string) Element {
	return NewElement("").UnsafeHtml(content)
}

// Returns a new section element
// If children are provided, they will be added as the content of the section.
func Section(children ...Element) Element {
	return NewElement("section").Children(children...)
}

// Returns a new hidden input element with the given name and value.
func HiddenInput(name string, value string) Element {
	return NewElement("input").Value(value).Type("hidden").Name(name)
}

// Returns a new element based on a template file.
// func Template(filename string) Element {
// 	html, err := os.ReadFile(filename + ".html")
// 	if err != nil {
// 		html, err = os.ReadFile(filename + "tmpl")
// 	}
// 	if err != nil {
// 		html, err = os.ReadFile(filename)
// 	}
// 	if err != nil {
// 		panic(err)
// 	}
// 	return NewElement("").UnsafeHtml(string(html))
// }

// returns a script tag
// If js is provided, it will be used as the UNESCAPED content of the script tag.
func Script(js string) Element {
	return NewElement("script").UnsafeHtml(js)
}

// returns a stylesheet link
// The url parameter specifies the URL of the stylesheet.
// The link tag is set to have a rel attribute of "stylesheet".
// It is used to link external CSS files to the HTML document.
func Stylesheet(url string) Element {
	return NewElement("link").Href(url).Tag("link").Attr("rel", "stylesheet")
}

// returns a style tag with the given css code WITHOUT ESCAPING
func Style(css string) Element {
	return NewElement("style").UnsafeHtml(css)
}

// returns an svg tag
// The path parameter should be the SVG content as a string.
// The width and height parameters specify the dimensions of the SVG.
// The vbX and vbY parameters specify the viewBox dimensions.
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

// returns a list item (li) tag
// If children are provided, they will be added as the content of the list item.
func Li(children ...Element) Element {
	return NewElement("li").Children(children...)
}

// returns a unordered list (ul) tag
// If children are provided, they will be added as the content of the unordered list.
func Ul(children ...Element) Element {
	return NewElement("ul").Children(children...)
}

// returns an iframe tag
// The title parameter sets the title attribute of the iframe.
// The width and height parameters specify the dimensions of the iframe.
// The src parameter specifies the source URL of the iframe.
// The iframe is set to allow autoplay, fullscreen, picture-in-picture, and clipboard-write.
// It also has a class of "w-full h-full" for styling.
// The data-ready attribute is set to "true" to indicate that the iframe is ready.
// The frameborder attribute is set to "0" to remove the border around the iframe.
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

// returns a head element
// If children are provided, they will be added as the content.
func Head(children ...Element) Element {
	return NewElement("head").Children(children...)
}

// returns a body element
// If children are provided, they will be added as the content.
func Body(children ...Element) Element {
	return NewElement("body").Children(children...)
}

// returns an Alpine template element
// If children are provided, they will be added as the content of the template.
func Template(children ...Element) Element {
	return NewElement("template").Children(children...)
}

// returns a table element
// If children are provided, they will be added as the content of the table.
func Table(children ...Element) Element {
	return NewElement("table").Children(children...)
}

// returns a table row element
// If children are provided, they will be added as the content of the row.
func Tr(children ...Element) Element {
	return NewElement("tr").Children(children...)
}

// returns a table header row element
// If children are provided, they will be added as the content of the header row.
func Th(children ...Element) Element {
	return NewElement("th").Children(children...)
}

// returns a table cell element
// If children are provided, they will be added as the content of the cell.
func Td(children ...Element) Element {
	return NewElement("td").Children(children...)
}

// returns a simple string content as an element
// This is useful for creating text nodes without any HTML tags.
func String(text string) Element {
	return NewElement("").Text(text)
}

// Oob returns a div element with the given value as an HTMX out-of-band (OOB) swap value.
func Oob(value string, children ...Element) Element {
	return Div("", children...).HXSwapOob(value)
}
