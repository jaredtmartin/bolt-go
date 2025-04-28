package bolt

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestClasses(t *testing.T) {
	e := NewElement("div").Class("red green blue").Class("green yellow")

	result := e.Render()
	expected := "<div class=\"blue green red yellow\"></div>"
	// fmt.Println("result", result)
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}
func TestRender(t *testing.T) {
	result := NewElement("script").Render()
	expected := "<script></script>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
	// Make sure unsafe HTML characters are escaped
	result = NewElement("script").Text("<script>alert('hello')</script>").Render()
	expected = "<script>&lt;script&gt;alert('hello')&lt;/script&gt;</script>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}
func TestRenderNullElement(t *testing.T) {
	result := NewElement("img").Render()
	expected := "<img>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}
func TestRenderNullElementWithText(t *testing.T) {
	result := NewElement("img").Text("Hello").Render()
	expected := "<img>Hello</img>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}
func TestRenderNullElementWithChildren(t *testing.T) {
	result := NewElement("img").Children(NewElement("p").Text("Hello")).Render()
	expected := "<img><p>Hello</p></img>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}
func TestLink(t *testing.T) {
	result := NewElement("link").Render()
	expected := "<link>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}
func TestRemoveClasses(t *testing.T) {
	e := NewElement("div").Class("red green blue yellow")
	e.RemoveClass("green yellow")
	result := e.Render()
	expected := "<div class=\"blue red\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}
func TestAddStyles(t *testing.T) {
	e := NewElement("div")
	e.Style("color: red; background: blue; border: 2px sold green")
	e.Style("color: purple; font-weight: bold")
	result := e.Render()
	expected := "<div style=\"background: blue; border: 2px sold green; color: purple; font-weight: bold;\"></div>"
	if result != expected {
		t.Fatalf(`Styles() = %q, expected %q`, result, expected)
	}
}
func TestAttributes(t *testing.T) {
	e := NewElement("div")
	e.Attr("style", "color: red; background: blue; border: 2px sold green")
	e.Attr("class", "green yellow")
	e.Attr("hx-post", "nope")
	e.Attr("hx-target", "#target")
	e.Attr("hx-post", "yep")
	result := e.Render()
	expected := "<div hx-post=\"yep\" hx-target=\"#target\" style=\"background: blue; border: 2px sold green; color: red;\" class=\"green yellow\"></div>"
	if result != expected {
		t.Fatalf(`Attributes() = %q, expected %q`, result, expected)
	}
}
func TestAttributesWithoutClasses(t *testing.T) {
	e := NewElement("div")
	e.Attr("style", "color: red; background: blue; border: 2px sold green")
	e.Attr("hx-post", "nope")
	result := e.Render()
	expected := "<div hx-post=\"nope\" style=\"background: blue; border: 2px sold green; color: red;\"></div>"
	if result != expected {
		t.Fatalf(`Attributes() = %q, expected %q`, result, expected)
	}
}
func TestAttributesWithoutStyles(t *testing.T) {
	e := NewElement("div")
	e.Attr("class", "green yellow")
	e.Attr("hx-post", "nope")
	result := e.Render()
	expected := "<div hx-post=\"nope\" class=\"green yellow\"></div>"
	if result != expected {
		t.Fatalf(`Attributes() = %q, expected %q`, result, expected)
	}
}
func TestText(t *testing.T) {
	e := NewElement("div")
	e.Text("hello")
	result := e.Render()
	expected := "<div>hello</div>"
	if result != expected {
		t.Fatalf(`Text() = %q, expected %q`, result, expected)
	}
}
func TestChild(t *testing.T) {
	e := NewElement("div")
	e.Children(NewElement("p").Text("hello"))
	result := e.Render()
	expected := "<div><p>hello</p></div>"
	if result != expected {
		t.Fatalf(`Text() = %q, expected %q`, result, expected)
	}
}
func TestPuttingItAllTogether(t *testing.T) {
	e := NewElement("div").Class("red m-1 p-2").Style("color: red; background: blue; border: 2px sold green")
	e.Children(NewElement("p").Text("hello"))
	result := e.Render()
	expected := "<div style=\"background: blue; border: 2px sold green; color: red;\" class=\"m-1 p-2 red\"><p>hello</p></div>"
	if result != expected {
		t.Fatalf(`Text() = %q, expected %q`, result, expected)
	}
}
func TestRemoveStyle(t *testing.T) {
	e := NewElement("div").Style("color: red; background: green;")
	e.RemoveStyle("color")
	result := e.Render()
	expected := "<div style=\"background: green;\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}
func TestRemoveAttr(t *testing.T) {
	e := NewElement("div").Attr("hx-boost", "true").Attr("id", "input")
	e.RemoveAttr("hx-boost")
	result := e.Render()
	expected := "<div id=\"input\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// Post

func TestPost(t *testing.T) {
	e := NewElement("div").Post("https://example.com")
	result := e.Render()
	expected := "<div hx-post=\"https://example.com\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// Get

func TestGet(t *testing.T) {
	e := NewElement("div").Get("https://example.com")
	result := e.Render()
	expected := "<div hx-get=\"https://example.com\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// Put

func TestPut(t *testing.T) {
	e := NewElement("div").Put("https://example.com")
	result := e.Render()
	expected := "<div hx-put=\"https://example.com\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// Patch

func TestPatch(t *testing.T) {
	e := NewElement("div").Patch("https://example.com")
	result := e.Render()
	expected := "<div hx-patch=\"https://example.com\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// Delete

func TestDelete(t *testing.T) {
	e := NewElement("div").Delete("https://example.com")
	result := e.Render()
	expected := "<div hx-delete=\"https://example.com\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// Confirm

func TestConfirm(t *testing.T) {
	e := NewElement("div").Confirm("Are you sure?")
	result := e.Render()
	expected := "<div hx-confirm-question=\"Are you sure?\" hx-confirm=\"true\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// Target

func TestTarget(t *testing.T) {
	e := NewElement("div").Target("#target")
	result := e.Render()
	expected := "<div hx-target=\"#target\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// Trigger

func TestTrigger(t *testing.T) {
	e := NewElement("div").Trigger("click")
	result := e.Render()
	expected := "<div hx-trigger=\"click\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// Swap

func TestSwap(t *testing.T) {
	e := NewElement("div").Swap("target")
	result := e.Render()
	expected := "<div hx-swap=\"target\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// Id

func TestId(t *testing.T) {
	e := NewElement("div").Id("target")
	result := e.Render()
	expected := "<div id=\"target\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// Name

func TestName(t *testing.T) {
	e := NewElement("div").Name("target")
	result := e.Render()
	expected := "<div name=\"target\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// For

func TestFor(t *testing.T) {
	e := NewElement("div").For("target")
	result := e.Render()
	expected := "<div for=\"target\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// Type

func TestType(t *testing.T) {
	e := NewElement("div").Type("target")
	result := e.Render()
	expected := "<div type=\"target\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// Src

func TestSrc(t *testing.T) {
	e := NewElement("div").Src("target")
	result := e.Render()
	expected := "<div src=\"target\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// Value

func TestValue(t *testing.T) {
	e := NewElement("div").Value("target")
	result := e.Render()
	expected := "<div value=\"target\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// Href

func TestHref(t *testing.T) {
	e := NewElement("div").Href("target")
	result := e.Render()
	expected := "<a href=\"target\"></a>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// Placeholder

func TestPlaceholder(t *testing.T) {
	e := NewElement("div").Placeholder("target")
	result := e.Render()
	expected := "<div placeholder=\"target\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// OnClick

func TestOnClick(t *testing.T) {
	e := NewElement("div").OnClick("target")
	result := e.Render()
	expected := "<div onclick=\"target\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// Indicator

func TestIndicator(t *testing.T) {
	e := NewElement("div").Indicator("target")
	result := e.Render()
	expected := "<div hx-indicator=\"target\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// Oob

func TestOob(t *testing.T) {
	e := NewElement("div").Oob("target")
	result := e.Render()
	expected := "<div hx-oob=\"target\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}
func TestOn(t *testing.T) {
	e := NewElement("div").On("click", `console.log('Hello')`)
	result := e.Render()
	expected := `<div hx-on:click="console.log('Hello')"></div>`
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// Vals

func TestVals(t *testing.T) {
	e := NewElement("div").Vals(`{ "name": "Fred" }`)
	result := e.Render()
	expected := "<div hx-vals=\"{ &quot;name&quot;: &quot;Fred&quot; }\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// SwapOob

func TestSwapOob(t *testing.T) {
	e := NewElement("div").SwapOob("target")
	result := e.Render()
	expected := "<div hx-swap-oob=\"target\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// PushUrl

func TestPushUrl(t *testing.T) {
	e := NewElement("div").PushUrl()
	result := e.Render()
	expected := "<div hx-push-url=\"true\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
	e.PushUrl(false)
	result = e.Render()
	expected = "<div></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
	e.PushUrl(true)
	result = e.Render()
	expected = "<div hx-push-url=\"true\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// Boost

func TestBoost(t *testing.T) {
	e := NewElement("div").Boost()
	result := e.Render()
	expected := "<div hx-boost=\"true\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}

	e.Boost(false)
	result = e.Render()
	expected = "<div></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
	e.Boost(true)
	result = e.Render()
	expected = "<div hx-boost=\"true\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

func TestHasClass(t *testing.T) {
	e := NewElement("div").Class("red green blue")
	result := e.HasClass("red")
	expected := true
	if result != expected {
		t.Fatalf(`result = %v, expected %v`, result, expected)
	}
	result = e.HasClass("purple")
	expected = false
	if result != expected {
		t.Fatalf(`result = %v, expected %v`, result, expected)
	}
}

// func TestSend(t *testing.T) {
// 	app := fiber.New()
// 	app.Get("/hello", func(c *fiber.Ctx) error {
// 		e := NewElement("div").Class("red green blue")
// 		return e.Send(c)
// 	})
// 	req := httptest.NewRequest("GET", "/hello", nil)
// 	resp, _ := app.Test(req, -1)
// 	assert.Equalf(t, 200, resp.StatusCode, "should get valid response")
// 	b, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

//		body := string(b)
//		assert.Equalf(t, "<div class=\"blue green red\"></div>", body, "should get valid body")
//		assert.Equalf(t, "text/html", resp.Header.Get("Content-Type"), "should get html")
//	}
func TestXData(t *testing.T) {
	e := NewElement("div").XData("{count: 0}")
	result := e.Render()
	expected := "<div x-data=\"{count: 0}\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}
func TestClick(t *testing.T) {
	e := NewElement("div").Click("increment()")
	result := e.Render()
	expected := "<div @click=\"increment()\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}
func TestXBind(t *testing.T) {
	e := NewElement("div").XBind("class", "isActive ? 'active' : ''")
	result := e.Render()
	expected := "<div x-bind:class=\"isActive ? 'active' : ''\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}
func TestXOn(t *testing.T) {
	e := NewElement("div").XOn("click", "handleClick()")
	result := e.Render()
	expected := "<div x-on:click=\"handleClick()\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}
func TestXText(t *testing.T) {
	e := NewElement("div").XText("count")
	result := e.Render()
	expected := "<div x-text=\"count\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}
func TestXHtml(t *testing.T) {
	e := NewElement("div").XHtml("count")
	result := e.Render()
	expected := "<div x-html=\"count\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}
func TestXModel(t *testing.T) {
	e := NewElement("input").XModel("user.name")
	result := e.Render()
	expected := "<input x-model=\"user.name\">"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}
func TestXShow(t *testing.T) {
	e := NewElement("div").XShow("isVisible")
	result := e.Render()
	expected := "<div x-show=\"isVisible\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}
func TestXInit(t *testing.T) {
	e := NewElement("div").XInit("setup()")
	result := e.Render()
	expected := "<div x-init=\"setup()\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}
func TestXCloak(t *testing.T) {
	e := NewElement("div").XCloak()
	result := e.Render()
	expected := "<div x-cloak></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}
func TestUnsafeHtml(t *testing.T) {
	e := NewElement("div").UnsafeHtml("<p>Hello <strong>World</strong></p>")
	result := e.Render()
	expected := "<div><p>Hello <strong>World</strong></p></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}

	e = NewElement("div").UnsafeHtml("<script>alert('xss')</script>")
	result = e.Render()
	expected = "<div><script>alert('xss')</script></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}

	e = NewElement("div").UnsafeHtml("Multiple\nLines\nOf\nContent")
	result = e.Render()
	expected = "<div>Multiple\nLines\nOf\nContent</div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}

}
func TestGetChildren(t *testing.T) {
	// Test empty children
	e := NewElement("div")
	children := e.GetChildren()
	if len(children) != 0 {
		t.Fatalf("Expected 0 children, got %d", len(children))
	}

	// Test single child
	child := NewElement("p").Text("hello")
	e.Children(child)
	children = e.GetChildren()
	if len(children) != 1 {
		t.Fatalf("Expected 1 child, got %d", len(children))
	}
	if children[0].Render() != "<p>hello</p>" {
		t.Fatalf("Expected child to render as <p>hello</p>, got %s", children[0].Render())
	}

	// Test multiple children
	e = NewElement("div")
	child1 := NewElement("p").Text("first")
	child2 := NewElement("span").Text("second")
	child3 := NewElement("div").Text("third")
	e.Children(child1, child2, child3)
	children = e.GetChildren()
	if len(children) != 3 {
		t.Fatalf("Expected 3 children, got %d", len(children))
	}
	expected := []string{
		"<p>first</p>",
		"<span>second</span>",
		"<div>third</div>",
	}
	for i, child := range children {
		if child.Render() != expected[i] {
			t.Fatalf("Child %d: expected %s, got %s", i, expected[i], child.Render())
		}
	}

	// Test nested children
	e = NewElement("div")
	nested := NewElement("div").Children(
		NewElement("p").Text("nested"),
		NewElement("span").Text("content"),
	)
	e.Children(nested)
	children = e.GetChildren()
	if len(children) != 1 {
		t.Fatalf("Expected 1 child, got %d", len(children))
	}
	expectedNested := "<div><p>nested</p><span>content</span></div>"
	if children[0].Render() != expectedNested {
		t.Fatalf("Expected nested children to render as %s, got %s", expectedNested, children[0].Render())
	}
}
func TestGetChild(t *testing.T) {
	// Test getting child at valid index
	e := NewElement("div")
	child1 := NewElement("p").Text("first")
	child2 := NewElement("span").Text("second")
	e.Children(child1, child2)

	result, ok := e.GetChild(0)
	if !ok {
		t.Fatal("Expected to get child at index 0")
	}
	if result.Render() != "<p>first</p>" {
		t.Fatalf("Expected first child to render as <p>first</p>, got %s", result.Render())
	}

	result, ok = e.GetChild(1)
	if !ok {
		t.Fatal("Expected to get child at index 1")
	}
	if result.Render() != "<span>second</span>" {
		t.Fatalf("Expected second child to render as <span>second</span>, got %s", result.Render())
	}

	// Test getting child at invalid index
	_, ok = e.GetChild(-1)
	if ok {
		t.Fatal("Expected false when getting child at negative index")
	}

	_, ok = e.GetChild(2)
	if ok {
		t.Fatal("Expected false when getting child at out of bounds index")
	}

	// Test getting child from empty element
	empty := NewElement("div")
	_, ok = empty.GetChild(0)
	if ok {
		t.Fatal("Expected false when getting child from empty element")
	}

	// Test getting child from element with nested children
	nested := NewElement("div").Children(
		NewElement("div").Children(
			NewElement("p").Text("nested"),
		),
	)
	result, ok = nested.GetChild(0)
	if !ok {
		t.Fatal("Expected to get nested child")
	}
	if result.Render() != "<div><p>nested</p></div>" {
		t.Fatalf("Expected nested child to render as <div><p>nested</p></div>, got %s", result.Render())
	}
}
