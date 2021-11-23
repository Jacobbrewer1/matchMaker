package main

import "testing"

func Test_getPlayerCount(t *testing.T) {
	tests := []struct {
		name                string
		players             []playerType
		expectedMaleCount   int
		expectedFemaleCount int
	}{
		{"Even Count", setupPlayersArrayTest(10), 5, 5},
		{"One third Female", setupThirdFemaleArrayTest(), 7, 3},
		{"All male", setupSingleGenderPlayersTest(true), 10, 0},
		{"All female", setupSingleGenderPlayersTest(false), 0, 10},
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
		{"1 Player", setupPlayersArrayTest(1), 0},
		{"3 Players", setupPlayersArrayTest(3), 3},
		{"5 Players", setupPlayersArrayTest(5), 10},
		{"7 Players", setupPlayersArrayTest(7), 21},
		{"10 Players", setupPlayersArrayTest(10), 45},
		{"11 Players", setupPlayersArrayTest(11), 55},
		{"13 Players", setupPlayersArrayTest(13), 78},
		{"17 Players", setupPlayersArrayTest(17), 136},
		{"19 Players", setupPlayersArrayTest(19), 171},
		{"21 Players", setupPlayersArrayTest(21), 210},
		{"57 Players", setupPlayersArrayTest(57), 1596},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInt := maxNumber(len(tt.players), pairsText, playersText)
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
		{"1 Pair", setupPlayersArrayTest(1), true},
		{"3 Pairs", setupPlayersArrayTest(3), false},
		{"4 Pairs", setupPlayersArrayTest(4), false},
		{"7 Pairs", setupPlayersArrayTest(7), false},
		{"10 Pairs", setupPlayersArrayTest(10), false},
		{"11 Pairs", setupPlayersArrayTest(11), false},
		{"17 Pairs", setupPlayersArrayTest(17), false},
		{"21 Pairs", setupPlayersArrayTest(21), false},
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
		{"1 Pair", setupPartnersArrayTest(1), true},
		{"3 Pairs", setupPartnersArrayTest(3), false},
		{"4 Pairs", setupPartnersArrayTest(4), false},
		{"7 Pairs", setupPartnersArrayTest(7), false},
		{"10 Pairs", setupPartnersArrayTest(10), false},
		{"11 Pairs", setupPartnersArrayTest(11), false},
		{"17 Pairs", setupPartnersArrayTest(17), false},
		{"21 Pairs", setupPartnersArrayTest(21), false},
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
		{"1 Player", setupPlayersArrayTest(1)},
		{"3 Players", setupPlayersArrayTest(3)},
		{"7 Players", setupPlayersArrayTest(7)},
		{"10 Players", setupPlayersArrayTest(10)},
		{"11 Players", setupPlayersArrayTest(11)},
		{"17 Players", setupPlayersArrayTest(17)},
		{"21 Players", setupPlayersArrayTest(21)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPairs := generateRandomPairs(tt.players)
			if checkForDuplicatesPairsTest(gotPairs) {
				t.Errorf("generateRandomPairs() = %v, expected %v", false, true)
			}
		})
	}
}

func Test_generateGames(t *testing.T) {
	tests := []struct {
		name         string
		players      []playerType
		expectedFail bool
	}{
		{"1 Player", setupPlayersArrayTest(1), true},
		{"3 Players", setupPlayersArrayTest(3), true},
		{"7 Players", setupPlayersArrayTest(7), false},
		{"10 Players", setupPlayersArrayTest(10), false},
		{"11 Players", setupPlayersArrayTest(11), false},
		{"17 Players", setupPlayersArrayTest(17), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGames, gotFailed := generateRandomGames(tt.players)
			if gotFailed != tt.expectedFail {
				t.Errorf("generateRandomPairs() = %v, expected %v", gotFailed, tt.expectedFail)
			}
			if checkForDuplicatesGamesDoublesTest(gotGames) {
				t.Errorf("generateRandomPairs() contains duplicates, expected no duplicates")
			}
		})
	}
}
