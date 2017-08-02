package main

import (
	"fmt"
)

func main() {
	const(
		X = 2
		Y = 7
		Z = X + Y

		S1 = "今日"
		S2 = "晴れ"
		S = S1 + "は" + S2
	)

	fmt.Printf(S)
}