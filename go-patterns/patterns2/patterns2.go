package main

import "fmt"

func patterns2(x int) {
	for i := 0; i < x; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	x := 5
	patterns2(x)
}
