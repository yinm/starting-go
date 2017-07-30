package main

import "fmt"

func plus(x, y int) int {
	return x + y
}

func main() {
	result := plus(3, 5)
	fmt.Printf("和=%d\n", result)
}
