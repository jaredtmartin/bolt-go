package bolt

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

// Traverses the AST and returns an Element
func traverseAST(n *html.Node, depth int) Element {
	// Print indentation based on the node's depth in the tree
	fmt.Printf("%s", strings.Repeat("  ", depth))
	el := None()
	// Recurse through child nodes
	switch n.Type {
	case html.ElementNode:
		fmt.Printf("Element: <%s>\n", n.Data)
		el = NewElement(n.Data)
	case html.TextNode:
		// Trim space for cleaner output of text nodes
		content := strings.TrimSpace(n.Data)
		if content != "" {
			fmt.Printf("Text: \"%s\"\n", content)
			el = String(n.Data)
		}
	case html.CommentNode:
		fmt.Printf("Comment: <!--%s-->\n", n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		el.AddChild(traverseAST(c, depth+1))
	}
	return el
}

// Parses an HTML string into a bolt.Element
func FromHtml(s string) (Element, error) {
	// Parse html string into element
	r := strings.NewReader(s)

	// Parse the HTML content into the AST
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Printf("error parsing HTML: %v\n", err)
		return nil, err
	}
	return traverseAST(doc, 0), nil
}
