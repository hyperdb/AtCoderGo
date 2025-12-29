package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	var result []string
	// 文字列を読み込み
	fmt.Scan(&S)
	// 一文字ずつに分割
	k := strings.Split(S, "")
	// 文字列を処理
	for i := range k {
		// 'B'なら直前の文字を削除
		if k[i] == "B" {
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
			continue
		}
		// 'B'でなければ結果に追加
		result = append(result, k[i])
	}
	// 結果を結合して出力
	fmt.Println(strings.Join(result, ""))
}
