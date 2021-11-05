package main

import "testing"

func Test_getPlayerCount(t *testing.T) {
	tests := []struct {
		name                string
		players             []playerType
		expectedMaleCount   int
		expectedFemaleCount int
	}{
		{"Even Count", setupPlayersArray(10), 5, 5},
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

func Test_maxNumber(t *testing.T) {
	tests := []struct {
		name        string
		players     []playerType
		expectedMax int
	}{
		{"1 Player", setupPlayersArray(1), 0},
		{"3 Players", setupPlayersArray(3), 3},
		{"5 Players", setupPlayersArray(5), 10},
		{"7 Players", setupPlayersArray(7), 21},
		{"10 Players", setupPlayersArray(10), 45},
		{"11 Players", setupPlayersArray(11), 55},
		{"13 Players", setupPlayersArray(13), 78},
		{"17 Players", setupPlayersArray(17), 136},
		{"19 Players", setupPlayersArray(19), 171},
		{"21 Players", setupPlayersArray(21), 210},
		{"57 Players", setupPlayersArray(57), 1596},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInt := maxNumber(len(tt.players))
			if gotInt != tt.expectedMax {
				t.Errorf("maxNumber() = %v, expected %v", gotInt, tt.expectedMax)
			}
		})
	}
}

func Test_getPlayers(t *testing.T) {
	tests := []struct {
		name         string
		players      []playerType
		expectedFail bool
	}{
		{"1 Pair", setupPlayersArray(1), true},
		{"3 Pairs", setupPlayersArray(3), false},
		{"4 Pairs", setupPlayersArray(4), false},
		{"7 Pairs", setupPlayersArray(7), false},
		{"10 Pairs", setupPlayersArray(10), false},
		{"11 Pairs", setupPlayersArray(11), false},
		{"17 Pairs", setupPlayersArray(17), false},
		{"21 Pairs", setupPlayersArray(21), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, fail := getRandomPlayer(len(tt.players), tt.players)
			if fail != tt.expectedFail {
				t.Errorf("getPlayers() = %v, expected %v", fail, tt.expectedFail)
			}
		})
	}
}

func Test_getPairs(t *testing.T) {
	tests := []struct {
		name         string
		players      []partners
		expectedFail bool
	}{
		{"1 Pair", setupPartnersArray(1), true},
		{"3 Pairs", setupPartnersArray(3), false},
		{"4 Pairs", setupPartnersArray(4), false},
		{"7 Pairs", setupPartnersArray(7), false},
		{"10 Pairs", setupPartnersArray(10), false},
		{"11 Pairs", setupPartnersArray(11), false},
		{"17 Pairs", setupPartnersArray(17), false},
		{"21 Pairs", setupPartnersArray(21), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, fail := getRandomPair(len(tt.players), tt.players)
			if fail != tt.expectedFail {
				t.Errorf("getPlayers() = %v, expected %v", fail, tt.expectedFail)
			}
		})
	}
}

func Test_generatePairs(t *testing.T) {
	tests := []struct {
		name    string
		players []playerType
	}{
		{"1 Player", setupPlayersArray(1)},
		{"3 Players", setupPlayersArray(3)},
		{"7 Players", setupPlayersArray(7)},
		{"10 Players", setupPlayersArray(10)},
		{"11 Players", setupPlayersArray(11)},
		{"17 Players", setupPlayersArray(17)},
		{"21 Players", setupPlayersArray(21)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPairs := generateRandomPairs(tt.players)
			if checkForDuplicatesPairs(gotPairs) {
				t.Errorf("generateRandomPairs() = %v, expected %v", false, true)
			}
		})
	}
}

func Test_generateGames(t *testing.T) {
	tests := []struct {
		name    string
		players []playerType
	}{
		{"1 Player", setupPlayersArray(1)},
		{"3 Players", setupPlayersArray(3)},
		{"7 Players", setupPlayersArray(7)},
		{"10 Players", setupPlayersArray(10)},
		{"11 Players", setupPlayersArray(11)},
		{"17 Players", setupPlayersArray(17)},
		{"21 Players", setupPlayersArray(21)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGames := generateRandomGames(tt.players)
			if checkForDuplicatesGamesDoubles(gotGames) {
				t.Errorf("generateRandomPairs() = %v, expected %v", false, true)
			}
		})
	}
}
