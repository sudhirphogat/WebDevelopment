package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/sudhirphogat/WebDevelopment/pkg/config"
)

var functions = template.FuncMap{}

var app *config.AppConfig

//sets config to template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	//We will get the cache from the app
	tc := app.TemplateCache

	//*** This is not required as I changed the nanem of RenderTemplateTest to CreateTemplateCache
	//we need to handle it differently as we want to show the data in portal
	//Currently using this code we only check that func is working but no data in the frontend
	//_, err := RenderTemplateTest(w)
	//if err != nil {
	//	fmt.Println("error getting template cache", err)
	//}

	//****Cahche is not a better way as this is rendering all the pages and storing it
	// this might impact once the traffic is high or n number of pages
	//tc, err := CreateTemplateCache()
	//if err != nil {
	//	log.Fatal(err)
	//}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Template cannot be created from template cache")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template", err)
	}

	//*** below code is good in case we do not have a layout template as it direcly parses the template
	//parshedTemplate, _ := template.ParseFiles("./templates/" + tmpl)

	//err = parshedTemplate.Execute(w, nil)
	//if err != nil {
	//	fmt.Println("error parsing template", err)
	//	return
	//}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	mycache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return mycache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		//fmt.Println("page is currently", page)

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
