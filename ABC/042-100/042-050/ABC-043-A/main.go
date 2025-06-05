package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N) // Nを入力

	// 1からNまでの和を公式で計算して出力
	fmt.Println((N + 1) * N / 2)
}
