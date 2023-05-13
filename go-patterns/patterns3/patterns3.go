package main

import "fmt"

func pattern3(x int) {
	for i := 1; i <= x; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}

func main() {
	x := 5
	pattern3(x)
}
