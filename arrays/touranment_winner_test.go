package arrays

import "testing"

func TestTournamentWinner(t *testing.T) {
	tests := []struct {
		competitions [][]string
		results      []int
		winner       string
	}{
		// Test Case: Java wins twice
		{
			[][]string{
				{"HTML", "Java"},
				{"Java", "Python"},
				{"Python", "HTML"},
			},
			[]int{0, 1, 1},
			"Java",
		},

		// Test Case: Bulls win all matches
		{
			[][]string{
				{"Bulls", "Eagles"},
				{"Bulls", "Bears"},
				{"Bears", "Eagles"},
			},
			[]int{1, 0, 0},
			"Bulls",
		},
		// Test Case: Bulls win all matches
		{
			[][]string{
				{"HTML", "Java"},
				{"Java", "Python"},
				{"Python", "HTML"},
				{"C#", "Python"},
				{"Java", "C#"},
				{"C#", "HTML"},
				{"SQL", "C#"},
				{"HTML", "SQL"},
				{"SQL", "Python"},
				{"SQL", "Java"},
			},
			[]int{0, 1, 1, 1, 0, 1, 0, 1, 1, 0},
			"C#",
		},
	}

	for i, test := range tests {
		got := TournamentWinner(test.competitions, test.results)

		if got != test.winner {
			t.Errorf("Test case #%d failed: expected %v, got %v", i, test.winner, got)
		}
	}
}
