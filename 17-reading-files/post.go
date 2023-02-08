package blogposts

import (
	"bufio"
	"io"
	"strings"
)

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

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(preFix string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), preFix)
	}

	titleLine := readMetaLine(titleSeparator)

	descriptionLine := readMetaLine(descriptionSeparator)

	tagsline := strings.Split(readMetaLine(tagsSeparator), ", ")

	post := Post{Title: titleLine, Description: descriptionLine, Tags: tagsline}

	return post, nil
}
