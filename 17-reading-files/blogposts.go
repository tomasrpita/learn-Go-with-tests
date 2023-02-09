package blogposts

import (
	"io/fs"
)

// NewPostsFromFS returns a collection of blog posts from a file system. If it does not conform to the format then it'll return an error
func NewPostsFromFs(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post

	for _, file := range dir {
		post, err := getPost(fileSystem, file.Name())
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}
	return posts, nil
}

// getPost returns a post from a file
func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)

}
