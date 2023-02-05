package blogposts

import (
	"io/fs"
	"testing/fstest"
)

type Post struct{}

func NewPotsFromFs(fileSystem fstest.MapFS) []Post {
	dir, _ := fs.ReadDir(fileSystem, ".")
	var posts []Post
	for range dir {
		posts = append(posts, Post{})
	}
	return posts
}
