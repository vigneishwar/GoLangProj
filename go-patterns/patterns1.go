package main

import "fmt"

func pattern1(x int) {

	for i := 0; i < x; i++ {
		for j := 0; j <= x; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func main() {
	x := 4
	pattern1(x)
}
