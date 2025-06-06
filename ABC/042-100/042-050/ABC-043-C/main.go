package main

import (
	"fmt"
	"math"
)

func getIntList(N int) []int {
	array := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&array[i])
	}
	return array
}

func getCost(data []int, average int) int {
	cost := 0.0
	for i := 0; i < len(data); i++ {
		cost += math.Pow(float64(data[i]-average), 2.0) // 平均からの偏差の二乗を計算
	}
	return int(cost)
}

func main() {
	N := 0
	fmt.Scan(&N) // 数字を入力

	if N == 0 {
		fmt.Println(0) // Nが0の場合は0を出力
	} else {
		a := getIntList(N) // 数字のリストを取得
		sum := 0
		for i := 0; i < N; i++ {
			sum += a[i] // リストの合計を計算
		}
		average := float64(sum) / float64(N) // 平均を計算

		cost := 0
		if math.Floor(average) == average {
			cost = getCost(a, int(average)) // 平均が整数の場合
		} else {
			cost_low := getCost(a, int(math.Floor(average))) // 平均の下の整数
			cost_high := getCost(a, int(math.Ceil(average))) // 平均の上の整数

			if cost_low < cost_high {
				cost = cost_low // 下の整数の方がコストが小さい場合
			} else {
				cost = cost_high // 上の整数の方がコストが小さい場合
			}
		}
		fmt.Println(cost)
	}
}
