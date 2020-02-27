package main

import (
	"fmt"
)

func myAtoi(str string) int {
	ans := 0
	isFu := false
	isBegin := false
	for _, ch := range str {
		if ch == ' ' && !isBegin {
			continue
		}

		if '0' <= ch && ch <= '9' {
			ans *= 10
			if isFu {
				ans -= int(ch - '0')
			} else {
				ans += int(ch - '0')
			}
		} else if ch == '-' {
			if isBegin {
				break
			}
			isFu = true
		} else if ch == '+' {
			if isBegin {
				break
			}
		} else {
			break
		}
		if ans < -(1 << 31) {
			return -(1 << 31)
		}
		if ans > (1<<31)-1 {
			return (1 << 31) - 1
		}
		if ch != ' ' {
			isBegin = true
		}
	}
	return ans
}

func main() {
	fmt.Println(myAtoi("41"))
	fmt.Println(myAtoi("  -123asdf"))
	fmt.Println(myAtoi("-21312-"))
	fmt.Println(myAtoi("a123"))
	fmt.Println(myAtoi("+123"))
	fmt.Println(myAtoi("-aw213123"))
	fmt.Println(myAtoi("213124124124124124124124124321"))
	fmt.Println(myAtoi("-213124124124124124124124124321"))
}
