package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type playerType struct {
	fname string
	lname string
	// Men are true and Women are false
	gender  bool
	ability int
}

var players []playerType

var templates *template.Template

func handleFilePath() {
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	templates = template.Must(template.New("").ParseGlob("./templates/*.html"))
}

func landingPage(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "home", nil)
	if err != nil {
		log.Panicln(err)
		return
	}
}

func addPlayerHandler(w http.ResponseWriter, r *http.Request) {
	fname := r.FormValue("firstName")
	lname := r.FormValue("lastName")
	gender := r.FormValue("gender")
	ability := r.FormValue("ability")

	atoi, err := strconv.Atoi(ability)
	if err != nil {
		log.Panicln(err)
		return
	}

	if addPlayer(fname, lname, gender, atoi) {
		log.Printf("Player added: %v, %v, %v", fname, lname, gender)

		pagePlayerData := struct {
			PlayerId int
			Fname    string
			Lname    string
			Gender   bool
			Ability  int
		}{
			len(players), fname, lname, false, atoi,
		}

		if gender == "male" {
			pagePlayerData.Gender = true
		}

		// return the last element added to slice before it was always returning the first element.
		err := templates.ExecuteTemplate(w, "displayPlayer", pagePlayerData)
		if err != nil {
			log.Panicln(err)
			return
		}
	}
}

func addPlayer(fname string, lname string, gender string, ability int) bool {

	if fname == "" || lname == "" || gender == "" || (ability < 1 || 10 < ability) {
		return false
	}

	var playerToBeAdded playerType

	// Set men to true and women to false
	if strings.ToLower(gender) == "male" {
		playerToBeAdded = playerType{fname, lname, true, ability}
	} else {
		playerToBeAdded = playerType{fname, lname, false, ability}
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
