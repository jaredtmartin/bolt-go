package bolt_test

import (
	"testing"

	svg "github.com/jaredtmartin/bolt-go/svg"
	"github.com/stretchr/testify/assert"
)

func TestGroup(t *testing.T) {
	group := svg.Group(
		svg.Circle(0, 0, 10),
		svg.Circle(10, 10, 10),
		svg.Circle(20, 20, 10),
	)
	result := group.Render()
	expected := `<g><circle cx="0" cy="0" r="10"></circle><circle cx="10" cy="10" r="10"></circle><circle cx="20" cy="20" r="10"></circle></g>`
	assert.Equalf(t, expected, result, "should match")
}
