// /repl_test.go
package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
    
cases := []struct {
	name	 string
	input    string
	expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		// add more cases here
		{
			name:     "single pokemon",
			input:    "Pikachu",
			expected: []string{"pikachu"},
		},
		{
			name:     "multiple pokemon",
			input:    "Pikachu Charizard Bulbasaur",
			expected: []string{"pikachu", "charizard", "bulbasaur"},
		},
		{
			name:     "extra spaces",
			input:    "   Pikachu    Eevee     Snorlax   ",
			expected: []string{"pikachu", "eevee", "snorlax"},
		},
		{
			name:     "tabs and newlines",
			input:    "Pikachu\nCharmander\tSquirtle",
			expected: []string{"pikachu", "charmander", "squirtle"},
		},
		{
			name:     "already lowercase",
			input:    "mewtwo gengar psyduck",
			expected: []string{"mewtwo", "gengar", "psyduck"},
		},
		{
			name:     "mixed case",
			input:    "JigglyPUFF PSYDUCK meOwTH",
			expected: []string{"jigglypuff", "psyduck", "meowth"},
		},
		{
			name:     "empty string",
			input:    "",
			expected: []string{},
		},
		{
			name:     "only whitespace",
			input:    "     \n\t   ",
			expected: []string{},
		},
		{
			name:     "pokemon with punctuation stays attached",
			input:    "Pikachu, I choose you!",
			expected: []string{"pikachu,", "i", "choose", "you!"},
		},
		{
			name:     "numbers included",
			input:    "Route 1 has 3 Pidgey",
			expected: []string{"route", "1", "has", "3", "pidgey"},
		},
	}



	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Test %s failed: expected %d words, got %d", c.name, len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Test %s failed: expected word %d to be %s, got %s", c.name, i, expectedWord, word)
			}
		}
	}



}


