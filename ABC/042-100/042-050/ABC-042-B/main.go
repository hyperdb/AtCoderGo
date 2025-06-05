package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var n, l int
	fmt.Scan(&n, &l)

	// n個の文字列を格納するスライスを作成
	s := make([]string, n)

	// 文字列を入力
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}

	// 辞書順にソート
	sort.Strings(s)

	// すべての文字列を連結
	result := strings.Join(s, "")

	// 結果を出力
	fmt.Println(result)
}
