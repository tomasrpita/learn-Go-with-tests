package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

// Post represents a post on a blog
type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

// NewPost creates a new post from a file
func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(preFix string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), preFix)
	}

	return Post{
		Title:       readMetaLine(titleSeparator),
		Description: readMetaLine(descriptionSeparator),
		Tags:        strings.Split(readMetaLine(tagsSeparator), ", "),
		Body:        readBody(scanner),
	}, nil

}

// readBody reads the body of the post
func readBody(scanner *bufio.Scanner) string {
	scanner.Scan() // ignore a line
	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	// Remove last \n
	return strings.TrimSuffix(buf.String(), "\n")
}
