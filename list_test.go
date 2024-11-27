package bolt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestNewList(t *testing.T) {
	l := List([]string{"apple", "bear", "coffee"})
	result := l.Render()
	assert.Equalf(t, `<div><p>apple</p><p>bear</p><p>coffee</p></div>`, result, "should match")
}
func TestNewListWithRenderRow(t *testing.T) {
	renderRow := func(item string) Element {
		return Div().Text(fmt.Sprintf("%v", item))
	}
	l := List([]string{"apple", "bear", "coffee"}, renderRow)
	result := l.Render()
	assert.Equalf(t, `<div><div>apple</div><div>bear</div><div>coffee</div></div>`, result, "should match")
}
