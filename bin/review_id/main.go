package main

import (
	"strings"
	"time"
)

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// int64をbase62文字列に変換する関数
func toBase62(num int64) string {
	if num == 0 {
		return string(base62Chars[0])
	}

	var result strings.Builder
	base := int64(len(base62Chars))

	for num > 0 {
		remainder := num % base
		result.WriteByte(base62Chars[remainder]) // 文字を追加
		num /= base
	}

	// 結果は逆順になっているため、元に戻す
	runes := []rune(result.String())
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {

	unixTime := time.Now().Unix()

	base62Name := toBase62(unixTime)

	println("review-" + base62Name) // 例: "1z0" などのbase62文字列が出力される
}
