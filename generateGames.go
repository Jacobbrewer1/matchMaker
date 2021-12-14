package main

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

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

func setupPlayersArray() {
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
	setupPlayersArray() // TODO : Remove this when ready for release

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

	maxGames, _ := maxNumber(len(players))

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
			playerCheck := []playerType{
				{p1.playerOne.fname, p1.playerOne.lname, p1.playerOne.gender, p1.playerOne.ability},
				{p1.playerTwo.fname, p1.playerTwo.lname, p1.playerTwo.gender, p1.playerTwo.ability},
				{p2.playerOne.fname, p2.playerOne.lname, p2.playerOne.gender, p2.playerOne.ability},
				{p2.playerTwo.fname, p2.playerTwo.lname, p2.playerTwo.gender, p2.playerTwo.ability},
			}
			if checkPair(playerCheck) || formalityChecker(p1, p2) {
				continue
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

// Returns true if the game is legitimate
func formalityChecker(pairOne partners, pairTwo partners) bool {
	playerArray := []playerType{
		{pairOne.playerOne.fname, pairOne.playerOne.lname, pairOne.playerOne.gender, pairOne.playerOne.ability},
		{pairOne.playerTwo.fname, pairOne.playerTwo.lname, pairOne.playerTwo.gender, pairOne.playerTwo.ability},
		{pairTwo.playerOne.fname, pairTwo.playerOne.lname, pairTwo.playerOne.gender, pairTwo.playerOne.ability},
		{pairTwo.playerTwo.fname, pairTwo.playerTwo.lname, pairTwo.playerTwo.gender, pairTwo.playerTwo.ability},
	}

	male, female := getPlayerCount(playerArray)
	if (female != male) && (male != 0 && female != 0) {
		return false
	}

	switch getGameType(pairOne, pairTwo) {
	case doublesGameText, mixedGameText:
		log.Println("Game is legitimate")
		return true
	default:
		log.Println("Game is not legitimate")
		return false
	}
}

func getGameType(pairOne partners, pairTwo partners) string {
	if pairOne.playerOne.gender && pairOne.playerTwo.gender && pairTwo.playerOne.gender && pairTwo.playerTwo.gender ||
		(!pairOne.playerOne.gender && !pairOne.playerTwo.gender && !pairTwo.playerOne.gender && !pairTwo.playerTwo.gender) {
		return doublesGameText
	} else if pairOne.playerOne.gender && !pairOne.playerTwo.gender && pairTwo.playerOne.gender && !pairTwo.playerTwo.gender ||
		(!pairOne.playerOne.gender && pairOne.playerTwo.gender && !pairTwo.playerOne.gender && pairTwo.playerTwo.gender) {
		return mixedGameText
	} else if !pairOne.playerOne.gender && pairOne.playerTwo.gender && !pairTwo.playerOne.gender && pairTwo.playerTwo.gender ||
		(pairOne.playerOne.gender && !pairOne.playerTwo.gender && pairTwo.playerOne.gender && !pairTwo.playerTwo.gender) {
		return mixedGameText
	}
	return ""
}

func checkPair(dubsPlayers []playerType) bool {
	for playerId, person := range dubsPlayers {
		for checkPlayerId, selectedPlayer := range dubsPlayers {
			if playerId == checkPlayerId {
				continue
			}
			if person == selectedPlayer {
				log.Printf("Duplicate Player found: %v", person)
				return true
			}
		}
	}
	return false
}

func generateRandomPairs(people []playerType) []partners {
	playersTotal := len(people)
	if playersTotal < 2 {
		log.Println("Not enough played added, no games generated")
		return nil
	}

	_, maxPairs := maxNumber(playersTotal)
	if maxPairs == 0 {
		log.Println("How enough players to generate doubles games")
		return nil
	}

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

func maxNumber(number int) (int, int) {
	pairs := 0
	games := 0
	if number < 2 {
		log.Println("Not enough players added to generate pairs")
	} else {
		pairs = number * (number - 1) / 2
		log.Printf("Max number of %v for %v %v: %v", pairsText, number, playersText, number)
	}

	if number < 4 {
		log.Printf("Not enough pairs added to generate %v", gamesText)
		number = 0
	} else {
		games = number / 2 // TODO : This equation needs to be corrected before the release
		log.Printf("Max number of %v for %v %v: %v", gamesText, number+1, pairsText, number)
	}
	return games, pairs
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
