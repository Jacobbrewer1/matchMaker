package main

import "testing"

func Test_addPlayer(t *testing.T) {
	tests := []struct {
		name           string
		players        []playerType
		genders        []string
		ability        int
		expectedResult bool
	}{
		// The ability of the player is generated dynamically in the setupPlayersArray() method
		//If it is redundant in the test, leave it as 5 for a placeholder
		{"Even Count", setupPlayersArray(10), setupGenderStrings(10, false),5, true},
		{"All male", setupSingleGenderPlayers(true), setupGenderStrings(10, false),5, true},
		{"All female", setupSingleGenderPlayers(false), setupGenderStrings(10, false),5, true},
		{"Ability fail too low", setupPlayersArray(10), setupGenderStrings(10,false), 0, true},
		{"Ability fail too high", setupPlayersArray(10), setupGenderStrings(10,false), 21, true},
		{"Ability pass boundary", setupPlayersArray(10), setupGenderStrings(10,false), 1, true},
		{"Fail players", setupPlayersArray(10), setupGenderStrings(10, true), 5,false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for personId, person := range tt.players {
				got := addPlayer(person.fname, person.lname, tt.genders[personId], person.ability)
				if got != tt.expectedResult {
					t.Errorf("addPlayer() = %v, expected %v", got, tt.expectedResult)
				}
			}
		})
	}

}
