package bolt

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"golang.org/x/exp/slices"
)

type Element interface {
	// Returns the HTML tag name of the element.
	GetTag() string
	// Sets the HTML tag name of the element.
	Tag(tag string) Element
	// Sets the defer attribute of the element.
	Defer(value ...bool) Element
	// Returns the id of the element.
	GetId() string
	// Returns the first child element with the given id recursively.
	ChildWithId(id string) (Element, bool)

	// Returns an array of all the classes of the element.
	GetClasses() []string
	// Returns true if the element has the given class.
	HasClass(class string) bool
	// Adds the given class or classes to the element.
	Class(classes ...string) Element
	// Removes the given class or classes from the element.
	RemoveClass(class ...string) Element
	// Returns a map of the styles on the element.
	GetStyles() map[string]string
	// Returns the value of the given style key.
	GetStyle(key string) string
	// Sets the given style key to the given value.
	Style(style string) Element
	// Removes the given style key from the element.
	RemoveStyle(key ...string) Element

	// Returns a map of the attributes on the element.
	GetAttrs() map[string]string
	// Returns the value of the given attribute key.
	GetAttr(key string) string
	// Sets the given attribute key to the given value.
	Attr(key string, value ...string) Element
	// Removes the given attribute key from the element.
	RemoveAttr(key ...string) Element
	// Returns the children of the element.
	GetChildren() []Element
	// Returns the child element at the given index.
	GetChild(idx int) (Element, bool)
	// Appends the element as a child.
	AddChild(child Element) Element
	// Prepends the element as a child.
	PrependChild(child Element) Element
	// Sets the children of the element.
	Children(children ...Element) Element
	// Returns all descendants with the given class.
	GetChildrenWithClass(class string) []Element
	// Returns the first descendant with the given class.
	GetFirstChildWithClass(class string) (Element, bool)
	// Returns the nth descendant with the given class.
	GetNthChildWithClass(class string, n int) (Element, bool)
	// Returns the rendered children and text.
	GetContent() string
	// Renders the element as HTML string.
	Render() string

	// Common attrs

	// Sets the text content of the element.
	Text(content string) Element
	// Sets the HTML content of the element without escaping.
	UnsafeHtml(content string) Element
	// Sets the hx-post attribute of the element.
	HXPost(url string) Element
	// Sets the hx-get attribute of the element.
	HXGet(url string) Element
	// Sets the hx-put attribute of the element.
	HXPut(url string) Element
	// Sets the hx-patch attribute of the element.
	HXPatch(url string) Element
	// Sets the hx-delete attribute of the element.
	HXDelete(url string) Element
	// Sets the hx-confirm attribute of the element.
	HXConfirm(prompt string) Element
	// Sets the hx-target attribute of the element.
	HXTarget(target string) Element
	// Sets the hx-trigger attribute of the element.
	HXTrigger(trigger string) Element
	// Sets the hx-swap attribute of the element.
	HXSwap(value string) Element
	// Sets the hx-select attribute of the element.
	HXSelect(value string) Element
	// Sets the id attribute of the element.
	Id(value string) Element
	// Sets the name attribute of the element.
	Name(name string) Element
	// Sets the for attribute of the element.
	For(name string) Element

	// Sets the type attribute of the element.
	Type(tipe string) Element

	// Sets the type attribute to "submit".
	Submit() Element

	// Sets the method attribute of the element to the given string.
	Method(method string) Element

	// Sets the src attribute of the element.
	Src(src string) Element
	// Sets the value attribute of the element.
	Value(value string) Element
	// Sets the href attribute of the element.
	Href(href string) Element
	// Sets the placeholder attribute of the element.
	Placeholder(placeholder string) Element
	// Sets the alt attribute of the element.
	Alt(alt string) Element
	// Sets the onclick attribute of the element.
	OnClick(value string) Element
	// Sets the hx-indicator attribute of the element.
	HXIndicator(indicator string) Element
	// Sets the hx-oob attribute of the element.
	HXOob(oob string) Element
	// Sets the hx-swap-oob attribute of the element.
	HXSwapOob(oob string) Element
	// Sets the hx-vals attribute of the element.
	HXVals(json string) Element
	// Sets the hx-push-url attribute of the element.
	HXPushUrl(pushUrl ...bool) Element
	// Sets hx-on attribute for htmx
	HXOn(event string, value string) Element
	// Sets the hx-boost attribute of the element.
	HXBoost(boost ...bool) Element
	// Sets x-data attribute for alpine.js
	XData(data string) Element
	// Sets @click attribute for alpine.js
	AtClick(handler string) Element
	// Sets x-bind attribute for alpine.js
	XBind(attr string, value string) Element
	// Sets x-on attribute for alpine.js
	XOn(event string, value string) Element
	// Sets x-text attribute for alpine.js
	XText(value string) Element
	// Sets x-html attribute for alpine.js
	XHtml(value string) Element
	// Sets x-model attribute for alpine.js
	XModel(value string) Element
	//  Sets x-show attribute for alpine.js
	XShow(condition string) Element
	// Sets x-init attribute for alpine.js
	XInit(value string) Element
	// Sets x-ref attribute for alpine.js
	XRef(value string) Element
	// Sets x-cloak attribute for alpine.js
	XCloak() Element

	Send(w http.ResponseWriter)
	// Redirect(w http.ResponseWriter, url string)
	// Debug(prefix ...string) Element
}

type DefaultElement struct {
	tag        string
	classes    map[string]bool
	styles     map[string]string
	attributes map[string]string
	children   []Element
	text       string
}

// type Element struct {
// 	tag        string
// 	classes    map[string]bool
// 	styles     map[string]string
// 	attributes map[string]string
// 	children   []*Element
// 	text       string
// }

//	type ElementBuilder interface {
//		HasClass(class string) bool
//		AddClass(class string)
//		RemoveClass(class string)
//		Classes() []string
//	}
var htmlEscaper = strings.NewReplacer(
	`&`, "&amp;",
	// `'`, "&#39;", // "&#39;" is shorter than "&apos;" and apos was not in HTML until HTML5.
	`<`, "&lt;",
	`>`, "&gt;",
	`"`, "&quot;", // "&#34;" is shorter than "&quot;".
)

// Escapes a string to ensure it is safe HTML.
// Doesnt remove apostrophe
func EscapeString(s string) string {
	return htmlEscaper.Replace(s)
}

// Returns the HTML tag name of the element.
func (e *DefaultElement) GetTag() string {
	return e.tag
}

// Sets the HTML tag name of the element.
func (e *DefaultElement) Tag(tag string) Element {
	e.tag = EscapeString(tag)
	return e
}

// Adds one or more classes separated by spaces to the class attribute of the element.
// Removes duplicate classes
// Removes empty strings
func (e *DefaultElement) add_classes(class string) {
	// split string by space
	for _, c := range strings.Split(class, " ") {
		if c == "" {
			continue
		}
		e.classes[c] = true
	}
}

// Returns an array of all the classes of the element.
func (e *DefaultElement) GetClasses() []string {
	classes := make([]string, 0, len(e.classes))
	for class := range e.classes {
		classes = append(classes, class)
	}
	slices.Sort(classes)
	return classes
}

// Returns true if the element has the given class.
func (e *DefaultElement) HasClass(class string) bool {
	return e.classes[class]
}

// Adds the given class or classes to the element.
func (e *DefaultElement) Class(classes ...string) Element {
	e.add_classes(strings.Join(classes, " "))
	return e
}

// Returns the id of the element.
func (e *DefaultElement) GetId() string {
	return e.attributes["id"]
}

// Removes the given class or classes from the element.
func (e *DefaultElement) RemoveClass(classes ...string) Element {
	for _, c := range classes {
		cls := strings.Split(c, " ")
		for _, v := range cls {
			delete(e.classes, strings.Trim(v, " "))
		}
	}
	return e
}

// renders the elements classes into a space separated string
func (e *DefaultElement) render_classes() string {
	return strings.Join(e.GetClasses(), " ")
}

// Adds CSS styles to the element
func (e *DefaultElement) add_styles(styles string) {
	styles_split := strings.Split(styles, ";")
	re := regexp.MustCompile(`([^"]*):([^"]*)`)
	for _, style := range styles_split {
		matched := re.FindStringSubmatch(strings.Trim(style, " "))
		if len(matched) == 3 {
			e.styles[matched[1]] = strings.Trim(matched[2], " ")
		}
	}
}

// Returns a map of the styles on the element.
func (e *DefaultElement) GetStyles() map[string]string {
	return e.styles
}

// Returns the value of the given style key.
func (e *DefaultElement) GetStyle(key string) string {
	return e.styles[key]
}

// Sets the given style key to the given value.
func (e *DefaultElement) Style(style string) Element {
	e.add_styles(style)
	return e
}

// Removes the given style key from the element.
func (e *DefaultElement) RemoveStyle(styles ...string) Element {
	for _, s := range styles {
		delete(e.styles, s)
	}
	return e
}

// Renders the internal map of styles into a proper CSS string
func (e *DefaultElement) render_styles() string {
	var styles []string
	for key, value := range e.styles {
		styles = append(styles, key+": "+value+";")
	}
	slices.Sort(styles)
	return strings.Trim(strings.Join(styles, " "), " ")
}

// Adds an attribute to the element. If the key is "style" or "class", it will be combined with existing style or class attributes.
func (e *DefaultElement) add_attribute(key string, value string) {
	if key == "style" {
		e.styles = make(map[string]string)
		e.add_styles(value)
	} else if key == "class" {
		e.classes = make(map[string]bool)
		e.add_classes(value)
	} else {
		e.attributes[key] = value
	}
}

// Renders the internal map of attributes into a proper HTML string including class and style
func (e *DefaultElement) render_attributes() string {
	var attributes []string
	for key, value := range e.attributes {
		// attributes = append(attributes, key+"='"+value+"'")
		if value == "" {
			attributes = append(attributes, key)
			continue
		}
		attributes = append(attributes, EscapeString(key)+"=\""+EscapeString(value)+"\"")
	}
	slices.Sort(attributes)
	if len(e.styles) > 0 {
		attributes = append(attributes, "style=\""+EscapeString(e.render_styles())+"\"")
	}
	if len(e.classes) > 0 {
		attributes = append(attributes, "class=\""+EscapeString(e.render_classes())+"\"")
	}
	return strings.Join(attributes, " ")
}

// Returns a map of the attributes on the element.
func (e *DefaultElement) GetAttrs() map[string]string {
	return e.attributes
}

// Returns the value of the given attribute key.
func (e *DefaultElement) GetAttr(key string) string {
	return e.attributes[key]
}

// Sets the given attribute key to the given value.
func (e *DefaultElement) Attr(key string, value ...string) Element {
	v := ""
	if len(value) > 0 {
		v = value[0]
	}
	e.add_attribute(key, v)
	return e
}

// Removes the given attribute key from the element.
func (e *DefaultElement) RemoveAttr(keys ...string) Element {
	for _, s := range keys {
		delete(e.attributes, s)
	}
	return e
}

// Sets the children of the element.
func (e *DefaultElement) Children(elements ...Element) Element {
	e.children = elements
	return e
}

// Appends the element as a child.
func (e *DefaultElement) AddChild(child Element) Element {
	e.children = append(e.children, child)
	return e
}

// Prepends the element as a child.
func (e *DefaultElement) PrependChild(child Element) Element {
	e.children = append([]Element{child}, e.children...)
	return e
}

// Returns all descendants with the given class.
func (e *DefaultElement) GetChildrenWithClass(class string) []Element {
	results := make([]Element, 0)
	for _, child := range e.children {
		if child.HasClass(class) {
			results = append(results, child)
		}
		results = append(results, child.GetChildrenWithClass(class)...)
	}
	return results
}

// Returns the nth descendant with the given class.
func (e *DefaultElement) GetNthChildWithClass(class string, n int) (Element, bool) {
	if n < 1 {
		return nil, false
	}
	results := make([]Element, 0)
	for _, child := range e.children {
		if child.HasClass(class) {
			results = append(results, child)
		}
		if len(results) >= n {
			return results[n-1], true
		}
		results = append(results, child.GetChildrenWithClass(class)...)
		if len(results) >= n {
			return results[n-1], true
		}
	}
	return nil, false
}

// Returns the first descendant with the given class.
func (e *DefaultElement) GetFirstChildWithClass(class string) (Element, bool) {
	return e.GetNthChildWithClass(class, 1)
}

// Returns the first child element with the given id recursively.
func (e *DefaultElement) ChildWithId(id string) (Element, bool) {
	for _, child := range e.children {
		if child.GetId() == id {
			return child, true
		}
		res, exists := child.ChildWithId(id)
		if exists {
			return res, true
		}
	}
	return nil, false
}

// Returns the children of the element.
func (e *DefaultElement) GetChildren() []Element {
	return e.children
}

// Returns the child element at the given index.
func (e *DefaultElement) GetChild(idx int) (Element, bool) {
	if idx < 0 || idx >= len(e.children) {
		return nil, false
	}
	return e.children[idx], true
}

// Returns the rendered children and text.
func (e *DefaultElement) GetContent() string {
	return renderElements(e.children...) + e.text
}

// Renders an array of elements into a string
func renderElements(elements ...Element) string {
	var renderedStrings []string
	for _, element := range elements {
		if element == nil {
			continue
		}
		renderedStrings = append(renderedStrings, element.Render())
	}
	return strings.Join(renderedStrings, "")
}

// Returns a new Element with the given tag.
func NewElement(tag string) Element {
	element := NewDefaultElement(tag)
	return &element
}

// Returns a new DefaultElement with the given tag.
func NewDefaultElement(tag string) DefaultElement {
	return DefaultElement{
		tag:        tag,
		classes:    make(map[string]bool),
		styles:     make(map[string]string),
		attributes: make(map[string]string),
		children:   []Element{},
		text:       "",
	}
}

// Sets the text content of the element.
func (e *DefaultElement) Text(content string) Element {
	e.text = EscapeString(content)
	return e
}

// Sets the text content of the element without escaping.
func (e *DefaultElement) UnsafeHtml(content string) Element {
	e.text = content
	return e
}

// Sets the hx-post attribute of the element.
func (e *DefaultElement) HXPost(url string) Element {
	e.add_attribute("hx-post", url)
	return e
}

// Sets the hx-get attribute of the element.
func (e *DefaultElement) HXGet(url string) Element {
	e.add_attribute("hx-get", url)
	return e
}

// Sets the hx-put attribute of the element.
func (e *DefaultElement) HXPut(url string) Element {
	e.add_attribute("hx-put", url)
	return e
}

// Sets the hx-patch attribute of the element.
func (e *DefaultElement) HXPatch(url string) Element {
	e.add_attribute("hx-patch", url)
	return e
}

// Sets the hx-delete attribute of the element.
func (e *DefaultElement) HXDelete(url string) Element {
	e.add_attribute("hx-delete", url)
	return e
}

// Sets the hx-confirm attribute of the element.
func (e *DefaultElement) HXConfirm(prompt string) Element {
	e.add_attribute("hx-confirm", "true")
	e.add_attribute("hx-confirm-question", prompt)
	return e
}

// Sets the hx-target attribute of the element.
func (e *DefaultElement) HXTarget(target string) Element {
	e.add_attribute("hx-target", target)
	return e
}

// Sets the hx-trigger attribute of the element.
func (e *DefaultElement) HXTrigger(trigger string) Element {
	e.add_attribute("hx-trigger", trigger)
	return e
}

// Sets the hx-swap attribute of the element.
func (e *DefaultElement) HXSwap(value string) Element {
	e.add_attribute("hx-swap", value)
	return e
}

// Sets the hx-select attribute of the element.
func (e *DefaultElement) HXSelect(value string) Element {
	e.add_attribute("hx-select", value)
	return e
}

// Sets the id attribute of the element.
func (e *DefaultElement) Id(value string) Element {
	e.add_attribute("id", value)
	return e
}

// Sets the name attribute of the element.
func (e *DefaultElement) Name(name string) Element {
	e.add_attribute("name", name)
	return e
}

// Sets the for attribute of the element.
func (e *DefaultElement) For(value string) Element {
	e.add_attribute("for", value)
	return e
}

// Sets the type attribute of the element.
func (e *DefaultElement) Type(tipe string) Element {
	e.add_attribute("type", tipe)
	return e
}

// Sets the type attribute to "submit".
func (e *DefaultElement) Submit() Element {
	e.add_attribute("type", "submit")
	return e
}

// Sets the src attribute of the element.
func (e *DefaultElement) Src(src string) Element {
	e.add_attribute("src", src)
	return e
}

// Sets the value attribute of the element.
func (e *DefaultElement) Value(value string) Element {
	e.add_attribute("value", value)
	return e
}

// Sets the href attribute of the element.
func (e *DefaultElement) Href(href string) Element {
	e.add_attribute("href", href)
	e.tag = "a"
	return e
}

// Sets the placeholder attribute of the element.
func (e *DefaultElement) Placeholder(placeholder string) Element {
	e.add_attribute("placeholder", placeholder)
	return e
}

// Sets the alt attribute of the element.
func (e *DefaultElement) Alt(alt string) Element {
	e.add_attribute("alt", alt)
	return e
}

// Sets the onclick attribute of the element.
func (e *DefaultElement) OnClick(value string) Element {
	e.add_attribute("onclick", value)
	return e
}

// Sets the hx-indicator attribute of the element.
func (e *DefaultElement) HXIndicator(indicator string) Element {
	e.add_attribute("hx-indicator", indicator)
	return e
}

// Sets the hx-oob attribute of the element.
func (e *DefaultElement) HXOob(oob string) Element {
	e.add_attribute("hx-oob", oob)
	return e
}

// Sets the hx-swap-oob attribute of the element.
func (e *DefaultElement) HXSwapOob(oob string) Element {
	e.add_attribute("hx-swap-oob", oob)
	return e
}

// Sets the hx-vals attribute of the element.
func (e *DefaultElement) HXVals(json string) Element {
	e.add_attribute("hx-vals", json)
	return e
}

// Sets the hx-push-url attribute of the element.
func (e *DefaultElement) HXPushUrl(pushUrl ...bool) Element {
	if len(pushUrl) > 0 {
		if pushUrl[0] {
			e.add_attribute("hx-push-url", "true")
		} else {
			delete(e.attributes, "hx-push-url")
		}
	} else {
		e.add_attribute("hx-push-url", "true")
	}
	return e
}

// Sets hx-on attribute for htmx
func (e *DefaultElement) HXOn(event string, script string) Element {
	e.add_attribute(fmt.Sprintf("hx-on:%s", event), script)
	return e
}

// Sets the hx-boost attribute of the element.
func (e *DefaultElement) HXBoost(boost ...bool) Element {
	if len(boost) > 0 {
		if boost[0] {
			e.add_attribute("hx-boost", "true")
		} else {
			delete(e.attributes, "hx-boost")
		}
	} else {
		e.add_attribute("hx-boost", "true")
	}
	return e
}

// Sets the defer attribute of the element.
func (e *DefaultElement) Defer(value ...bool) Element {
	if len(value) > 0 {
		if value[0] {
			e.add_attribute("defer", "true")
		} else {
			delete(e.attributes, "defer")
		}
	} else {
		e.add_attribute("defer", "true")
	}
	return e
}

// Sets x-data attribute for alpine.js
func (e *DefaultElement) XData(data string) Element {
	e.add_attribute("x-data", data)
	return e
}

// Sets @click attribute for alpine.js
func (e *DefaultElement) AtClick(handler string) Element {
	e.add_attribute("@click", handler)
	return e
}

// Sets x-bind attribute for alpine.js
func (e *DefaultElement) XBind(attr string, value string) Element {
	e.add_attribute(fmt.Sprintf("x-bind:%s", attr), value)
	return e
}

// Sets x-on attribute for alpine.js
func (e *DefaultElement) XOn(event string, value string) Element {
	e.add_attribute(fmt.Sprintf("x-on:%s", event), value)
	return e
}

// Sets x-text attribute for alpine.js
func (e *DefaultElement) XText(value string) Element {
	e.add_attribute("x-text", value)
	return e
}

// Sets x-html attribute for alpine.js
func (e *DefaultElement) XHtml(value string) Element {
	e.add_attribute("x-html", value)
	return e
}

// Sets x-model attribute for alpine.js
func (e *DefaultElement) XModel(value string) Element {
	e.add_attribute("x-model", value)
	return e
}

// Sets x-show attribute for alpine.js
func (e *DefaultElement) XShow(condition string) Element {
	e.add_attribute("x-show", condition)
	return e
}

// Sets x-init attribute for alpine.js
func (e *DefaultElement) XInit(value string) Element {
	e.add_attribute("x-init", value)
	return e
}

// Sets x-cloak attribute for alpine.js
func (e *DefaultElement) XCloak() Element {
	e.add_attribute("x-cloak", "")
	return e
}

// Sets x-ref attribute for alpine.js
func (e *DefaultElement) XRef(value string) Element {
	e.add_attribute("x-ref", value)
	return e
}

// list of null elements that normally should not have content
var null_elements = [...]string{"area", "base", "br", "col", "embed", "hr", "img", "input", "link", "meta", "param", "source", "track", "wbr"}

// Returns true if the element is a null element.
func is_null_element(tag string) bool {
	for _, v := range null_elements {
		if v == tag {
			return true
		}
	}
	return false
}

// Renders the element with the children first and text at the end.
func (e *DefaultElement) RenderWithContent(text string, children ...Element) string {
	content := renderElements(children...) + text
	if e.tag == "" {
		return content
	}
	attr := e.render_attributes()
	if len(attr) > 0 {
		attr = " " + attr
	}
	if is_null_element(e.tag) && len(content) == 0 && len(e.text) == 0 {
		return "<" + e.tag + attr + ">"
	}
	return "<" + e.tag + attr + ">" + content + "</" + e.tag + ">"
}

// Renders the element as HTML string.
func (e *DefaultElement) Render() string {
	return e.RenderWithContent(e.text, e.children...)
}

// Sends the element to the http.ResponseWriter as text/html.
func (e *DefaultElement) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(e.Render()))
}

// Sets the method attribute of the element to the given string.
func (e *DefaultElement) Method(method string) Element {
	e.add_attribute("method", method)
	return e
}

// func (e *DefaultElement) Replace(key string, value string) Element {
// 	html := strings.ReplaceAll(e.children, key, value)
// 	return e.Text(html)
// }
