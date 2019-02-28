package main

import (
	"fmt"
	"golang-packr-demo-fail/handlers"
	"log"
	"net/http"
	"os"

	"github.com/gobuffalo/packr/v2"
	"github.com/gorilla/mux"
)

var staticBox = packr.New("Static", "./static")

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "ok")
	})
	router.HandleFunc("/", handlers.HomeHandler)
	router.HandleFunc("/about", handlers.AboutHandler)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(staticBox)))
	log.Println("Listening port " + os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}
