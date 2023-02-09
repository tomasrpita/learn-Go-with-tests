package blogrender

import "io"

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

func Render(w io.Writer, p Post) error {
	return nil
}
