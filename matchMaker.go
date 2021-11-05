package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

type playerType struct {
	fname string
	lname string
	// Men are true and Women are false
	gender bool
}

var players []playerType

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

	_ = addPlayer(fname, lname, gender)
}

func addPlayer(fname string, lname string, gender string) bool {

	if fname == "" || lname == "" || gender == "" {
		return false
	}

	var playerToBeAdded playerType

	// Set men to true and women to false
	if strings.ToLower(gender) == "male" {
		playerToBeAdded = playerType{fname, lname, true}
	} else {
		playerToBeAdded = playerType{fname, lname, false}
	}

	players = append(players, playerToBeAdded)

	return true
}

func getPlayerCount(people []playerType) (int, int) {
	var malePlayers int
	var femalePlayers int

	for _, person := range people {
		if person.gender == true {
			malePlayers++
			continue
		}
		femalePlayers++
	}

	log.Printf("Male count: %v", malePlayers)
	log.Printf("Female count: %v", femalePlayers)
	return malePlayers, femalePlayers
}


func main() {
	handleFilePath()
	http.HandleFunc("/", landingPage)

	log.Println("Listening...")

	http.HandleFunc("/addPlayer", addPlayerHandler)

	http.HandleFunc("/createGames", createGamesHandler)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Panicln(err)
	}
}
