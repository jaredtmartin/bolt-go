package bolt

import (
	"testing"
)

func TestDiv(t *testing.T) {
	e := Div()
	result := e.Render()
	expected := "<div></div>"
	if result != expected {
		t.Fatalf(`Div() = %q, expected %q`, result, expected)
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
	expected := "<img src=\"src\"/>"
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
	expected := "<input name=\"first_name\" type=\"text\"/>"
	if result != expected {
		t.Fatalf(`Input() = %q, expected %q`, result, expected)
	}
}

// Form
func TestForm(t *testing.T) {
	e := Form().Children(Input("name"))
	result := e.Render()
	expected := "<form><input name=\"name\" type=\"text\"/></form>"
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
func TestHtml(t *testing.T) {
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
