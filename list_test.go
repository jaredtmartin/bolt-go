package bolt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestNewList(t *testing.T) {
	l := List()
	var strs []interface{}
	strs = append(strs, "apple", "bear", "coffee")
	result := l.Render(strs...)
	assert.Equalf(t, `<div><p>apple</p><p>bear</p><p>coffee</p></div>`, result, "should match")
}
func TestNewListWithRenderRow(t *testing.T) {
	l := List()
	l.renderRow = func(item interface{}) *Element {
		return Div().Text(fmt.Sprintf("%v", item))
	}
	var strs []interface{}
	strs = append(strs, "apple", "bear", "coffee")
	result := l.Render(strs...)
	assert.Equalf(t, `<div><div>apple</div><div>bear</div><div>coffee</div></div>`, result, "should match")
}
func TestListElement(t *testing.T) {
	l := List()
	var strs []interface{}
	strs = append(strs, "apple", "bear", "coffee")
	result := l.Element(strs...).Render()
	assert.Equalf(t, `<div><p>apple</p><p>bear</p><p>coffee</p></div>`, result, "should match")
}
