package main

import "fmt"

func main() {
	fmt.Printf("%T\n", func(x, y int) int { return x + y })
}
