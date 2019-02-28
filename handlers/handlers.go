package handlers

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

// HomeHandler / route handler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	homeVars := struct {
		Title string
	}{
		Title: "Home",
	}

	templateLayout, err := ioutil.ReadFile("../templates/layout.html")
	if err != nil {
		log.Fatal(err)
	}
	templateHome, err := ioutil.ReadFile("../templates/home.html")
	if err != nil {
		log.Fatal(err)
	}
	t := template.New("")
	t.Parse(string(templateLayout))
	t.Parse(string(templateHome))
	err = t.ExecuteTemplate(w, "layout", homeVars)
}

// AboutHandler /about route handler
func AboutHandler(w http.ResponseWriter, r *http.Request) {
}
