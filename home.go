package main

import (
	"html/template"
	"log"
	"net/http"
)

type player struct {
	fname string
	lname string
	// Men are true and Women are false
	gender bool
}

var players []player

func handleFilePath() {
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
}

func landingPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/home.html"))
	tmpl.ExecuteTemplate(w, "home.html", nil)
}

func addPlayerHandler(w http.ResponseWriter, r *http.Request) {
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")
	gender := r.FormValue("gender")

	var tempPlayer player

	// Set men to true and women to false
	if gender == "male" {
		tempPlayer = player{fname, lname, true}
	} else {
		tempPlayer = player{fname, lname, false}
	}

	players = append(players, tempPlayer)

	log.Println("Player Added" +
		"\nName: " + fname + " " + lname +
		"\nGender: " + gender)
}

func createGamesHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	handleFilePath()
	http.HandleFunc("/", landingPage)

	log.Println("Listening...")

	http.HandleFunc("/test", addPlayerHandler)
	http.HandleFunc("/createGames", createGamesHandler)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Panicln(err)
	}
}
