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
		chars := strings.Split(fmt.Sprintf("%d", c), "")
		found := false // 禁止数字が見つかったかどうかのフラグ

		// 各文字がDに含まれているかチェック
		for _, ch := range chars {
			for _, d := range D {
				if ch == d {
					found = true
					break
				}
			}
			if found {
				break // 1つでも禁止数字があれば次の数値へ
			}
		}

		// 禁止数字が含まれていなければ出力して終了
		if !found {
			fmt.Println(c)
			break
		}
	}
}
