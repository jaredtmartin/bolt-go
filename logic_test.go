package bolt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIf(t *testing.T) {
	l := If(true, P("1"), P("2"), P("3"))
	result := l.Render()
	assert.Equalf(t, `<p>1</p><p>2</p><p>3</p>`, result, "should match")
	l = If(false, P("false"), P("false"))
	result = l.Render()
	assert.Equalf(t, ``, result, "should match")
}
func TestIfElse(t *testing.T) {
	l := IfElse(true, P("true"), P("false"), P("other"))
	result := l.Render()
	assert.Equalf(t, `<p>true</p>`, result, "should match")
	l = IfElse(false, P("true"), P("false"), P("other"))
	result = l.Render()
	assert.Equalf(t, `<p>false</p><p>other</p>`, result, "should match")
}
func TestFor(t *testing.T) {
	l := For([]string{"apple", "bear", "coffee"}, func(item string, idx int) Element {
		return P(item)
	})
	result := l.Render()
	assert.Equalf(t, `<div><p>apple</p><p>bear</p><p>coffee</p></div>`, result, "should match")
}
func TestSwitchCase(t *testing.T) {
	l := Switch("apple",
		Case("apple", P("apple")),
		Case("pear", P("pear")),
	)
	result := l.Render()
	assert.Equalf(t, `<p>apple</p>`, result, "should match")
	l = Switch("pear",
		Case("apple", P("apple")),
		Case("pear", P("pear")),
	)
	result = l.Render()
	assert.Equalf(t, `<p>pear</p>`, result, "should match")
	l = Switch("orange",
		Case("apple", P("apple")),
		Case("pear", P("pear")),
		Default(P("other")),
	)
	result = l.Render()
	assert.Equalf(t, `<p>other</p>`, result, "should match")
}
