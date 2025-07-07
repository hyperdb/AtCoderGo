package main

import (
	"fmt"
)

func main() {
	N := 0
	K := 0
	X := 0
	Y := 0

	fmt.Scan(&N)
	fmt.Scan(&K)
	fmt.Scan(&X)
	fmt.Scan(&Y)

	ans := 0
	if N <= K {
		ans = N * X
	} else {
		ans = K*X + (N-K)*Y
	}
	fmt.Println(ans)
}
