package bolt_test

import (
	"testing"

	"github.com/jaredtmartin/bolt-go"
)

func TestDiv(t *testing.T) {
	e := bolt.Div("")
	result := e.Render()
	expected := "<div></div>"
	if result != expected {
		t.Fatalf(`Div(") = %q, expected %q`, result, expected)
	}
}

func TestSpan(t *testing.T) {
	e := bolt.Span("Hello")
	result := e.Render()
	expected := "<span>Hello</span>"
	if result != expected {
		t.Fatalf(`Span() = %q, expected %q`, result, expected)
	}
}

func TestP(t *testing.T) {
	e := bolt.P("Hello")
	result := e.Render()
	expected := "<p>Hello</p>"
	if result != expected {
		t.Fatalf(`P() = %q, expected %q`, result, expected)
	}
}

func TestH1(t *testing.T) {
	e := bolt.H1("Hello")
	result := e.Render()
	expected := "<h1>Hello</h1>"
	if result != expected {
		t.Fatalf(`H1() = %q, expected %q`, result, expected)
	}
}

func TestH2(t *testing.T) {
	e := bolt.H2("Hello")
	result := e.Render()
	expected := "<h2>Hello</h2>"
	if result != expected {
		t.Fatalf(`H2() = %q, expected %q`, result, expected)
	}
}

func TestH3(t *testing.T) {
	e := bolt.H3("Hello")
	result := e.Render()
	expected := "<h3>Hello</h3>"
	if result != expected {
		t.Fatalf(`H3() = %q, expected %q`, result, expected)
	}
}

func TestH4(t *testing.T) {
	e := bolt.H4("Hello")
	result := e.Render()
	expected := "<h4>Hello</h4>"
	if result != expected {
		t.Fatalf(`H4() = %q, expected %q`, result, expected)
	}
}

func TestH5(t *testing.T) {
	e := bolt.H5("Hello")
	result := e.Render()
	expected := "<h5>Hello</h5>"
	if result != expected {
		t.Fatalf(`H5() = %q, expected %q`, result, expected)
	}
}

func TestH6(t *testing.T) {
	e := bolt.H6("Hello")
	result := e.Render()
	expected := "<h6>Hello</h6>"
	if result != expected {
		t.Fatalf(`H6() = %q, expected %q`, result, expected)
	}
}

// Img
func TestImg(t *testing.T) {
	e := bolt.Img("src")
	result := e.Render()
	expected := "<img src=\"src\">"
	if result != expected {
		t.Fatalf(`Img() = %q, expected %q`, result, expected)
	}
}

// Button
func TestButton(t *testing.T) {
	e := bolt.Button("Hello")
	result := e.Render()
	expected := "<button type=\"button\">Hello</button>"
	if result != expected {
		t.Fatalf(`Button() = %q, expected %q`, result, expected)
	}
}

// Label
func TestLabel(t *testing.T) {
	e := bolt.Label("Hello")
	result := e.Render()
	expected := "<label>Hello</label>"
	if result != expected {
		t.Fatalf(`Label() = %q, expected %q`, result, expected)
	}
}

// Input
func TestInput(t *testing.T) {
	e := bolt.Input("first_name")
	result := e.Render()
	expected := "<input name=\"first_name\" type=\"text\">"
	if result != expected {
		t.Fatalf(`Input() = %q, expected %q`, result, expected)
	}
}

// Form
func TestForm(t *testing.T) {
	e := bolt.Form().Children(bolt.Input("name"))
	result := e.Render()
	expected := "<form><input name=\"name\" type=\"text\"></form>"
	if result != expected {
		t.Fatalf(`Form() = %q, expected %q`, result, expected)
	}
}

// A
func TestA(t *testing.T) {
	e := bolt.A("www.google.com", bolt.Html("Hello"))
	result := e.Render()
	expected := "<a href=\"www.google.com\">Hello</a>"
	if result != expected {
		t.Fatalf(`A() = %q, expected %q`, result, expected)
	}
}

// None
func TestNone(t *testing.T) {
	e := bolt.None()
	result := e.Render()
	expected := ""
	if result != expected {
		t.Fatalf(`None() = %q, expected %q`, result, expected)
	}
}

// Fragment
func TestFragment(t *testing.T) {
	e := bolt.Fragment()
	result := e.Render()
	expected := ""
	if result != expected {
		t.Fatalf(`Fragment() = %q, expected %q`, result, expected)
	}
}

// Html
func TestUnsafeHtmlElement(t *testing.T) {
	e := bolt.Html("<div>Hello</div>")
	result := e.Render()
	expected := "<div>Hello</div>"
	if result != expected {
		t.Fatalf(`Html() = %q, expected %q`, result, expected)
	}
}
func TestSection(t *testing.T) {
	e := bolt.Section(bolt.P("Hello"))
	result := e.Render()
	expected := "<section><p>Hello</p></section>"
	if result != expected {
		t.Fatalf(`Section() = %q, expected %q`, result, expected)
	}
}

//	func TestTemplate(t *testing.T) {
//		result := Template("sample").Render()
//		expected := "<p>Hello Html!</p>"
//		if result != expected {
//			t.Fatalf(`Template("sample") = %q, expected %q`, result, expected)
//		}
//		result = Template("sample.html").Render()
//		if result != expected {
//			t.Fatalf(`Template("sample.html") = %q, expected %q`, result, expected)
//		}
//		result = Template("sample.tmpl").Render()
//		expected = "<p>Hello Tmpl!</p>"
//		if result != expected {
//			t.Fatalf(`Template("sample.tmpl") = %q, expected %q`, result, expected)
//		}
//	}
func TestTemplate(t *testing.T) {
	result := bolt.Template(bolt.P("Content")).Render()
	expected := "<template><p>Content</p></template>"
	if result != expected {
		t.Fatalf(`Template(P("Content")) = %q, expected %q`, result, expected)
	}
}
func TestSvg(t *testing.T) {
	e := bolt.Svg().ViewBox(0, 0, 100, 100).Height(100).Width(100).Class("fill-current").Children(
		bolt.Path().D("M10 10 H 90 V 90 H 10 Z"),
	)
	result := e.Render()
	expected := `<svg height="100" version="1.1" viewBox="0 0 100 100" width="100" xmlns:xlink="http://www.w3.org/1999/xlink" xmlns="http://www.w3.org/2000/svg" class="fill-current"><path d="M10 10 H 90 V 90 H 10 Z"></path></svg>`

	if result != expected {
		t.Fatalf(`Svg() = %q, expected %q`, result, expected)
	}
	e = bolt.Svg().ViewBox(24, 24, 24, 24)
	e.Class("fill-current").UnsafeHtml(`<path d="M15 19l-7-7 7-7" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"></path>`)
	result = e.Render()
	expected = `<svg version="1.1" viewBox="24 24 24 24" xmlns:xlink="http://www.w3.org/1999/xlink" xmlns="http://www.w3.org/2000/svg" class="fill-current"><path d="M15 19l-7-7 7-7" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"></path></svg>`
	if result != expected {
		t.Fatalf(`Svg() = %q, expected %q`, result, expected)
	}
}
func TestVideoIframe(t *testing.T) {
	e := bolt.VideoIframe("Hello", 23, 25, "https://www.youtube.com/embed/dQw4w9WgXcQ")
	result := e.Render()
	expected := `<iframe allow="autoplay; fullscreen; picture-in-picture; clipboard-write" data-ready="true" frameborder="0" height="25" src="https://www.youtube.com/embed/dQw4w9WgXcQ" title="Hello" width="23" class="h-full w-full"></iframe>`
	if result != expected {
		t.Fatalf(`Svg() = %q, expected %q`, result, expected)
	}
}
func TestScript(t *testing.T) {
	// Test embedded script
	e := bolt.Script("console.log('test')")
	result := e.Render()
	expected := `<script>console.log('test')</script>`
	if result != expected {
		t.Fatalf(`Script("console.log('test')") = %q, expected %q`, result, expected)
	}
	// Test linked script
	e = bolt.Script("").Src("/src/test.js")
	result = e.Render()
	expected = `<script src="/src/test.js"></script>`
	if result != expected {
		t.Fatalf(`Script("/src/test.js") = %q, expected %q`, result, expected)
	}
	// Test linked script that is deferred
	e = bolt.Script("").Src("/src/test.js").Defer()
	result = e.Render()
	expected = `<script defer="true" src="/src/test.js"></script>`
	if result != expected {
		t.Fatalf(`Script("/src/test.js") = %q, expected %q`, result, expected)
	}
}
func TestHead(t *testing.T) {
	// Test empty head
	e := bolt.Head()
	result := e.Render()
	expected := "<head></head>"
	if result != expected {
		t.Fatalf(`Head() = %q, expected %q`, result, expected)
	}

	// Test head with single child
	e = bolt.Head(bolt.Script("console.log('test')"))
	result = e.Render()
	expected = "<head><script>console.log('test')</script></head>"
	if result != expected {
		t.Fatalf(`Head() with single child = %q, expected %q`, result, expected)
	}

	// Test head with multiple children
	e = bolt.Head(
		bolt.Script("console.log('test')"),
		bolt.Html("<meta charset=\"UTF-8\">"),
		bolt.Html("<title>Test Page</title>"),
	)
	result = e.Render()
	expected = "<head><script>console.log('test')</script><meta charset=\"UTF-8\"><title>Test Page</title></head>"
	if result != expected {
		t.Fatalf(`Head() with multiple children = %q, expected %q`, result, expected)
	}
}
func TestBody(t *testing.T) {
	// Test empty body
	e := bolt.Body()
	result := e.Render()
	expected := "<body></body>"
	if result != expected {
		t.Fatalf(`Body() = %q, expected %q`, result, expected)
	}

	// Test body with single child
	e = bolt.Body(bolt.P("Hello"))
	result = e.Render()
	expected = "<body><p>Hello</p></body>"
	if result != expected {
		t.Fatalf(`Body() = %q, expected %q`, result, expected)
	}

	// Test body with multiple children
	e = bolt.Body(bolt.H1("Title"), bolt.P("Paragraph"), bolt.Div("", bolt.P("Content")))
	result = e.Render()
	expected = "<body><h1>Title</h1><p>Paragraph</p><div><p>Content</p></div></body>"
	if result != expected {
		t.Fatalf(`Body() = %q, expected %q`, result, expected)
	}

	// Test body with nested elements
	e = bolt.Body(bolt.Div("", bolt.P("Nested")))
	result = e.Render()
	expected = "<body><div><p>Nested</p></div></body>"
	if result != expected {
		t.Fatalf(`Body() = %q, expected %q`, result, expected)
	}
}
func TestStylesheet(t *testing.T) {
	// Test basic stylesheet
	e := bolt.Stylesheet("/styles/main.css")
	result := e.Render()
	expected := `<link href="/styles/main.css" rel="stylesheet">`
	if result != expected {
		t.Fatalf(`Stylesheet("/styles/main.css") = %q, expected %q`, result, expected)
	}
}

func TestHiddenInput(t *testing.T) {
	e := bolt.HiddenInput("name", "value")
	result := e.Render()
	expected := `<input name="name" type="hidden" value="value">`
	if result != expected {
		t.Fatalf(`HiddenInput("name", "value") = %q, expected %q`, result, expected)
	}
}

func TestLi(t *testing.T) {
	e := bolt.Li(bolt.P("Content"))
	result := e.Render()
	expected := `<li><p>Content</p></li>`
	if result != expected {
		t.Fatalf(`Li(P("Content")) = %q, expected %q`, result, expected)
	}
}

func TestTable(t *testing.T) {
	e := bolt.Table(
		bolt.Tr(
			bolt.Th(bolt.P("Header 1")),
			bolt.Th(bolt.P("Header 2")),
		),
		bolt.Tr(
			bolt.Td(bolt.P("Row 1, Cell 1")),
			bolt.Td(bolt.P("Row 1, Cell 2")),
		),
	)
	result := e.Render()
	expected := `<table><tr><th><p>Header 1</p></th><th><p>Header 2</p></th></tr><tr><td><p>Row 1, Cell 1</p></td><td><p>Row 1, Cell 2</p></td></tr></table>`
	if result != expected {
		t.Fatalf(`Table() = %q, expected %q`, result, expected)
	}
}
func TestTr(t *testing.T) {
	e := bolt.Tr(
		bolt.Td(bolt.P("Row 1, Cell 1")),
		bolt.Td(bolt.P("Row 1, Cell 2")),
	)
	result := e.Render()
	expected := `<tr><td><p>Row 1, Cell 1</p></td><td><p>Row 1, Cell 2</p></td></tr>`
	if result != expected {
		t.Fatalf(`Tr() = %q, expected %q`, result, expected)
	}
}
func TestTh(t *testing.T) {
	e := bolt.Th(
		bolt.Td(bolt.P("Header 1")),
		bolt.Td(bolt.P("Header 2")),
	)
	result := e.Render()
	expected := `<th><td><p>Header 1</p></td><td><p>Header 2</p></td></th>`
	if result != expected {
		t.Fatalf(`Th() = %q, expected %q`, result, expected)
	}
}
func TestTd(t *testing.T) {
	e := bolt.Td(bolt.P("Content"))
	result := e.Render()
	expected := `<td><p>Content</p></td>`
	if result != expected {
		t.Fatalf(`Td(P("Content")) = %q, expected %q`, result, expected)
	}
}
func TestString(t *testing.T) {
	e := bolt.Template(bolt.P("Content"))
	result := e.Render()
	expected := `<template><p>Content</p></template>`
	if result != expected {
		t.Fatalf(`Template(P("Content")) = %q, expected %q`, result, expected)
	}
}

func TestStyle(t *testing.T) {
	// Test with simple CSS
	e := bolt.Style("body { background: #fff; }")
	result := e.Render()
	expected := `<style>body { background: #fff; }</style>`
	if result != expected {
		t.Fatalf(`Style("body { background: #fff; }") = %q, expected %q`, result, expected)
	}

	// Test with empty CSS
	e = bolt.Style("")
	result = e.Render()
	expected = `<style></style>`
	if result != expected {
		t.Fatalf(`Style("") = %q, expected %q`, result, expected)
	}

	// Test with multiline CSS
	css := `
body {
	color: #333;
}
h1 {
	font-size: 2em;
}`
	e = bolt.Style(css)
	result = e.Render()
	expected = `<style>` + css + `</style>`
	if result != expected {
		t.Fatalf(`Style(multiline) = %q, expected %q`, result, expected)
	}
}
func TestUl(t *testing.T) {
	// Test empty ul
	e := bolt.Ul()
	result := e.Render()
	expected := "<ul></ul>"
	if result != expected {
		t.Fatalf(`Ul() = %q, expected %q`, result, expected)
	}

	// Test ul with one child
	e = bolt.Ul(bolt.Li(bolt.P("Item 1")))
	result = e.Render()
	expected = "<ul><li><p>Item 1</p></li></ul>"
	if result != expected {
		t.Fatalf(`Ul(Li(P("Item 1"))) = %q, expected %q`, result, expected)
	}

	// Test ul with multiple children
	e = bolt.Ul(
		bolt.Li(bolt.P("Item 1")),
		bolt.Li(bolt.P("Item 2")),
		bolt.Li(bolt.P("Item 3")),
	)
	result = e.Render()
	expected = "<ul><li><p>Item 1</p></li><li><p>Item 2</p></li><li><p>Item 3</p></li></ul>"
	if result != expected {
		t.Fatalf(`Ul(...) = %q, expected %q`, result, expected)
	}

	// Test ul with nested ul
	e = bolt.Ul(
		bolt.Li(bolt.P("Item 1")),
		bolt.Li(bolt.Ul(bolt.Li(bolt.P("Subitem 1")), bolt.Li(bolt.P("Subitem 2")))),
	)
	result = e.Render()
	expected = "<ul><li><p>Item 1</p></li><li><ul><li><p>Subitem 1</p></li><li><p>Subitem 2</p></li></ul></li></ul>"
	if result != expected {
		t.Fatalf(`Ul with nested Ul = %q, expected %q`, result, expected)
	}
}
