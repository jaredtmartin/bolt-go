package bolt_test

import (
	"testing"

	"github.com/jaredtmartin/bolt-go"
	"github.com/stretchr/testify/assert"
)

func TestOptionStrings(t *testing.T) {
	options := []string{"red", "blue", "green"}
	result := bolt.OptionStrings(options)

	assert.Equal(t, 3, len(result))
	assert.Equal(t, "red", result[0].Label)
	assert.Equal(t, "red", result[0].Value)
	assert.Equal(t, "blue", result[1].Label)
	assert.Equal(t, "blue", result[1].Value)
	assert.Equal(t, "green", result[2].Label)
	assert.Equal(t, "green", result[2].Value)
}

func TestNewField(t *testing.T) {
	f := bolt.NewField("username", "Username", "admin", "text")

	assert.Equal(t, "div", f.GetTag())
	assert.Equal(t, "Username", f.Label.GetContent())
	assert.Equal(t, "username", f.Input.GetAttr("name"))
	assert.Equal(t, "admin", f.Input.GetAttr("value"))
	assert.Equal(t, "text", f.Input.GetAttr("type"))
	assert.Equal(t, "username-field", f.Input.GetAttr("id"))
	assert.Equal(t, "username-field-error", f.Error.GetAttr("id"))

	// Verify children order
	children := f.GetChildren()
	assert.Len(t, children, 3)
	assert.Equal(t, f.Label, children[0])
	assert.Equal(t, f.Input, children[1])
	assert.Equal(t, f.Error, children[2])
}

func TestHiddenField(t *testing.T) {
	f := bolt.HiddenField("csrf_token", "", "xyz123")

	// HiddenField sets tag to "" (fragment-like)
	assert.Equal(t, "", f.GetTag())
	assert.Equal(t, "hidden", f.Input.GetAttr("type"))
	assert.Equal(t, "csrf_token", f.Input.GetAttr("name"))
	assert.Equal(t, "xyz123", f.Input.GetAttr("value"))

	result := f.Render()
	assert.Equal(t, `<input name="csrf_token" type="hidden" value="xyz123">`, result)
}

func TestField_Required(t *testing.T) {
	f := bolt.TextField("name", "Name", "")

	// Test enabling required
	f.Required()
	assert.Equal(t, "required", f.Input.GetAttr("required"))
	assert.Equal(t, "doFieldValidation(event)", f.Input.GetAttr("onblur"))

	// Test disabling required
	f.Required(false)
	assert.Empty(t, f.Input.GetAttr("required"))
	assert.Empty(t, f.Input.GetAttr("onblur"))

	// Test explicitly enabling required
	f.Required(true)
	assert.Equal(t, "required", f.Input.GetAttr("required"))
}

func TestField_EmailAndPassword(t *testing.T) {
	f := bolt.TextField("input", "Input", "")

	// Test Email
	f.Email()
	assert.Equal(t, "email", f.Input.GetAttr("type"))
	f.Email(false)
	assert.Equal(t, "text", f.Input.GetAttr("type"))

	// Test Password
	f.Password()
	assert.Equal(t, "password", f.Input.GetAttr("type"))
	f.Password(false)
	assert.Equal(t, "text", f.Input.GetAttr("type"))
}

func TestField_Checked(t *testing.T) {
	f := bolt.Radio("choice", "Choice", "A")

	f.Checked()
	assert.Equal(t, "checked", f.Input.GetAttr("checked"))

	f.Checked(false)
	assert.Empty(t, f.Input.GetAttr("checked"))

	f.Checked(true)
	assert.Equal(t, "checked", f.Input.GetAttr("checked"))
}
