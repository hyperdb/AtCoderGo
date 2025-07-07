package main

import (
	"fmt"
)

func main() {
	var N, K, X, Y int

	// 入力を一度に読み込み
	fmt.Scan(&N, &K, &X, &Y)

	// 料金計算
	ans := 0
	if N <= K {
		// K個以下なら全て料金X
		ans = N * X
	} else {
		// K個は料金X、超過分は料金Y
		ans = K*X + (N-K)*Y
	}
	fmt.Println(ans)
}
