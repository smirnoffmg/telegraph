package telegraph

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

// Allowed tags and attributes
var allowedTags = map[string]struct{}{
	"a":          {},
	"aside":      {},
	"b":          {},
	"blockquote": {},
	"br":         {},
	"code":       {},
	"em":         {},
	"figcaption": {},
	"figure":     {},
	"h3":         {},
	"h4":         {},
	"hr":         {},
	"i":          {},
	"iframe":     {},
	"img":        {},
	"li":         {},
	"ol":         {},
	"p":          {},
	"pre":        {},
	"s":          {},
	"strong":     {},
	"u":          {},
	"ul":         {},
	"video":      {},
}

var allowedAttrs = map[string]struct{}{
	"href": {}, "src": {},
}

// domToNode converts an HTML node to a Node
func domToNode(n *html.Node) interface{} {
	if n.Type == html.TextNode {
		textContent := strings.TrimSpace(n.Data)
		if textContent != "" {
			return textContent
		}
		return nil
	}

	if n.Type != html.ElementNode {
		return nil
	}

	// Ensure the tag is allowed
	if _, ok := allowedTags[n.Data]; !ok {
		return nil
	}

	nodeElement := NodeElement{
		Tag:   n.Data,
		Attrs: make(map[string]string),
	}

	for _, attr := range n.Attr {
		if _, ok := allowedAttrs[attr.Key]; ok {
			nodeElement.Attrs[attr.Key] = attr.Val
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		child := domToNode(c)
		if child != nil {
			nodeElement.Children = append(nodeElement.Children, child)
		}
	}

	return nodeElement
}

// HTMLToContent transforms HTML string to a slice of Nodes
func HTMLToContent(htmlStr string) ([]Node, error) {
	doc, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		return nil, err
	}

	var bodyNode *html.Node
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "body" {
			bodyNode = n
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	if bodyNode == nil {
		return nil, fmt.Errorf("no body element found")
	}

	var content []Node
	for c := bodyNode.FirstChild; c != nil; c = c.NextSibling {
		node := domToNode(c)
		if node != nil {
			content = append(content, node)
		}
	}

	return content, nil
}
