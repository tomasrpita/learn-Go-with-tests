package blogrender_test

import (
	"bytes"
	"io"
	blogrender "learn-Go-with-tests/18-templating"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

var (
	aPost = blogrender.Post{
		Title:       "hello world",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}
)

func TestRender(t *testing.T) {

	postRederer, err := blogrender.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := postRederer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}
		approvals.VerifyString(t, buf.String())
	})

	t.Run("it renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blogrender.Post{
			{Title: "Hello World"},
			{Title: "Hello World 2"},
		}

		if err := postRederer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}
		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkXxx(b *testing.B) {
	postRederer, err := blogrender.NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postRederer.Render(io.Discard, aPost)
	}
}
