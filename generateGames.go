package main

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

const pairsText = "pair(s)"
const playersText = "player(s)"
const gamesText = "game(s)"

type singles struct {
	playerOne playerType
	playerTwo playerType
}

type partners struct {
	playerOne playerType
	playerTwo playerType
}

type doublesFormat struct {
	pairOne partners
	pairTwo partners
}

var pairings []partners
var generatedDoubles []doublesFormat

func setupPlayersArrayTest() {
	playerCount := 0
	for {
		var tempPlayer playerType

		tempPlayer.fname = "Test"
		tempPlayer.lname = strconv.Itoa(playerCount)
		if playerCount%2 == 0 {
			tempPlayer.gender = true
		} else {
			tempPlayer.gender = false
		}

		players = append(players, tempPlayer)
		playerCount++
		if playerCount > 9 {
			break
		}
	}
}

func createGamesHandler(w http.ResponseWriter, r *http.Request) {
	// maleCount, femaleCount := getPlayerCount(players)

	setupPlayersArrayTest()

	randomGames, fail := generateRandomGames(players)
	if fail {
		log.Panicln("No games generated. Stopping process now")
		return
	}
	log.Println("Games generated successfully")

	for gameId, game := range randomGames {
		pageGameData := struct {
			GameId  int
			Player1 string
			Player2 string
			Player3 string
			Player4 string
		}{
			gameId + 1, game.pairOne.playerOne.fname, game.pairOne.playerTwo.fname, game.pairTwo.playerOne.fname, game.pairTwo.playerTwo.fname,
		}

		// return the last element added to slice before it was always returning the first element.
		err := templates.ExecuteTemplate(w, "gamesDisplayTmpl", pageGameData)
		if err != nil {
			log.Panicln(err)
			return
		}
		log.Println("Template gamesDisplayTmpl executed successfully")
	}
}

func generateRandomGames(players []playerType) ([]doublesFormat, bool) {
	if len(players) < 4 {
		log.Println("Not enough played added, no games generated")
		return nil, true
	}

	pairs := generateRandomPairs(players)

	pairsTotal := len(pairs)
	if pairsTotal < 2 {
		log.Println("Not enough played added, no games generated")
		return nil, false
	}

	maxGames := maxNumber(pairsTotal, gamesText, pairsText)

	var tempDubsArray []doublesFormat
	var tempDubs doublesFormat
	var p1 partners
	var p2 partners

	gameCount := 0

	var check bool

	for {
		for {
			fail := false
			p1, p2, fail = getRandomPair(pairsTotal, pairs)
			if fail {
				log.Println("Not enough played added, no games generated")
				return nil, true
			}
			if p1 != p2 {
				break
			}
		}

		tempDubs.pairOne = p1
		tempDubs.pairTwo = p2

		check = false
		for _, dub := range tempDubsArray {
			if dub == tempDubs {
				check = true
				break
			}
		}

		if check {
			continue
		}

		tempDubsArray = append(tempDubsArray, tempDubs)

		gameCount++
		if gameCount == maxGames {
			break
		}
	}

	return tempDubsArray, false
}

func generateRandomPairs(people []playerType) []partners {
	playersTotal := len(people)
	if playersTotal < 2 {
		log.Println("Not enough played added, no games generated")
		return nil
	}

	maxPairs := maxNumber(playersTotal, pairsText, playersText)

	var tempPairArray []partners
	var tempPair partners
	var tempP1 playerType
	var tempP2 playerType

	pairCount := 0

	var check bool

	for {
		for {
			fail := false
			tempP1, tempP2, fail = getRandomPlayer(playersTotal, people)
			if fail {
				log.Println("Not enough played added, no games generated")
			}
			if tempP1 != tempP2 {
				break
			}
		}

		tempPair.playerOne = tempP1
		tempPair.playerTwo = tempP2

		check = false
		for _, pair := range tempPairArray {
			if pair == tempPair {
				check = true
				break
			}
		}

		if check {
			continue
		}

		tempPairArray = append(tempPairArray, tempPair)

		pairCount++
		if pairCount == maxPairs {
			break
		}
	}

	return tempPairArray
}

func maxNumber(maxNumber int, textOne string, textTwo string) int {
	max := maxNumber * (maxNumber - 1) / 2
	log.Printf("Max number of %v for %v %v: %v", textOne, maxNumber, textTwo, max)
	return max
}

func getRandomPlayer(max int, people []playerType) (playerType, playerType, bool) {
	if max < 2 {
		return playerType{
				fname:  "",
				lname:  "",
				gender: true,
			},
			playerType{
				fname:  "",
				lname:  "",
				gender: false,
			},
			true
	}

	n1 := rand.Intn(max)
	n2 := rand.Intn(max)
	var p1 playerType
	var p2 playerType

	if n1 < n2 {
		p1 = people[n1]
		p2 = people[n2]
	} else {
		p1 = people[n2]
		p2 = people[n1]
	}

	return p1, p2, false
}

func getRandomPair(max int, pairs []partners) (partners, partners, bool) {
	if max < 2 {
		return partners{
				playerOne: playerType{"", "", true, 0},
				playerTwo: playerType{"", "", false, 0},
			},
			partners{
				playerOne: playerType{"", "", true, 0},
				playerTwo: playerType{"", "", false, 0},
			},
			true
	}

	n1 := rand.Intn(max)
	n2 := rand.Intn(max)
	var p1 partners
	var p2 partners

	if n1 < n2 {
		p1 = pairs[n1]
		p2 = pairs[n2]
	} else {
		p1 = pairs[n2]
		p2 = pairs[n1]
	}

	return p1, p2, false
}
