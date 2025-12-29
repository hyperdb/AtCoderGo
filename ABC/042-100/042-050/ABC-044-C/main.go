package main

import (
	"fmt"
	"math"
)

func getCost(data []int, ave float64) int {
	cost := 0
	for i := range(data) {
		cost += int(math.Pow(float64(data[i]) - ave, 2))
	}
	return cost
}


func main() {
	var n int
	fmt.Scan(&n)

	// nが0の場合の特別処理
	if n == 0 {
		fmt.Println("0")
	} else {
		var p []int

		// 整数列を読み込み
		p = make([]int, n)
		for i := range(n) {
			fmt.Scan(&p[i])
		}

		// 平均値を計算
		var sumP int
		for i := range(p) {
			sumP += p[i]
		}
		aveP := float64(sumP) / float64(len(p))

		// 小数部を取り出し
		_, frac := math.Modf(aveP)

		// 平均が整数なら
		var cost int
		if frac == 0 {
			cost = getCost(p, aveP)
		} else {
			// 平均が整数でないなら、切り上げと切り捨ての両方を試す
			cost_floor := getCost(p, math.Floor(aveP))
			cost_ceil := getCost(p, math.Ceil(aveP))
			// 小さい方を選ぶ
			cost = min(cost_floor, cost_ceil)
		}
		fmt.Println(cost)
	}


}
