package main

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

const pairs_text = "pair(s)"
const players_text = "player(s)"
const games_text = "game(s)"

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

	_, _ = generateRandomGames(players)
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

	maxGames := maxNumber(pairsTotal, games_text, pairs_text)

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

	for gameId, dub := range tempDubsArray {
		log.Printf("Game %v: %v", gameId, dub)
	}

	return tempDubsArray, false
}

func generateRandomPairs(people []playerType) []partners {
	playersTotal := len(people)
	if playersTotal < 2 {
		log.Println("Not enough played added, no games generated")
		return nil
	}

	maxPairs := maxNumber(playersTotal, pairs_text, players_text)

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

	for pairId, pair := range tempPairArray {
		log.Printf("Pair %v: %v", pairId, pair)
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
				playerOne: playerType{"", "", true},
				playerTwo: playerType{"", "", false},
			},
			partners{
				playerOne: playerType{"", "", true},
				playerTwo: playerType{"", "", false},
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
