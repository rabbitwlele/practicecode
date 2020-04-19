package main

import "fmt"

func test1() {
	{
		defer func() {
			fmt.Println(1)
		}()
	}

	{
		defer func() {
			fmt.Println(2)
		}()
	}
}
func main() {
	test1()
}
