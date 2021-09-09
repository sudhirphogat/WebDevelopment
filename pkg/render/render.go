package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	_, err := RenderTemplateTest(w)
	if err != nil {
		fmt.Println("error getting template cache", err)
	}

	parshedTemplate, _ := template.ParseFiles("./templates/" + tmpl)

	err = parshedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}

}

func RenderTemplateTest(w http.ResponseWriter) (map[string]*template.Template, error) {
	mycache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return mycache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("page is currently", page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return mycache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return mycache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return mycache, err
			}
		}
		mycache[name] = ts

	}
	return mycache, nil

}
