package main

import (
	"github.com/NYTimes/gziphandler"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

type handleFunc func(http.ResponseWriter, *http.Request)

func logHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		h(w, r)
	}
}

func webHandlers(h handleFunc) http.Handler {
	return gziphandler.GzipHandler(logHandler(http.HandlerFunc(h)))
}

func handleFilePath() {
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
}

func landingPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/home.html"))
	tmpl.ExecuteTemplate(w, "home.html", nil)
}

func main() {
	handleFilePath()
	http.HandleFunc("/", landingPage)

	log.Println("Listening...")
	r := mux.NewRouter()

	r.Handle("/test", webHandlers(usage))

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Panicln(err)
	}
}

func usage(w http.ResponseWriter, r *http.Request) {
	log.Println("Test")
}
