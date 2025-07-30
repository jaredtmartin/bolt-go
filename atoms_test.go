package bolt

import (
	"testing"
)

func TestDiv(t *testing.T) {
	e := Div("")
	result := e.Render()
	expected := "<div></div>"
	if result != expected {
		t.Fatalf(`Div(") = %q, expected %q`, result, expected)
	}
}

func TestSpan(t *testing.T) {
	e := Span("Hello")
	result := e.Render()
	expected := "<span>Hello</span>"
	if result != expected {
		t.Fatalf(`Span() = %q, expected %q`, result, expected)
	}
}

func TestP(t *testing.T) {
	e := P("Hello")
	result := e.Render()
	expected := "<p>Hello</p>"
	if result != expected {
		t.Fatalf(`P() = %q, expected %q`, result, expected)
	}
}

func TestH1(t *testing.T) {
	e := H1("Hello")
	result := e.Render()
	expected := "<h1>Hello</h1>"
	if result != expected {
		t.Fatalf(`H1() = %q, expected %q`, result, expected)
	}
}

func TestH2(t *testing.T) {
	e := H2("Hello")
	result := e.Render()
	expected := "<h2>Hello</h2>"
	if result != expected {
		t.Fatalf(`H2() = %q, expected %q`, result, expected)
	}
}

func TestH3(t *testing.T) {
	e := H3("Hello")
	result := e.Render()
	expected := "<h3>Hello</h3>"
	if result != expected {
		t.Fatalf(`H3() = %q, expected %q`, result, expected)
	}
}

func TestH4(t *testing.T) {
	e := H4("Hello")
	result := e.Render()
	expected := "<h4>Hello</h4>"
	if result != expected {
		t.Fatalf(`H4() = %q, expected %q`, result, expected)
	}
}

func TestH5(t *testing.T) {
	e := H5("Hello")
	result := e.Render()
	expected := "<h5>Hello</h5>"
	if result != expected {
		t.Fatalf(`H5() = %q, expected %q`, result, expected)
	}
}

func TestH6(t *testing.T) {
	e := H6("Hello")
	result := e.Render()
	expected := "<h6>Hello</h6>"
	if result != expected {
		t.Fatalf(`H6() = %q, expected %q`, result, expected)
	}
}

// Img
func TestImg(t *testing.T) {
	e := Img("src")
	result := e.Render()
	expected := "<img src=\"src\">"
	if result != expected {
		t.Fatalf(`Img() = %q, expected %q`, result, expected)
	}
}

// Button
func TestButton(t *testing.T) {
	e := Button("Hello")
	result := e.Render()
	expected := "<button type=\"button\">Hello</button>"
	if result != expected {
		t.Fatalf(`Button() = %q, expected %q`, result, expected)
	}
}

// Label
func TestLabel(t *testing.T) {
	e := Label("Hello")
	result := e.Render()
	expected := "<label>Hello</label>"
	if result != expected {
		t.Fatalf(`Label() = %q, expected %q`, result, expected)
	}
}

// Input
func TestInput(t *testing.T) {
	e := Input("first_name")
	result := e.Render()
	expected := "<input name=\"first_name\" type=\"text\">"
	if result != expected {
		t.Fatalf(`Input() = %q, expected %q`, result, expected)
	}
}

// Form
func TestForm(t *testing.T) {
	e := Form().Children(Input("name"))
	result := e.Render()
	expected := "<form><input name=\"name\" type=\"text\"></form>"
	if result != expected {
		t.Fatalf(`Form() = %q, expected %q`, result, expected)
	}
}

// A
func TestA(t *testing.T) {
	e := A("www.google.com", Html("Hello"))
	result := e.Render()
	expected := "<a href=\"www.google.com\">Hello</a>"
	if result != expected {
		t.Fatalf(`A() = %q, expected %q`, result, expected)
	}
}

// None
func TestNone(t *testing.T) {
	e := None()
	result := e.Render()
	expected := ""
	if result != expected {
		t.Fatalf(`None() = %q, expected %q`, result, expected)
	}
}

// Fragment
func TestFragment(t *testing.T) {
	e := Fragment()
	result := e.Render()
	expected := ""
	if result != expected {
		t.Fatalf(`Fragment() = %q, expected %q`, result, expected)
	}
}

// Html
func TestUnsafeHtmlElement(t *testing.T) {
	e := Html("<div>Hello</div>")
	result := e.Render()
	expected := "<div>Hello</div>"
	if result != expected {
		t.Fatalf(`Html() = %q, expected %q`, result, expected)
	}
}
func TestSection(t *testing.T) {
	e := Section(P("Hello"))
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
	result := Template(P("Content")).Render()
	expected := "<template><p>Content</p></template>"
	if result != expected {
		t.Fatalf(`Template(P("Content")) = %q, expected %q`, result, expected)
	}
}
func TestSvg(t *testing.T) {
	e := Svg("M10 10 H 90 V 90 H 10 Z", 100, 100, 100, 100)
	result := e.Render()
	expected := `<svg height="100" version="1.1" viewBox="0 0 100 100" width="100" xmlns:xlink="http://www.w3.org/1999/xlink" xmlns="http://www.w3.org/2000/svg" class="fill-current">M10 10 H 90 V 90 H 10 Z</svg>`
	if result != expected {
		t.Fatalf(`Svg() = %q, expected %q`, result, expected)
	}
	e = Svg(`<path d="M15 19l-7-7 7-7" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"></path>`, 24, 24, 24, 24).Class("fill-current")
	result = e.Render()
	expected = `<svg height="24" version="1.1" viewBox="0 0 24 24" width="24" xmlns:xlink="http://www.w3.org/1999/xlink" xmlns="http://www.w3.org/2000/svg" class="fill-current"><path d="M15 19l-7-7 7-7" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"></path></svg>`
	if result != expected {
		t.Fatalf(`Svg() = %q, expected %q`, result, expected)
	}
}
func TestVideoIframe(t *testing.T) {
	e := VideoIframe("Hello", 23, 25, "https://www.youtube.com/embed/dQw4w9WgXcQ")
	result := e.Render()
	expected := `<iframe allow="autoplay; fullscreen; picture-in-picture; clipboard-write" data-ready="true" frameborder="0" height="25" src="https://www.youtube.com/embed/dQw4w9WgXcQ" title="Hello" width="23" class="h-full w-full"></iframe>`
	if result != expected {
		t.Fatalf(`Svg() = %q, expected %q`, result, expected)
	}
}
func TestScript(t *testing.T) {
	// Test embedded script
	e := Script("console.log('test')")
	result := e.Render()
	expected := `<script>console.log('test')</script>`
	if result != expected {
		t.Fatalf(`Script("console.log('test')") = %q, expected %q`, result, expected)
	}
	// Test linked script
	e = Script("").Src("/src/test.js")
	result = e.Render()
	expected = `<script src="/src/test.js"></script>`
	if result != expected {
		t.Fatalf(`Script("/src/test.js") = %q, expected %q`, result, expected)
	}
	// Test linked script that is deferred
	e = Script("").Src("/src/test.js").Defer()
	result = e.Render()
	expected = `<script defer="true" src="/src/test.js"></script>`
	if result != expected {
		t.Fatalf(`Script("/src/test.js") = %q, expected %q`, result, expected)
	}
}
func TestHead(t *testing.T) {
	// Test empty head
	e := Head()
	result := e.Render()
	expected := "<head></head>"
	if result != expected {
		t.Fatalf(`Head() = %q, expected %q`, result, expected)
	}

	// Test head with single child
	e = Head(Script("console.log('test')"))
	result = e.Render()
	expected = "<head><script>console.log('test')</script></head>"
	if result != expected {
		t.Fatalf(`Head() with single child = %q, expected %q`, result, expected)
	}

	// Test head with multiple children
	e = Head(
		Script("console.log('test')"),
		Html("<meta charset=\"UTF-8\">"),
		Html("<title>Test Page</title>"),
	)
	result = e.Render()
	expected = "<head><script>console.log('test')</script><meta charset=\"UTF-8\"><title>Test Page</title></head>"
	if result != expected {
		t.Fatalf(`Head() with multiple children = %q, expected %q`, result, expected)
	}
}
func TestBody(t *testing.T) {
	// Test empty body
	e := Body()
	result := e.Render()
	expected := "<body></body>"
	if result != expected {
		t.Fatalf(`Body() = %q, expected %q`, result, expected)
	}

	// Test body with single child
	e = Body(P("Hello"))
	result = e.Render()
	expected = "<body><p>Hello</p></body>"
	if result != expected {
		t.Fatalf(`Body() = %q, expected %q`, result, expected)
	}

	// Test body with multiple children
	e = Body(H1("Title"), P("Paragraph"), Div("", P("Content")))
	result = e.Render()
	expected = "<body><h1>Title</h1><p>Paragraph</p><div><p>Content</p></div></body>"
	if result != expected {
		t.Fatalf(`Body() = %q, expected %q`, result, expected)
	}

	// Test body with nested elements
	e = Body(Div("", P("Nested")))
	result = e.Render()
	expected = "<body><div><p>Nested</p></div></body>"
	if result != expected {
		t.Fatalf(`Body() = %q, expected %q`, result, expected)
	}
}
func TestStylesheet(t *testing.T) {
	// Test basic stylesheet
	e := Stylesheet("/styles/main.css")
	result := e.Render()
	expected := `<link href="/styles/main.css" rel="stylesheet">`
	if result != expected {
		t.Fatalf(`Stylesheet("/styles/main.css") = %q, expected %q`, result, expected)
	}
}

func TestHiddenInput(t *testing.T) {
	e := HiddenInput("name", "value")
	result := e.Render()
	expected := `<input name="name" type="hidden" value="value">`
	if result != expected {
		t.Fatalf(`HiddenInput("name", "value") = %q, expected %q`, result, expected)
	}
}

func TestLi(t *testing.T) {
	e := Li(P("Content"))
	result := e.Render()
	expected := `<li><p>Content</p></li>`
	if result != expected {
		t.Fatalf(`Li(P("Content")) = %q, expected %q`, result, expected)
	}
}

func TestTable(t *testing.T) {
	e := Table(
		Tr(
			Th(P("Header 1")),
			Th(P("Header 2")),
		),
		Tr(
			Td(P("Row 1, Cell 1")),
			Td(P("Row 1, Cell 2")),
		),
	)
	result := e.Render()
	expected := `<table><tr><th><p>Header 1</p></th><th><p>Header 2</p></th></tr><tr><td><p>Row 1, Cell 1</p></td><td><p>Row 1, Cell 2</p></td></tr></table>`
	if result != expected {
		t.Fatalf(`Table() = %q, expected %q`, result, expected)
	}
}
func TestTr(t *testing.T) {
	e := Tr(
		Td(P("Row 1, Cell 1")),
		Td(P("Row 1, Cell 2")),
	)
	result := e.Render()
	expected := `<tr><td><p>Row 1, Cell 1</p></td><td><p>Row 1, Cell 2</p></td></tr>`
	if result != expected {
		t.Fatalf(`Tr() = %q, expected %q`, result, expected)
	}
}
func TestTh(t *testing.T) {
	e := Th(
		Td(P("Header 1")),
		Td(P("Header 2")),
	)
	result := e.Render()
	expected := `<th><td><p>Header 1</p></td><td><p>Header 2</p></td></th>`
	if result != expected {
		t.Fatalf(`Th() = %q, expected %q`, result, expected)
	}
}
func TestTd(t *testing.T) {
	e := Td(P("Content"))
	result := e.Render()
	expected := `<td><p>Content</p></td>`
	if result != expected {
		t.Fatalf(`Td(P("Content")) = %q, expected %q`, result, expected)
	}
}
func TestString(t *testing.T) {
	e := Template(P("Content"))
	result := e.Render()
	expected := `<template><p>Content</p></template>`
	if result != expected {
		t.Fatalf(`Template(P("Content")) = %q, expected %q`, result, expected)
	}
}

func TestOobElement(t *testing.T) {
	e := Oob("swap-value", P("Content"))
	result := e.Render()
	expected := `<div hx-swap-oob="swap-value"><p>Content</p></div>`
	if result != expected {
		t.Fatalf(`Oob("swap-value", P("Content")) = %q, expected %q`, result, expected)
	}
}
func TestStyle(t *testing.T) {
	// Test with simple CSS
	e := Style("body { background: #fff; }")
	result := e.Render()
	expected := `<style>body { background: #fff; }</style>`
	if result != expected {
		t.Fatalf(`Style("body { background: #fff; }") = %q, expected %q`, result, expected)
	}

	// Test with empty CSS
	e = Style("")
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
	e = Style(css)
	result = e.Render()
	expected = `<style>` + css + `</style>`
	if result != expected {
		t.Fatalf(`Style(multiline) = %q, expected %q`, result, expected)
	}
}
func TestUl(t *testing.T) {
	// Test empty ul
	e := Ul()
	result := e.Render()
	expected := "<ul></ul>"
	if result != expected {
		t.Fatalf(`Ul() = %q, expected %q`, result, expected)
	}

	// Test ul with one child
	e = Ul(Li(P("Item 1")))
	result = e.Render()
	expected = "<ul><li><p>Item 1</p></li></ul>"
	if result != expected {
		t.Fatalf(`Ul(Li(P("Item 1"))) = %q, expected %q`, result, expected)
	}

	// Test ul with multiple children
	e = Ul(
		Li(P("Item 1")),
		Li(P("Item 2")),
		Li(P("Item 3")),
	)
	result = e.Render()
	expected = "<ul><li><p>Item 1</p></li><li><p>Item 2</p></li><li><p>Item 3</p></li></ul>"
	if result != expected {
		t.Fatalf(`Ul(...) = %q, expected %q`, result, expected)
	}

	// Test ul with nested ul
	e = Ul(
		Li(P("Item 1")),
		Li(Ul(Li(P("Subitem 1")), Li(P("Subitem 2")))),
	)
	result = e.Render()
	expected = "<ul><li><p>Item 1</p></li><li><ul><li><p>Subitem 1</p></li><li><p>Subitem 2</p></li></ul></li></ul>"
	if result != expected {
		t.Fatalf(`Ul with nested Ul = %q, expected %q`, result, expected)
	}
}
