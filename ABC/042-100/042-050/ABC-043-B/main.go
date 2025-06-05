package main

import (
	"fmt"
	"strings"
)

func main() {
	s := ""
	fmt.Scan(&s) // 文字列を入力

	chars := strings.Split(s, "") // 文字列を1文字ずつ分割

	buf := make([]string, 0, len(s)) // バッファを作成
	for c := range chars {
		if chars[c] == "0" || chars[c] == "1" {
			buf = append(buf, chars[c])
		} else if chars[c] == "B" && len(buf) > 0 { // "B"が見つかり、バッファが空でない場合
			tmp := []string{}
			buf = append(tmp, buf[:len(buf)-1]...)
		}
	}
	fmt.Println(strings.Join(buf, "")) // バッファの内容を連結して出力
}
