package arrays

const HOME_TEAM_WON = 1

func TournamentWinner(competitions [][]string, results []int) string {
	// Write your code here.
	currentBestteam := ""
	scores := map[string]int{currentBestteam: 0}

	for i := 0; i < len(competitions); i++ {
		competition := competitions[i]
		home := competition[0]
		away := competition[1]
		result := results[i]
		winner := away

		if result == HOME_TEAM_WON {
			winner = home
		}

		scores[winner] = scores[winner] + 1

		if scores[winner] > scores[currentBestteam] {
			currentBestteam = winner
		}

	}

	return currentBestteam
}
