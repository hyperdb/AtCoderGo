/*
 * ABC-044-C
 */
package main

import (
	"fmt"
	"math"
)

func getCost(data []int, ave float64) int {
	cost := 0
	for i := range data {
		cost += int(math.Pow(float64(data[i])-ave, 2))
	}
	return cost
}

func main() {
	var num int
	var target int
	fmt.Scan(&num, &target)

	// nが0の場合の特別処理
	if num == 0 {
		fmt.Println("0")
	} else {
		var cards []int

		// 整数列を読み込み
		cards = make([]int, num)
		for i := range num {
			fmt.Scan(&cards[i])
		}

		// 平均値を計算
		var sum int
		for i := range cards {
			sum += cards[i]
		}
		ave := float64(sum) / float64(len(cards))

		// 小数部を取り出し
		_, frac := math.Modf(ave)

		// 平均が整数なら
		var cost int
		if frac == 0 {
			cost = getCost(cards, ave)
		} else {
			// 平均が整数でないなら、切り上げと切り捨ての両方を試す
			cost_floor := getCost(cards, math.Floor(ave))
			cost_ceil := getCost(cards, math.Ceil(ave))
			// 小さい方を選ぶ
			cost = min(cost_floor, cost_ceil)
		}
		fmt.Println(cost)
	}

}
