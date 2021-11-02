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
