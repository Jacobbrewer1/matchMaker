package main

import (
	"log"
	"math/rand"
	"strconv"
)

func setupPlayersArray(numberOfPlayers int) []playerType {
	numberOfPlayers--
	playerCount := 0
	var testPlayers []playerType
	for {
		var tempPlayer playerType

		tempPlayer.fname = "Test"
		tempPlayer.lname = strconv.Itoa(playerCount)
		if playerCount%2 == 0 {
			tempPlayer.gender = true
		} else {
			tempPlayer.gender = false
		}

		testPlayers = append(testPlayers, tempPlayer)
		playerCount++
		if playerCount > numberOfPlayers {
			break
		}
	}
	return testPlayers
}

func setupThirdFemaleArray() []playerType {
	playerCount := 0
	var testPlayers []playerType
	for {
		var tempPlayer playerType

		tempPlayer.fname = "Test"
		tempPlayer.lname = strconv.Itoa(playerCount)
		if playerCount%4 == 0 {
			tempPlayer.gender = false
		} else {
			tempPlayer.gender = true
		}

		testPlayers = append(testPlayers, tempPlayer)
		playerCount++
		if playerCount > 9 {
			break
		}
	}
	return testPlayers
}

func setupSingleGenderPlayers(gender bool) []playerType {
	playerCount := 0
	var testPlayers []playerType
	for {
		var tempPlayer playerType

		tempPlayer.fname = "Test"
		tempPlayer.lname = strconv.Itoa(playerCount)
		tempPlayer.gender = gender

		testPlayers = append(testPlayers, tempPlayer)
		playerCount++
		if playerCount > 9 {
			break
		}
	}
	return testPlayers
}

func setupErrorPlayers() []playerType {
	playerCount := 0
	var testPlayers []playerType
	for {
		var tempPlayer playerType

		tempPlayer.fname = "Test"
		tempPlayer.lname = strconv.Itoa(playerCount)

		testPlayers = append(testPlayers, tempPlayer)
		playerCount++
		if playerCount > 9 {
			break
		}
	}
	return testPlayers
}

func setupGenderStrings(playerCount int, blank bool) []string {
	var genders []string
	localCount := 0
	for {
		if blank && localCount < playerCount {
			genders = append(genders, "")
			localCount++
			continue
		}
		if localCount%2 == 0 {
			genders = append(genders, "Male")
		} else {
			genders = append(genders, "Female")
		}

		localCount++
		if localCount > playerCount {
			break
		}
	}
	return genders
}

func checkForDuplicatesPairs(pairings []partners) bool {
	for pairId, pair := range pairings {
		for selectedPairId, selectedPair := range pairings {
			if pairId == selectedPairId {
				continue
			}
			if pair == selectedPair {
				log.Printf("Pair Id: %v Pair: %v", pairId, selectedPairId)
				return false
			}
		}
	}
	return false
}

func checkForDuplicatesGamesDoubles(pairings []doublesFormat) bool {
	for pairId, pair := range pairings {
		for selectedPairId, selectedPair := range pairings {
			if pairId == selectedPairId {
				continue
			}
			if pair == selectedPair {
				log.Printf("Pair Id: %v Pair: %v", pairId, selectedPairId)
				return false
			}
		}
	}
	return false
}

func setupPartnersArray(numberOfPairs int) []partners {
	playersArray := setupPlayersArray(numberOfPairs * 2)
	max := len(playersArray)

	pairCount := 0

	var testPairArray []partners

	for {
		var tempPair partners

		var p1 playerType
		var p2 playerType

		for {
			n1 := rand.Intn(max)
			n2 := rand.Intn(max)

			if n1 < n2 {
				p1 = playersArray[n1]
				p2 = playersArray[n2]
			} else {
				p1 = playersArray[n2]
				p2 = playersArray[n1]
			}

			if p1 != p2 {
				break
			}
		}

		tempPair.playerOne = p1
		tempPair.playerTwo = p2

		check := false
		for _, pair := range testPairArray {
			if pair == tempPair {
				check = true
				break
			}
		}

		if check {
			continue
		}

		testPairArray = append(testPairArray, tempPair)

		pairCount++
		if pairCount == numberOfPairs {
			break
		}

		testPairArray = append(testPairArray, tempPair)
		pairCount++
		if pairCount > numberOfPairs {
			break
		}
	}
	return testPairArray
}
