package main

import "testing"

func Test_addPlayer(t *testing.T) {
	tests := []struct {
		name           string
		players        []playerType
		genders        []string
		expectedResult bool
	}{
		{"Even Count", setupPlayersArray(10), setupGenderStrings(10, false), true},
		{"All male", setupSingleGenderPlayers(true), setupGenderStrings(10, false), true},
		{"All female", setupSingleGenderPlayers(false), setupGenderStrings(10, false), true},
		{"Fail players", setupPlayersArray(10), setupGenderStrings(10, true), false},
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
