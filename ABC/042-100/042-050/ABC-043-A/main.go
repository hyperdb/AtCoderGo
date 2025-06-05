package main

import (
	"fmt"
)

func main() {
	N := 0
	fmt.Scan(&N)

	fmt.Println(int(((N + 1) * N) / 2))
}
