package main

import "fmt"

func testClose() {
	ch := make(chan int, 4)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	for x := range ch {
		fmt.Println(x)
	}
}
func main() {
	testClose()
}
