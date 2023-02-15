package blogrender

import (
	"embed"
	"html/template"
	"io"
	"strings"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

// type PostViewModel struct {
// 	Title, SanitisedTitle, Body, Description string
// 	Tags                                     []string
// }

func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {

	// if err := r.templ.Execute(w, p); err != nil {
	// 	return err
	// }

	// return nil
	return r.templ.ExecuteTemplate(w, "blog.gohtml", p)
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	// indexTemplate := `<ol>{{range .}}<li><a href="/post/{{.SanitisedTitle}}">{{.Title}}</a></li>{{end}}</ol>`

	// templ, err := template.New("index").Parse(indexTemplate)
	// if err != nil {
	// 	return err
	// }

	// if err := templ.Execute(w, posts); err != nil {
	// 	return err
	// }

	// return nil
	return r.templ.ExecuteTemplate(w, "index.gohtml", posts)
}
