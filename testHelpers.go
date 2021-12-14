package main

import (
	"log"
	"math/rand"
	"strconv"
)

func setupPlayersArrayTest(numberOfPlayers int) []playerType {
	numberOfPlayers--
	playerCount := 0
	var testPlayers []playerType
	abilityTest := 1
	for {
		var tempPlayer playerType

		tempPlayer.fname = "Test"
		tempPlayer.lname = strconv.Itoa(playerCount)
		if playerCount%2 == 0 {
			tempPlayer.gender = true
		} else {
			tempPlayer.gender = false
		}

		tempPlayer.ability = abilityTest
		abilityTest++
		if abilityTest > 10 {
			abilityTest = 1
		}

		testPlayers = append(testPlayers, tempPlayer)
		playerCount++
		if playerCount > numberOfPlayers {
			break
		}
	}
	return testPlayers
}

func setupThirdFemaleArrayTest() []playerType {
	playerCount := 0
	var testPlayers []playerType
	abilityTest := 1
	for {
		var tempPlayer playerType

		tempPlayer.fname = "Test"
		tempPlayer.lname = strconv.Itoa(playerCount)
		if playerCount%4 == 0 {
			tempPlayer.gender = false
		} else {
			tempPlayer.gender = true
		}

		tempPlayer.ability = abilityTest
		abilityTest++
		if abilityTest > 10 {
			abilityTest = 1
		}

		testPlayers = append(testPlayers, tempPlayer)
		playerCount++
		if playerCount > 9 {
			break
		}
	}
	return testPlayers
}

func setupSingleGenderPlayersTest(gender bool) []playerType {
	playerCount := 0
	var testPlayers []playerType
	abilityTest := 1
	for {
		var tempPlayer playerType

		tempPlayer.fname = "Test"
		tempPlayer.lname = strconv.Itoa(playerCount)
		tempPlayer.gender = gender

		tempPlayer.ability = abilityTest
		abilityTest++
		if abilityTest > 10 {
			abilityTest = 1
		}

		testPlayers = append(testPlayers, tempPlayer)
		playerCount++
		if playerCount > 9 {
			break
		}
	}
	return testPlayers
}

func setupErrorPlayersTest() []playerType {
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

func setupGenderStringsTest(playerCount int, blank bool) []string {
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

func checkForDuplicatesPairsTest(pairings []partners) bool {
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

func checkForDuplicatesGamesDoublesTest(pairings []doublesFormat) bool {
	for pairId, pair := range pairings {
		for selectedPairId, selectedPair := range pairings {
			if pairId == selectedPairId {
				continue
			}
			if pair == selectedPair {
				log.Printf("Pair Id: %v Pair: %v", pairId, selectedPairId)
				return true
			}
			players := []playerType{
				{pair.pairOne.playerOne.fname, pair.pairOne.playerOne.lname, pair.pairOne.playerOne.gender, pair.pairOne.playerOne.ability},
				{pair.pairOne.playerTwo.fname, pair.pairOne.playerTwo.lname, pair.pairOne.playerTwo.gender, pair.pairOne.playerTwo.ability},
				{pair.pairTwo.playerOne.fname, pair.pairTwo.playerOne.lname, pair.pairTwo.playerOne.gender, pair.pairTwo.playerOne.ability},
				{pair.pairTwo.playerTwo.fname, pair.pairTwo.playerTwo.lname, pair.pairTwo.playerTwo.gender, pair.pairTwo.playerTwo.ability},
			}
			if checkGamePairTest(players) {
				log.Printf("Pair Id: %v Pair: %v", pairId, selectedPairId)
				return true
			}
		}
	}
	return false
}

func checkGamePairTest(players []playerType) bool {
	for playerId, player := range players {
		for checkPlayerId, selectedPlayer := range players {
			if playerId == checkPlayerId {
				continue
			}
			if player == selectedPlayer {
				return true
			}
		}
	}
	return false
}

func setupPartnersArrayTest(numberOfPairs int) []partners {
	playersArray := setupPlayersArrayTest(numberOfPairs * 2)
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

func setupPartnersTest(setup int) partners {
	var pair partners
	switch setup {
	case 0:
		p := partners{
			playerOne: playerType{"Steve", "Jobs", true, 10},
			playerTwo: playerType{"Nick", "Fury", true, 10},
		}
		pair = p
		break
	case 1:
		p := partners{
			playerOne: playerType{"Lewis", "Hamilton", true, 10},
			playerTwo: playerType{"Bill", "Gates", true, 10},
		}
		pair = p
		break
	case 2:
		p := partners{
			playerOne: playerType{"Amy", "Winehouse", false, 10},
			playerTwo: playerType{"Lucy", "Sparks", false, 10},
		}
		pair = p
		break
	case 3:
		p := partners{
			playerOne: playerType{"Lauren", "Hamilton", false, 10},
			playerTwo: playerType{"Penny", "Hofstedder", false, 10},
		}
		pair = p
		break
	case 4:
		p := partners{
			playerOne: playerType{"Lewis", "Hamilton", true, 10},
			playerTwo: playerType{"Lucy", "Rosewarn", false, 10},
		}
		pair = p
		break
	case 5:
		p := partners{
			playerOne: playerType{"Jacob", "Tucker", true, 10},
			playerTwo: playerType{"Tash", "Lado", false, 10},
		}
		pair = p
		break
	}

	return pair
}
