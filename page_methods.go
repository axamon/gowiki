package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

func (p *Page) save() error {

	getPages()

	p.TimeStamp = time.Now().Format("2006-01-02T15:04")

	m[p.Title] = p

	savePages()

	filename := "pages/" + p.Title + ".html"
	f, err := os.Create(filename)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	tplVars := map[string]interface{}{
		"Body":      template.HTML(p.Body),
		"Title":     p.Title,
		"TimeStamp": p.TimeStamp,
	}

	err = templates.ExecuteTemplate(f, "page.html", tplVars)
	if err != nil {
		log.Println(err)
	}
	return err
	// return ioutil.WriteFile(filename, []byte(p.Body), 0600)
}
