package telegraph_test

import (
	"encoding/json"
	"testing"

	"github.com/smirnoffmg/telegraph"
)

const testContentJSON = `[{"tag":"p","children":["Hello, ",{"tag":"b","children":["world"]},"! This is an ",{"tag":"a","attrs":{"href":"https://example.com"},"children":["example link"]},"."]}]`

func TestHTMLToContent(t *testing.T) {
	htmlStr := `<p>Hello, <b>world</b>! This is an <a href="https://example.com">example link</a>.</p>`
	content, err := telegraph.HTMLToContent(htmlStr)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(content) != 1 {
		t.Fatalf("Expected 1 Node, got %d", len(content))
	}

	contentJSON, err := json.Marshal(content)

	if err != nil {
		t.Fatalf("Failed to marshal content: %v", err)
	}

	if string(contentJSON) != testContentJSON {
		t.Errorf("Expected\n%s\ngot\n%s", testContentJSON, string(contentJSON))
	}
}
