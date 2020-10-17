package main

import (
	"html/template"
	"net/http"
	"regexp"
)

var templates = template.Must(template.ParseFiles("templates/page.html", "templates/header.html", "templates/index.html", "templates/edit.html", "templates/view.html", "templates/footer.html"))
var validPath = regexp.MustCompile("^/(index|edit|save|view)/([a-zA-Z0-9]+)$")

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	var templates = template.Must(template.ParseFiles("templates/page.html", "templates/header.html", "templates/index.html", "templates/edit.html", "templates/view.html", "templates/footer.html"))

	tplVars := map[string]interface{}{
		"Body":      template.HTML(p.Body),
		"Title":     p.Title,
		"TimeStamp": p.TimeStamp,
	}

	err := templates.ExecuteTemplate(w, tmpl+".html", tplVars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	getPages()
	var pages []*Page
	for _, v := range m {
		pages = append(pages, v)
	}
	err := templates.ExecuteTemplate(w, "index.html", pages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// loadPage loads pages from html files.
func loadPage(title string) (*Page, error) {
	getPages()

	if _, ok := m[title]; ok {
		return m[title], nil
	}
	// filename := "pages/" + title + ".html"
	// body, err := ioutil.ReadFile(filename)
	// if err != nil {
	// 	return nil, err
	// }
	return &Page{Title: title, Body: ""}, nil
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: body}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}
