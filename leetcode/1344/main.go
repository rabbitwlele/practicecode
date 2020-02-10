package main

import (
	"fmt"
)

func angleClock(hour int, minutes int) float64 {
	if hour == 12 {
		hour = 0
	}
	hh := float64(hour)*5*6 + float64(minutes)*5*6/60
	mm := float64(minutes) * 6
	//fmt.Println(hh, mm)
	ans := hh - mm
	if ans < 0 {
		ans = -ans
	}
	if ans > 180{
		ans = 360-ans
	}
	return ans
}
func main() {
	fmt.Println(angleClock(12, 30))
	fmt.Println(angleClock(3, 30))
	fmt.Println(angleClock(3, 15))
	fmt.Println(angleClock(4, 50))
	fmt.Println(angleClock(12, 0))
	fmt.Println(angleClock(1, 57))
}
