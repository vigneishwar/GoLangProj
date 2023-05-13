package main

import "fmt"

func patterns4(x int) {
	for i := 1; i <= x; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(i, " ")
		}
		fmt.Println()
	}
}

func main() {
	x := 6
	patterns4(x)
}
