package main

import "fmt"

func alphabetBoardPath(target string) string {
	var ans string
	x, y := 0, 0
	for _, char := range []rune(target) {
		xx, yy := pos(char)
		for {
			if x == xx && y == yy {
				break
			}
			if x > xx && in(x-1, y) {
				x--
				ans += "U"
				continue
			}
			if x < xx && in(x+1, y) {
				x++
				ans += "D"
				continue
			}
			if y < yy && in(x, y+1) {
				y++
				ans += "R"
				continue
			}
			if y > yy && in(x, y-1) {
				y--
				ans += "L"
				continue
			}
			fmt.Println(ans)
		}
		ans += "!"
	}
	return ans
}

func in(x, y int) bool {
	if x < 5 && y < 5 {
		return true
	}
	if x == 5 && y == 0 {
		return true
	}
	return false
}

func pos(a rune) (int, int) {
	c := code(a)
	return c / 5, c % 5
}
func code(a rune) int {
	return (int)(a) - 97
}

func main() {
	fmt.Println(alphabetBoardPath("zdz"))
}
