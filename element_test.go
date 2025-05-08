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

	// Make sure it doesnt add the class attr if there are no classes
	result = NewElement("div").Render()
	expected = "<div></div>"
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
	e.Attr("alpha")
	result := e.Render()
	expected := "<div alpha hx-post=\"yep\" hx-target=\"#target\" style=\"background: blue; border: 2px sold green; color: red;\" class=\"green yellow\"></div>"
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
	e := NewElement("div").HXPost("https://example.com")
	result := e.Render()
	expected := "<div hx-post=\"https://example.com\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// Get

func TestGet(t *testing.T) {
	e := NewElement("div").HXGet("https://example.com")
	result := e.Render()
	expected := "<div hx-get=\"https://example.com\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// Put

func TestPut(t *testing.T) {
	e := NewElement("div").HXPut("https://example.com")
	result := e.Render()
	expected := "<div hx-put=\"https://example.com\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// Patch

func TestPatch(t *testing.T) {
	e := NewElement("div").HXPatch("https://example.com")
	result := e.Render()
	expected := "<div hx-patch=\"https://example.com\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// Delete

func TestDelete(t *testing.T) {
	e := NewElement("div").HXDelete("https://example.com")
	result := e.Render()
	expected := "<div hx-delete=\"https://example.com\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// Confirm

func TestConfirm(t *testing.T) {
	e := NewElement("div").HXConfirm("Are you sure?")
	result := e.Render()
	expected := "<div hx-confirm-question=\"Are you sure?\" hx-confirm=\"true\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// Target

func TestTarget(t *testing.T) {
	e := NewElement("div").HXTarget("#target")
	result := e.Render()
	expected := "<div hx-target=\"#target\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// Trigger

func TestTrigger(t *testing.T) {
	e := NewElement("div").HXTrigger("click")
	result := e.Render()
	expected := "<div hx-trigger=\"click\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// Swap

func TestSwap(t *testing.T) {
	e := NewElement("div").HXSwap("target")
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
	e := NewElement("div").HXIndicator("target")
	result := e.Render()
	expected := "<div hx-indicator=\"target\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// Oob

func TestOob(t *testing.T) {
	e := NewElement("div").HXOob("target")
	result := e.Render()
	expected := "<div hx-oob=\"target\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}
func TestOn(t *testing.T) {
	e := NewElement("div").HXOn("click", `console.log('Hello')`)
	result := e.Render()
	expected := `<div hx-on:click="console.log('Hello')"></div>`
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// Vals

func TestVals(t *testing.T) {
	e := NewElement("div").HXVals(`{ "name": "Fred" }`)
	result := e.Render()
	expected := "<div hx-vals=\"{ &quot;name&quot;: &quot;Fred&quot; }\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// SwapOob

func TestSwapOob(t *testing.T) {
	e := NewElement("div").HXSwapOob("target")
	result := e.Render()
	expected := "<div hx-swap-oob=\"target\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// PushUrl

func TestPushUrl(t *testing.T) {
	e := NewElement("div").HXPushUrl()
	result := e.Render()
	expected := "<div hx-push-url=\"true\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
	e.HXPushUrl(false)
	result = e.Render()
	expected = "<div></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
	e.HXPushUrl(true)
	result = e.Render()
	expected = "<div hx-push-url=\"true\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

// Boost

func TestBoost(t *testing.T) {
	e := NewElement("div").HXBoost()
	result := e.Render()
	expected := "<div hx-boost=\"true\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}

	e.HXBoost(false)
	result = e.Render()
	expected = "<div></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
	e.HXBoost(true)
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
	e := NewElement("div").AtClick("increment()")
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
func TestChildWithId(t *testing.T) {
	e := NewElement("div").Id("parent").Children(
		NewElement("span").Id("child"),
		NewElement("div").Children(
			NewElement("inner").Id("deep"),
		),
	)

	result, ok := e.ChildWithId("child")
	if !ok {
		t.Fatal("Expected to find child with id 'child'")
	}
	if result.GetId() != "child" {
		t.Fatalf("Expected child id to be 'child', got %s", result.GetId())
	}

	result, ok = e.ChildWithId("deep")
	if !ok {
		t.Fatal("Expected to find child with id 'deep'")
	}
	if result.GetId() != "deep" {
		t.Fatalf("Expected child id to be 'child', got %s", result.GetId())
	}
	if result.GetTag() != "inner" {
		t.Fatalf("Expected child tag to be 'inner', got %s", result.GetTag())
	}

	_, ok = e.ChildWithId("nonexistent")
	if ok {
		t.Fatal("Expected false when searching for nonexistent child id")
	}
}

func TestChildrenWithClass(t *testing.T) {
	e := NewElement("div")
	child1 := NewElement("p").Class("highlight")
	child2 := NewElement("span").Class("highlight bold")
	child3 := NewElement("div").Class("normal")
	e.Children(child1, child2, child3)

	results := e.GetChildrenWithClass("highlight")
	if len(results) != 2 {
		t.Fatalf("Expected 2 children with class 'highlight', got %d", len(results))
	}

	results = e.GetChildrenWithClass("nonexistent")
	if len(results) != 0 {
		t.Fatal("Expected empty slice for nonexistent class")
	}
}

func TestFirstChildWithClass(t *testing.T) {
	e := NewElement("div")
	child0 := NewElement("span").Class("green")
	child1 := NewElement("p").Class("highlight")
	child2 := NewElement("span").Class("highlight")
	e.Children(child0, child1, child2)

	result, ok := e.GetFirstChildWithClass("highlight")
	if !ok {
		t.Fatal("Expected to find first child with class 'highlight'")
	}
	if result.GetTag() != "p" {
		t.Fatalf("Expected first child to be 'p', got %s", result.GetTag())
	}

	_, ok = e.GetFirstChildWithClass("nonexistent")
	if ok {
		t.Fatal("Expected false when searching for nonexistent class")
	}
}

func TestNthChildWithClass(t *testing.T) {
	e := NewElement("div").Children(
		NewElement("first").Class("green"),
		NewElement("second").Class("highlight"),
		NewElement("third").Class("green"),
		NewElement("fourth").Class("highlight"),
		NewElement("fifth").Class("highlight"),
		NewElement("parent").Children(
			NewElement("child-first").Class("green"),
			NewElement("child-second").Class("highlight"),
			NewElement("child-third").Class("green"),
			NewElement("child-fourth").Class("highlight"),
		),
	)
	result, ok := e.GetNthChildWithClass("highlight", 1)
	if !ok {
		t.Fatal("Expected to find first child with class 'highlight'")
	}
	if result.GetTag() != "second" {
		t.Fatalf("Expected first child to be 'second', got %s", result.GetTag())
	}

	result, ok = e.GetNthChildWithClass("highlight", 3)
	if !ok {
		t.Fatal("Expected to find third child with class 'highlight'")
	}
	if result.GetTag() != "fifth" {
		t.Fatalf("Expected third child to be 'fifth', got %s", result.GetTag())
	}

	result, ok = e.GetNthChildWithClass("highlight", 4)
	if !ok {
		t.Fatal("Expected to find deeper child with class 'highlight'")
	}
	if result.GetTag() != "child-second" {
		t.Fatalf("Expected deeper child to be 'child-second', got %s", result.GetTag())
	}

	_, ok = e.GetNthChildWithClass("highlight", 6)
	if ok {
		t.Fatal("Expected false when index out of bounds")
	}
	_, ok = e.GetNthChildWithClass("highlight", -5)
	if ok {
		t.Fatal("Expected false when index is negative")
	}

	_, ok = e.GetNthChildWithClass("nonexistent", 0)
	if ok {
		t.Fatal("Expected false when searching for nonexistent class")
	}
}

func TestPrependChild(t *testing.T) {
	e := NewElement("div")
	child1 := NewElement("p").Text("first")
	child2 := NewElement("span").Text("second")

	e.Children(child1)
	e.PrependChild(child2)

	result := e.Render()
	expected := "<div><span>second</span><p>first</p></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

func TestAddChild(t *testing.T) {
	e := NewElement("div")
	child1 := NewElement("p").Text("first")
	child2 := NewElement("span").Text("second")

	e.AddChild(child1)
	e.AddChild(child2)

	result := e.Render()
	expected := "<div><p>first</p><span>second</span></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

func TestGetTag(t *testing.T) {
	e := NewElement("div")
	if e.GetTag() != "div" {
		t.Fatalf("Expected tag 'div', got %s", e.GetTag())
	}

	e.Tag("span")
	if e.GetTag() != "span" {
		t.Fatalf("Expected tag 'span', got %s", e.GetTag())
	}
}

func TestGetStyle(t *testing.T) {
	e := NewElement("div").Style("color: red; font-size: 12px")

	if e.GetStyle("color") != "red" {
		t.Fatalf("Expected style 'color' to be 'red', got %s", e.GetStyle("color"))
	}

	if e.GetStyle("nonexistent") != "" {
		t.Fatal("Expected empty string for nonexistent style")
	}
}

func TestSelect(t *testing.T) {
	e := NewElement("div").HXSelect("#target")
	result := e.Render()
	expected := "<div hx-select=\"#target\"></div>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

func TestSubmit(t *testing.T) {
	e := NewElement("button").Submit()
	result := e.Render()
	expected := "<button type=\"submit\"></button>"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

func TestAlt(t *testing.T) {
	e := NewElement("img").Alt("description")
	result := e.Render()
	expected := "<img alt=\"description\">"
	if result != expected {
		t.Fatalf(`result = %q, expected %q`, result, expected)
	}
}

func TestContent(t *testing.T) {
	e := NewElement("div").Text("Hello World")
	if e.GetContent() != "Hello World" {
		t.Fatalf("Expected content 'Hello World', got %s", e.GetContent())
	}

	e = NewElement("div")
	if e.GetContent() != "" {
		t.Fatalf("Expected empty content, got %s", e.GetContent())
	}
}
func TestGetId(t *testing.T) {
	// Test with explicitly set ID
	e := NewElement("div").Id("test-id")
	result := e.GetId()
	if result != "test-id" {
		t.Fatalf(`GetId() = %q, expected "test-id"`, result)
	}

	// Test with no ID set
	e = NewElement("div")
	result = e.GetId()
	if result != "" {
		t.Fatalf(`GetId() = %q, expected empty string`, result)
	}

	// Test after removing ID
	e = NewElement("div").Id("test-id")
	e.RemoveAttr("id")
	result = e.GetId()
	if result != "" {
		t.Fatalf(`GetId() = %q, expected empty string after removal`, result)
	}

}
func TestGetStyles(t *testing.T) {
	// Test empty styles
	e := NewElement("div")
	styles := e.GetStyles()
	if len(styles) != 0 {
		t.Fatalf("Expected empty styles map, got %d styles", len(styles))
	}

	// Test single style
	e = NewElement("div").Style("color: red")
	styles = e.GetStyles()
	if len(styles) != 1 {
		t.Fatalf("Expected 1 style, got %d", len(styles))
	}
	if styles["color"] != "red" {
		t.Fatalf("Expected color:red, got color:%s", styles["color"])
	}

	// Test multiple styles
	e = NewElement("div").Style("color: red; font-size: 12px; margin: 5px")
	styles = e.GetStyles()
	if len(styles) != 3 {
		t.Fatalf("Expected 3 styles, got %d", len(styles))
	}
	expectedStyles := map[string]string{
		"color":     "red",
		"font-size": "12px",
		"margin":    "5px",
	}
	for key, value := range expectedStyles {
		if styles[key] != value {
			t.Fatalf("Expected %s:%s, got %s:%s", key, value, key, styles[key])
		}
	}

	// Test overwriting styles
	e = NewElement("div").Style("color: red").Style("color: blue")
	styles = e.GetStyles()
	if len(styles) != 1 {
		t.Fatalf("Expected 1 style after overwrite, got %d", len(styles))
	}
	if styles["color"] != "blue" {
		t.Fatalf("Expected color:blue after overwrite, got color:%s", styles["color"])
	}

	// Test after removing style
	e = NewElement("div").Style("color: red; font-size: 12px")
	e.RemoveStyle("color")
	styles = e.GetStyles()
	if len(styles) != 1 {
		t.Fatalf("Expected 1 style after removal, got %d", len(styles))
	}
	if _, exists := styles["color"]; exists {
		t.Fatal("Style 'color' should not exist after removal")
	}
	if styles["font-size"] != "12px" {
		t.Fatalf("Expected font-size:12px to remain, got font-size:%s", styles["font-size"])
	}
}
func TestGetAttrs(t *testing.T) {
	// Test empty attributes
	e := NewElement("div")
	attrs := e.GetAttrs()
	if len(attrs) != 0 {
		t.Fatalf("Expected empty attributes map, got %d attributes", len(attrs))
	}

	// Test single attribute
	e = NewElement("div").Attr("data-test", "value")
	attrs = e.GetAttrs()
	if len(attrs) != 1 {
		t.Fatalf("Expected 1 attribute, got %d", len(attrs))
	}
	if attrs["data-test"] != "value" {
		t.Fatalf("Expected data-test:value, got data-test:%s", attrs["data-test"])
	}

	// Test multiple attributes
	e = NewElement("div").
		Attr("data-test", "value").
		Attr("aria-label", "label").
		Attr("role", "button")
	attrs = e.GetAttrs()
	if len(attrs) != 3 {
		t.Fatalf("Expected 3 attributes, got %d", len(attrs))
	}
	expectedAttrs := map[string]string{
		"data-test":  "value",
		"aria-label": "label",
		"role":       "button",
	}
	for key, value := range expectedAttrs {
		if attrs[key] != value {
			t.Fatalf("Expected %s:%s, got %s:%s", key, value, key, attrs[key])
		}
	}

	// Test overwriting attributes
	e = NewElement("div").
		Attr("data-test", "first").
		Attr("data-test", "second")
	attrs = e.GetAttrs()
	if len(attrs) != 1 {
		t.Fatalf("Expected 1 attribute after overwrite, got %d", len(attrs))
	}
	if attrs["data-test"] != "second" {
		t.Fatalf("Expected data-test:second after overwrite, got data-test:%s", attrs["data-test"])
	}

	// Test after removing attribute
	e = NewElement("div").
		Attr("data-test", "value").
		Attr("aria-label", "label")
	e.RemoveAttr("data-test")
	attrs = e.GetAttrs()
	if len(attrs) != 1 {
		t.Fatalf("Expected 1 attribute after removal, got %d", len(attrs))
	}
	if _, exists := attrs["data-test"]; exists {
		t.Fatal("Attribute 'data-test' should not exist after removal")
	}
	if attrs["aria-label"] != "label" {
		t.Fatalf("Expected aria-label:label to remain, got aria-label:%s", attrs["aria-label"])
	}
}
