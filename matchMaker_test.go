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
		{"Even Count", setupPlayersArrayTest(10), setupGenderStringsTest(10, false), 5, true},
		{"All male", setupSingleGenderPlayersTest(true), setupGenderStringsTest(10, false), 5, true},
		{"All female", setupSingleGenderPlayersTest(false), setupGenderStringsTest(10, false), 5, true},
		{"Ability fail too low", setupPlayersArrayTest(10), setupGenderStringsTest(10, false), 0, false},
		{"Ability fail too high", setupPlayersArrayTest(10), setupGenderStringsTest(10, false), 21, false},
		{"Ability pass boundary", setupPlayersArrayTest(10), setupGenderStringsTest(10, false), 1, true},
		{"Fail players", setupPlayersArrayTest(10), setupGenderStringsTest(10, true), 5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.ability == 5 {
				for personId, person := range tt.players {
					got := addPlayer(person.fname, person.lname, tt.genders[personId], person.ability)
					if got != tt.expectedResult {
						t.Errorf("addPlayer() = %v, expected %v", got, tt.expectedResult)
					}
				}
			} else {
				person := tt.players[5]
				got := addPlayer(person.fname, person.lname, tt.genders[5], tt.ability)
				if got != tt.expectedResult {
					t.Errorf("addPlayer() = %v, expected %v", got, tt.expectedResult)
				}
			}
		})
	}

}
