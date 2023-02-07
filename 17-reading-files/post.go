package blogposts

import (
	"bufio"
	"io"
)

type Post struct {
	Title       string
	Description string
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readLine := func() string {
		scanner.Scan()
		return scanner.Text()
	}

	titleLine := readLine()

	descriptionLine := readLine()

	post := Post{Title: titleLine[len(titleSeparator):], Description: descriptionLine[len(descriptionSeparator):]}

	return post, nil
}
