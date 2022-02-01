package easy

const HOME_TEAM_WON = 1

func TournamentWinner(competitions [][]string, results []int) string {

	winner := ""
	points := 0
	var counterMap map[string]int = make(map[string]int)

	for i := 0; i < len(competitions); i++ {
		currentWinner := competitions[i][HOME_TEAM_WON-results[i]]
		currentPoint, ok := counterMap[currentWinner]
		if !ok {
			// first time occurance
			counterMap[currentWinner] = 3

		} else {
			counterMap[currentWinner] = 3 + currentPoint
		}
		if counterMap[currentWinner] > points {
			points = counterMap[currentWinner]
			winner = currentWinner
		}
	}

	return winner
}
