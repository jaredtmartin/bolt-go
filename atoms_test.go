package bolt

import (
	"testing"
)

func TestDiv(t *testing.T) {
	e := Div()
	result := e.Render()
	expected := "<div/>"
	if result != expected {
		t.Fatalf(`Div() = %q, want %q`, result, expected)
	}
}

func TestSpan(t *testing.T) {
	e := Span("Hello")
	result := e.Render()
	expected := "<span>Hello</span>"
	if result != expected {
		t.Fatalf(`Span() = %q, want %q`, result, expected)
	}
}

func TestP(t *testing.T) {
	e := P("Hello")
	result := e.Render()
	expected := "<p>Hello</p>"
	if result != expected {
		t.Fatalf(`P() = %q, want %q`, result, expected)
	}
}

func TestH1(t *testing.T) {
	e := H1("Hello")
	result := e.Render()
	expected := "<h1>Hello</h1>"
	if result != expected {
		t.Fatalf(`H1() = %q, want %q`, result, expected)
	}
}

func TestH2(t *testing.T) {
	e := H2("Hello")
	result := e.Render()
	expected := "<h2>Hello</h2>"
	if result != expected {
		t.Fatalf(`H2() = %q, want %q`, result, expected)
	}
}

func TestH3(t *testing.T) {
	e := H3("Hello")
	result := e.Render()
	expected := "<h3>Hello</h3>"
	if result != expected {
		t.Fatalf(`H3() = %q, want %q`, result, expected)
	}
}

func TestH4(t *testing.T) {
	e := H4("Hello")
	result := e.Render()
	expected := "<h4>Hello</h4>"
	if result != expected {
		t.Fatalf(`H4() = %q, want %q`, result, expected)
	}
}

func TestH5(t *testing.T) {
	e := H5("Hello")
	result := e.Render()
	expected := "<h5>Hello</h5>"
	if result != expected {
		t.Fatalf(`H5() = %q, want %q`, result, expected)
	}
}

func TestH6(t *testing.T) {
	e := H6("Hello")
	result := e.Render()
	expected := "<h6>Hello</h6>"
	if result != expected {
		t.Fatalf(`H6() = %q, want %q`, result, expected)
	}
}

// Img
func TestImg(t *testing.T) {
	e := Img("src")
	result := e.Render()
	expected := "<img src=\"src\"/>"
	if result != expected {
		t.Fatalf(`Img() = %q, want %q`, result, expected)
	}
}

// Button
func TestButton(t *testing.T) {
	e := Button("Hello")
	result := e.Render()
	expected := "<button>Hello</button>"
	if result != expected {
		t.Fatalf(`Button() = %q, want %q`, result, expected)
	}
}

// Label
func TestLabel(t *testing.T) {
	e := Label("Hello")
	result := e.Render()
	expected := "<label>Hello</label>"
	if result != expected {
		t.Fatalf(`Label() = %q, want %q`, result, expected)
	}
}

// Input
func TestInput(t *testing.T) {
	e := Input("first_name")
	result := e.Render()
	expected := "<input name=\"first_name\"/>"
	if result != expected {
		t.Fatalf(`Input() = %q, want %q`, result, expected)
	}
}

// Form
func TestForm(t *testing.T) {
	e := Form().Children(Input("name"))
	result := e.Render()
	expected := "<form><input name=\"name\"/></form>"
	if result != expected {
		t.Fatalf(`Form() = %q, want %q`, result, expected)
	}
}

// A
func TestA(t *testing.T) {
	e := A("www.google.com", Html("Hello"))
	result := e.Render()
	expected := "<a href=\"www.google.com\">Hello</a>"
	if result != expected {
		t.Fatalf(`A() = %q, want %q`, result, expected)
	}
}

// None
func TestNone(t *testing.T) {
	e := None()
	result := e.Render()
	expected := ""
	if result != expected {
		t.Fatalf(`None() = %q, want %q`, result, expected)
	}
}

// Fragment
func TestFragment(t *testing.T) {
	e := Fragment()
	result := e.Render()
	expected := ""
	if result != expected {
		t.Fatalf(`Fragment() = %q, want %q`, result, expected)
	}
}

// Html
func TestHtml(t *testing.T) {
	e := Html("<div>Hello</div>")
	result := e.Render()
	expected := "<div>Hello</div>"
	if result != expected {
		t.Fatalf(`Html() = %q, want %q`, result, expected)
	}
}
