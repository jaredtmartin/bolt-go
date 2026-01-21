package bolt

import "testing"

func TestFromHtml(t *testing.T) {
	htmlContent := `
<!DOCTYPE html>
<html>
<head>
    <title>My Title</title>
</head>
<body>
    <h1>My First Heading</h1>
    <p>My first paragraph.</p>
    <!-- A comment -->
</body>
</html>`
	result, err := FromHtml(htmlContent)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result == nil {
		t.Errorf("Expected a non-nil result, got nil")
	}
	expected := `<html><head><title>My Title</title></head><body><h1>My First Heading</h1><p>My first paragraph.</p></body></html>`
	s := result.Render()
	if s != expected {
		t.Errorf("Expected %q, got %q", expected, s)
	}
}
