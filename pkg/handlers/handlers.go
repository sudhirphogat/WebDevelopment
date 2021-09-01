package handlers

import (
	"net/http"

	"github.com/sudhirphogat/WebDevelopment/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")

}

//About is called when /about is called
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
