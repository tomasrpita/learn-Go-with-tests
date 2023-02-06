package blogposts_test

import (
	blogposts "learn-Go-with-tests/17-reading-files"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello-world.md":  {Data: []byte("Titlle:Post 1")},
		"hello-world2.md": {Data: []byte("Titlle:Post 2")},
	}

	posts, err := blogposts.NewPotsFromFs(fs)

	if err != nil {
		t.Fatal(err)
	}
	got := posts[0]
	want := blogposts.Post{Title: "Post 1"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v want %+v posts", got, want)
	}
}
