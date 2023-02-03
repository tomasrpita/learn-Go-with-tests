package blogposts_test

import (
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello-world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
	}

	posts := blogpots.NewPotsFromFs(fs)

	if len(posts) != len(fs) {
		t.Errorf("got %d want %d posts", len(posts), len(fs))
	}
}
