package main

func allSame(a []string) string {
	for i := range a[:len(a)-1] {
		if a[i] != a[i+1] {
			return ""
		}
	}
	return a[0]
}

func getWinner(game [][]string) string {
	n := len(game)

	// rows
	for _, row := range game {
		if winner := allSame(row); winner != "" {
			return winner
		}
	}

	// columns
	for j := 0; j < n; j++ {
		column := make([]string, 0, n)
		for _, row := range game {
			column = append(column, row[j])
		}
		winner := allSame(column)
		if winner != "" {
			return winner
		}
	}

	// diagonals
	diagonal := make([]string, 0, n)
	for i := 0; i < n; i++ {
		diagonal = append(diagonal, game[i][i])
	}
	winner := allSame(diagonal)
	if winner != "" {
		return winner
	}

	diagonal = make([]string, 0, n)
	for i := 0; i < n; i++ {
		j := n - 1 - i
		diagonal = append(diagonal, game[i][j])
	}
	winner = allSame(diagonal)
	if winner != "" {
		return winner
	}

	return ""
}
