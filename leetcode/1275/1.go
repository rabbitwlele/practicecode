package main

import "fmt"


func tictactoe(moves [][]int) string {
	var m [3][3]int
	for idx, move := range moves {
		if id%2 == 0 {
			m[move[0]][move[1]] = 1
		} else {
			m[move[0]][move[1]] = -1
		}
	}

	for i := 0; i < 3; i++ {
		sum := 0
		for j := 0; j < 3; j++ {
			sum += m[i][j]
		}
		if sum == 3 {
			return "A"
		}
		if sum == -3 {
			return "B"
		}
	}
	for i := 0; i < 3; i++ {
		sum := 0
		for j := 0; j < 3; j++ {
			sum += m[j][i]
		}
		if sum == 3 {
			return "A"
		}
		if sum == -3 {
			return "B"
		}
	}

	sum := m[0][0] + m[1][1] + m[2][2]
	if sum == 3 {
		return "A"
	}
	if sum == -3 {
		return "B"
	}

	sum = m[0][2] + m[1][1] + m[2][0]
	if sum == 3 {
		return "A"
	}
	if sum == -3 {
		return "B"
	}

	if len(moves) == 9 {
		return "Draw"
	}
	return "Pending"
}
func main() {
	fmt.Println("xxx")
}
