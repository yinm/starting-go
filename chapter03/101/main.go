package main

import "fmt"

func doSomething() (x, y int) {
	y = 5
	return
}

func main() {
	fmt.Println(doSomething())
}
