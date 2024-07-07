package bolt

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestNewSelect(t *testing.T) {
	options := []Option{
		{Label: "Male", Value: "male"},
		{Label: "Female", Value: "female"},
	}
	e := Select("gender", "Gender", options).Value("male")
	result := e.Render()
	log.Printf("result: %v\n", result)

	assert.Equalf(t, `<div><label for="gender-field">Gender</label><select id="gender-field" name="gender"><option selected="" value="male">Male</option><option value="female">Female</option></select><div id="gender-field-error"></div></div>`, result, "should match")
}
