package main

import "fmt"

func main() {
	x := 5
	if x := 2; true {
		fmt.Printf("%d\n", x)
	}

	fmt.Printf("%d\n", x)
}
