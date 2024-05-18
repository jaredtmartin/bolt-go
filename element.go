package bolt

import (
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
)

type Element struct {
	tag        string
	classes    map[string]bool
	styles     map[string]string
	attributes map[string]string
	children   string
}

// type ElementBuilder interface {
// 	HasClass(class string) bool
// 	AddClass(class string)
// 	RemoveClass(class string)
// 	Classes() []string
// }

func (e *Element) add_classes(class string) {
	// split string by space
	for _, c := range strings.Split(class, " ") {
		e.classes[c] = true
	}
}

func (e *Element) HasClass(class string) bool {
	return e.classes[class]
}

func (e *Element) get_classes() string {
	classes := make([]string, 0, len(e.classes))
	for class := range e.classes {
		classes = append(classes, class)
	}
	slices.Sort(classes)
	return strings.Join(classes, " ")
}

func (e *Element) add_styles(styles string) {
	styles_split := strings.Split(styles, ";")
	re := regexp.MustCompile(`([^"]*):([^"]*)`)
	for _, style := range styles_split {
		matched := re.FindStringSubmatch(strings.Trim(style, " "))
		if len(matched) == 3 {
			e.styles[matched[1]] = matched[2]
		}
	}
}
func (e *Element) get_styles() string {
	var styles []string
	for key, value := range e.styles {
		styles = append(styles, key+":"+value+";")
	}
	slices.Sort(styles)
	return strings.Trim(strings.Join(styles, " "), " ")
}
func (e *Element) add_attribute(key string, value string) {
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
func (e *Element) get_attributes() string {
	var attributes []string
	for key, value := range e.attributes {
		attributes = append(attributes, key+"=\""+value+"\"")
	}
	slices.Sort(attributes)
	if len(e.styles) > 0 {
		attributes = append(attributes, "style=\""+e.get_styles()+"\"")
	}
	if len(e.classes) > 0 {
		attributes = append(attributes, "class=\""+e.get_classes()+"\"")
	}
	return strings.Join(attributes, " ")
}
func NewElement(tag string) *Element {
	return &Element{
		tag:        tag,
		classes:    make(map[string]bool),
		styles:     make(map[string]string),
		attributes: make(map[string]string),
		children:   "",
	}
}
func (e *Element) RemoveClasses(classes ...string) *Element {
	for _, c := range classes {
		cls := strings.Split(c, " ")
		for _, v := range cls {
			delete(e.classes, strings.Trim(v, " "))
		}
	}
	return e
}
func (e *Element) RemoveStyles(styles ...string) *Element {
	for _, s := range styles {
		delete(e.styles, s)
	}
	return e
}
func (e *Element) RemoveAttributes(keys ...string) *Element {
	for _, s := range keys {
		delete(e.attributes, s)
	}
	return e
}
func (e *Element) Class(classes string) *Element {
	e.add_classes(classes)
	return e
}
func (e *Element) Style(styles string) *Element {
	e.add_styles(styles)
	return e
}
func (e *Element) Attr(key string, value string) *Element {
	e.add_attribute(key, value)
	return e
}
func (e *Element) Tag(tag string) *Element {
	e.tag = tag
	return e
}
func (e *Element) Text(content string) *Element {
	e.children = content
	return e
}
func (e *Element) Post(url string) *Element {
	e.add_attribute("hx-post", url)
	return e
}
func (e *Element) Get(url string) *Element {
	e.add_attribute("hx-get", url)
	return e
}
func (e *Element) Put(url string) *Element {
	e.add_attribute("hx-put", url)
	return e
}
func (e *Element) Patch(url string) *Element {
	e.add_attribute("hx-patch", url)
	return e
}
func (e *Element) Delete(url string) *Element {
	e.add_attribute("hx-delete", url)
	return e
}
func (e *Element) Confirm(prompt string) *Element {
	e.add_attribute("hx-confirm", "true")
	e.add_attribute("hx-confirm-question", prompt)
	return e
}
func (e *Element) Target(target string) *Element {
	e.add_attribute("hx-target", target)
	return e
}
func (e *Element) Trigger(trigger string) *Element {
	e.add_attribute("hx-trigger", trigger)
	return e
}
func (e *Element) Swap(value string) *Element {
	e.add_attribute("hx-swap", value)
	return e
}
func (e *Element) Id(value string) *Element {
	e.add_attribute("id", value)
	return e
}
func (e *Element) Name(name string) *Element {
	e.add_attribute("name", name)
	return e
}
func (e *Element) Type(tipe string) *Element {
	e.add_attribute("type", tipe)
	return e
}
func (e *Element) Src(src string) *Element {
	e.add_attribute("src", src)
	return e
}
func (e *Element) Value(value string) *Element {
	e.add_attribute("value", value)
	return e
}
func (e *Element) Href(href string) *Element {
	e.add_attribute("href", href)
	e.tag = "a"
	return e
}
func (e *Element) Placeholder(placeholder string) *Element {
	e.add_attribute("placeholder", placeholder)
	return e
}
func (e *Element) OnClick(value string) *Element {
	e.add_attribute("onclick", value)
	return e
}
func (e *Element) Indicator(indicator string) *Element {
	e.add_attribute("hx-indicator", indicator)
	return e
}
func (e *Element) Oob(oob string) *Element {
	e.add_attribute("hx-oob", oob)
	return e
}
func (e *Element) SwapOob(oob string) *Element {
	e.add_attribute("hx-swap-oob", oob)
	return e
}
func (e *Element) PushUrl(pushUrl ...bool) *Element {
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
func (e *Element) Boost(boost ...bool) *Element {
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

func (e *Element) Children(elements ...*Element) *Element {
	var renderedStrings []string
	for _, element := range elements {
		renderedStrings = append(renderedStrings, element.Render())
	}
	e.children = strings.Join(renderedStrings, "")
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
func (e *Element) Render() string {
	if e.tag == "" {
		return e.children
	}
	attr := e.get_attributes()
	if len(attr) > 0 {
		attr = " " + attr
	}
	if is_null_element(e.tag) {
		return "<" + e.tag + attr + "/>"
	}
	return "<" + e.tag + attr + ">" + e.children + "</" + e.tag + ">"
}

// For Fiber
func (e *Element) Send(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.SendString(e.Render())
}
func (e *Element) Replace(key string, value string) *Element {
	html := strings.ReplaceAll(e.children, key, value)
	return e.Text(html)
}
