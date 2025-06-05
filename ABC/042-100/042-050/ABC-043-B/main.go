package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s) // 文字列を入力

	chars := strings.Split(s, "") // 文字列を1文字ずつ分割

	buf := make([]string, 0, len(s)) // バッファを作成
	for _, c := range chars {
		if c == "0" || c == "1" {
			buf = append(buf, c)
		} else if c == "B" && len(buf) > 0 { // "B"が見つかり、バッファが空でない場合
			// バッファの末尾1文字を削除
			buf = buf[:len(buf)-1]
		}
	}
	fmt.Println(strings.Join(buf, "")) // バッファの内容を連結して出力
}
