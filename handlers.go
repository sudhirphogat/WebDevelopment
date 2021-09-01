package main

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.html")

}

//About is called when /about is called
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.html")
}
