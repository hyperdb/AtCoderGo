package main

import (
	"fmt"
	"sort"
)

func main() {
	// 3つの整数を受け取る
	nums := []int{0, 0, 0}
	fmt.Scan(&nums[0], &nums[1], &nums[2])

	// 昇順にソート（順序に依存しない判定のため）
	sort.Ints(nums)

	// 判定したい配列（5, 5, 7の組み合わせかどうか）
	target := []int{5, 5, 7}

	// 配列が一致するか判定
	result := "NO"
	if nums[0] == target[0] && nums[1] == target[1] && nums[2] == target[2] {
		result = "YES"
	}
	fmt.Println(result)
}
