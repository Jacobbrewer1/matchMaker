package main

import (
	"strconv"
	"testing"
)

func setupPlayersArray() []playerType {
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
		if playerCount > 9 {
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

func Test_addPlayer(t *testing.T) {
	tests := []struct {
		name           string
		players        []playerType
		genders        []string
		expectedResult bool
	}{
		{"Even Count", setupPlayersArray(), setupGenderStrings(10, false), true},
		{"All male", setupSingleGenderPlayers(true), setupGenderStrings(10, false), true},
		{"All female", setupSingleGenderPlayers(false), setupGenderStrings(10, false), true},
		{"Fail players", setupPlayersArray(), setupGenderStrings(10, true), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for personId, person := range tt.players {
				got := addPlayer(person.fname, person.lname, tt.genders[personId])
				if got != tt.expectedResult {
					t.Errorf("getPlayerCount() = %v, expected %v", got, tt.expectedResult)
				}
			}
		})
	}

}

func Test_getPlayerCount(t *testing.T) {
	tests := []struct {
		name                string
		players             []playerType
		expectedMaleCount   int
		expectedFemaleCount int
	}{
		{"Even Count", setupPlayersArray(), 5, 5},
		{"One third Female", setupThirdFemaleArray(), 7, 3},
		{"All male", setupSingleGenderPlayers(true), 10, 0},
		{"All female", setupSingleGenderPlayers(false), 0, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMale, gotFemale := getPlayerCount(tt.players)
			if gotMale != tt.expectedMaleCount {
				t.Errorf("getPlayerCount() = %v, expected %v", gotMale, tt.expectedMaleCount)
			}
			if gotFemale != tt.expectedFemaleCount {
				t.Errorf("got %v, expected %v", gotFemale, tt.expectedFemaleCount)
			}
		})
	}
}
