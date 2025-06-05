package main

import (
	"fmt"
	"strings"
)

func main() {
	n := 0
	k := 0

	fmt.Scan(&n, &k)

	D := make([]string, k)
	for i := 0; i < k; i++ {
		fmt.Scan(&D[i])
	}

	// n以上の数値をチェックする
	for c := n; c <= 100000; c++ {
		// 数値に含まれる文字を取得
		chars := strings.Split(fmt.Sprintf("%d", c), "")
		errors := 0
		// 各文字がDに含まれているかチェック
		// 含まれていればエラーとする
		for i := range chars {
			for j := range D {
				if chars[i] == D[j] {
					errors++
					break
				}
			}
		}
		// エラーがなければ出力
		// それ以外は次の数値をチェック
		if errors > 0 {
			continue
		} else {
			fmt.Println(c)
			break
		}
	}
}
