package blogrender_test

import (
	"bytes"
	"github.com/approvals/go-approval-tests"
	blogrender "learn-Go-with-tests/18-templating"
	"testing"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogrender.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := blogrender.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}
