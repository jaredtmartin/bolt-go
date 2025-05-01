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

func TestTemplate(t *testing.T) {
	result := Template("sample").Render()
	expected := "<p>Hello Html!</p>"
	if result != expected {
		t.Fatalf(`Template("sample") = %q, expected %q`, result, expected)
	}
	result = Template("sample.html").Render()
	if result != expected {
		t.Fatalf(`Template("sample.html") = %q, expected %q`, result, expected)
	}
	result = Template("sample.tmpl").Render()
	expected = "<p>Hello Tmpl!</p>"
	if result != expected {
		t.Fatalf(`Template("sample.tmpl") = %q, expected %q`, result, expected)
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
	e := Script("console.log('test')").Src("/src/test.js")
	result := e.Render()
	expected := `<script src="/src/test.js">console.log('test')</script>`
	if result != expected {
		t.Fatalf(`Script() = %q, expected %q`, result, expected)
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
