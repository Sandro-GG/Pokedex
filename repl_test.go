package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello   world    ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello   WORld  world woRLD ",
			expected: []string{"hello", "world", "world", "world"},
		},
		{
			input:    "   PIKACHu   bulbaSAUR charmander   MAGIKARP",
			expected: []string{"pikachu", "bulbasaur", "charmander", "magikarp"},
		},
		{
			input:    "ONYX",
			expected: []string{"onyx"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Number of words don't match for \"%s\"", c.input)
			continue
		}

		for i, word := range actual {
			if word != c.expected[i] {
				t.Errorf("input: \"%v\"; expected: \"%v\"", word, c.expected[i])
			}
		}
	}

}
