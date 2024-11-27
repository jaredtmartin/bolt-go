package bolt

import (
	"log"
	"net/http"
	"regexp"
	"strings"

	"golang.org/x/exp/slices"
)

type Element interface {
	GetTag() string
	Tag(tag string) Element

	GetClasses() []string
	HasClass(class string) bool
	Class(classes ...string) Element
	RemoveClass(class ...string) Element

	GetStyles() map[string]string
	GetStyle(key string) string
	Style(style string) Element
	RemoveStyle(key ...string) Element

	GetAttrs() map[string]string
	GetAttr(key string) string
	Attr(key string, value string) Element
	RemoveAttr(key ...string) Element

	GetChildren() []Element
	AddChild(child Element) Element
	PrependChild(child Element) Element
	Children(children ...Element) Element
	Content() string

	Render() string

	// Common attrs
	Text(content string) Element
	Post(url string) Element
	Get(url string) Element
	Put(url string) Element
	Patch(url string) Element
	Delete(url string) Element
	Confirm(prompt string) Element
	Target(target string) Element
	Trigger(trigger string) Element
	Swap(value string) Element
	Select(value string) Element
	Id(value string) Element
	Name(name string) Element
	Type(tipe string) Element
	Submit() Element
	Src(src string) Element
	Value(value string) Element
	Href(href string) Element
	Placeholder(placeholder string) Element
	Alt(alt string) Element
	OnClick(value string) Element
	Indicator(indicator string) Element
	Oob(oob string) Element
	SwapOob(oob string) Element
	Vals(json string) Element
	PushUrl(pushUrl ...bool) Element
	Boost(boost ...bool) Element

	Respond(w http.ResponseWriter)
	Debug(prefix ...string) Element
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

// type ElementBuilder interface {
// 	HasClass(class string) bool
// 	AddClass(class string)
// 	RemoveClass(class string)
// 	Classes() []string
// }

func (e *DefaultElement) GetTag() string {
	return e.tag
}
func (e *DefaultElement) Tag(tag string) Element {
	e.tag = tag
	return e
}
func (e *DefaultElement) add_classes(class string) {
	// split string by space
	for _, c := range strings.Split(class, " ") {
		e.classes[c] = true
	}
}
func (e *DefaultElement) GetClasses() []string {
	classes := make([]string, 0, len(e.classes))
	for class := range e.classes {
		classes = append(classes, class)
	}
	slices.Sort(classes)
	return classes
}
func (e *DefaultElement) HasClass(class string) bool {
	return e.classes[class]
}
func (e *DefaultElement) Class(classes ...string) Element {
	e.add_classes(strings.Join(classes, " "))
	return e
}
func (e *DefaultElement) RemoveClass(classes ...string) Element {
	for _, c := range classes {
		cls := strings.Split(c, " ")
		for _, v := range cls {
			delete(e.classes, strings.Trim(v, " "))
		}
	}
	return e
}
func (e *DefaultElement) render_classes() string {
	return strings.Join(e.GetClasses(), " ")
}
func (e *DefaultElement) add_styles(styles string) {
	styles_split := strings.Split(styles, ";")
	re := regexp.MustCompile(`([^"]*):([^"]*)`)
	for _, style := range styles_split {
		matched := re.FindStringSubmatch(strings.Trim(style, " "))
		if len(matched) == 3 {
			e.styles[matched[1]] = matched[2]
		}
	}
}
func (e *DefaultElement) GetStyles() map[string]string {
	return e.styles
}
func (e *DefaultElement) GetStyle(key string) string {
	return e.styles[key]
}
func (e *DefaultElement) Style(style string) Element {
	e.add_styles(style)
	return e
}
func (e *DefaultElement) RemoveStyle(styles ...string) Element {
	for _, s := range styles {
		delete(e.styles, s)
	}
	return e
}
func (e *DefaultElement) render_styles() string {
	var styles []string
	for key, value := range e.styles {
		styles = append(styles, key+":"+value+";")
	}
	slices.Sort(styles)
	return strings.Trim(strings.Join(styles, " "), " ")
}
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
func (e *DefaultElement) render_attributes() string {
	var attributes []string
	for key, value := range e.attributes {
		if key == "hx-vals" {
			attributes = append(attributes, key+"='"+value+"'")
		} else {
			attributes = append(attributes, key+"=\""+value+"\"")
		}
	}
	slices.Sort(attributes)
	if len(e.styles) > 0 {
		attributes = append(attributes, "style=\""+e.render_styles()+"\"")
	}
	if len(e.classes) > 0 {
		attributes = append(attributes, "class=\""+e.render_classes()+"\"")
	}
	return strings.Join(attributes, " ")
}
func (e *DefaultElement) GetAttrs() map[string]string {
	return e.attributes
}
func (e *DefaultElement) GetAttr(key string) string {
	return e.attributes[key]
}
func (e *DefaultElement) Attr(key string, value string) Element {
	e.add_attribute(key, value)
	return e
}
func (e *DefaultElement) RemoveAttr(keys ...string) Element {
	for _, s := range keys {
		delete(e.attributes, s)
	}
	return e
}
func (e *DefaultElement) Children(elements ...Element) Element {
	e.children = elements
	return e
}
func (e *DefaultElement) AddChild(child Element) Element {
	e.children = append(e.children, child)
	return e
}
func (e *DefaultElement) PrependChild(child Element) Element {
	e.children = append([]Element{child}, e.children...)
	return e
}
func (e *DefaultElement) Debug(prefix ...string) Element {
	s := ""
	if len(prefix) > 0 {
		s = prefix[0]
	}
	log.Printf("%s%s: c:%v s:%v a:%v text: %s\n", s, e.tag, e.classes, e.styles, e.attributes, e.text)
	// log.Printf("%sclasses: %v\n", s, e.classes)
	// log.Printf("%sstyles: %v\n", s, e.styles)
	// log.Printf("%sattributes: %v\n", s, e.attributes)
	// log.Printf("%schildren: \n", s)
	for _, child := range e.children {
		if child != nil {
			child.Debug(s + "  ")
		} else {
			log.Printf("%s  nil\n", s)
		}

	}
	return e
}
func (e *DefaultElement) GetChildren() []Element {
	return e.children
}
func (e *DefaultElement) renderChildren() string {
	var renderedStrings []string
	for _, element := range e.children {
		if element == nil {
			continue
		}
		renderedStrings = append(renderedStrings, element.Render())
	}
	return strings.Join(renderedStrings, "")
}
func (e *DefaultElement) Content() string {
	return e.renderChildren() + e.text
}
func NewElement(tag string) Element {
	element := NewDefaultElement(tag)
	return &element
}
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
func (e *DefaultElement) Text(content string) Element {
	e.text = content
	return e
}
func (e *DefaultElement) Post(url string) Element {
	e.add_attribute("hx-post", url)
	return e
}
func (e *DefaultElement) Get(url string) Element {
	e.add_attribute("hx-get", url)
	return e
}
func (e *DefaultElement) Put(url string) Element {
	e.add_attribute("hx-put", url)
	return e
}
func (e *DefaultElement) Patch(url string) Element {
	e.add_attribute("hx-patch", url)
	return e
}
func (e *DefaultElement) Delete(url string) Element {
	e.add_attribute("hx-delete", url)
	return e
}
func (e *DefaultElement) Confirm(prompt string) Element {
	e.add_attribute("hx-confirm", "true")
	e.add_attribute("hx-confirm-question", prompt)
	return e
}
func (e *DefaultElement) Target(target string) Element {
	e.add_attribute("hx-target", target)
	return e
}
func (e *DefaultElement) Trigger(trigger string) Element {
	e.add_attribute("hx-trigger", trigger)
	return e
}
func (e *DefaultElement) Swap(value string) Element {
	e.add_attribute("hx-swap", value)
	return e
}
func (e *DefaultElement) Select(value string) Element {
	e.add_attribute("hx-select", value)
	return e
}
func (e *DefaultElement) Id(value string) Element {
	e.add_attribute("id", value)
	return e
}
func (e *DefaultElement) Name(name string) Element {
	e.add_attribute("name", name)
	return e
}
func (e *DefaultElement) Type(tipe string) Element {
	e.add_attribute("type", tipe)
	return e
}
func (e *DefaultElement) Submit() Element {
	e.add_attribute("type", "submit")
	return e
}
func (e *DefaultElement) Src(src string) Element {
	e.add_attribute("src", src)
	return e
}
func (e *DefaultElement) Value(value string) Element {
	e.add_attribute("value", value)
	return e
}
func (e *DefaultElement) Href(href string) Element {
	e.add_attribute("href", href)
	e.tag = "a"
	return e
}
func (e *DefaultElement) Placeholder(placeholder string) Element {
	e.add_attribute("placeholder", placeholder)
	return e
}
func (e *DefaultElement) Alt(alt string) Element {
	e.add_attribute("alt", alt)
	return e
}
func (e *DefaultElement) OnClick(value string) Element {
	e.add_attribute("onclick", value)
	return e
}
func (e *DefaultElement) Indicator(indicator string) Element {
	e.add_attribute("hx-indicator", indicator)
	return e
}
func (e *DefaultElement) Oob(oob string) Element {
	e.add_attribute("hx-oob", oob)
	return e
}
func (e *DefaultElement) SwapOob(oob string) Element {
	e.add_attribute("hx-swap-oob", oob)
	return e
}
func (e *DefaultElement) Vals(json string) Element {
	e.add_attribute("hx-vals", json)
	return e
}
func (e *DefaultElement) PushUrl(pushUrl ...bool) Element {
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
func (e *DefaultElement) Boost(boost ...bool) Element {
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

var null_elements = [...]string{"area", "base", "br", "col", "embed", "hr", "img", "input", "link", "meta", "param", "source", "track", "wbr"}

func is_null_element(tag string) bool {
	for _, v := range null_elements {
		if v == tag {
			return true
		}
	}
	return false
}

func (e *DefaultElement) Render() string {
	children := e.renderChildren() + e.text
	if e.tag == "" {
		return children
	}
	attr := e.render_attributes()
	if len(attr) > 0 {
		attr = " " + attr
	}
	if is_null_element(e.tag) {
		return "<" + e.tag + attr + "/>"
	}
	return "<" + e.tag + attr + ">" + children + "</" + e.tag + ">"
}

// For Fiber
//
//	func (e *DefaultElement) Send(c *fiber.Ctx) error {
//		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
//		return c.SendString(e.Render())
//	}
func (e *DefaultElement) Respond(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(e.Render()))
}

// func (e *DefaultElement) Replace(key string, value string) Element {
// 	html := strings.ReplaceAll(e.children, key, value)
// 	return e.Text(html)
// }
