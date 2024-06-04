package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{0, 0, 0}

	fmt.Scan(&n[0], &n[1], &n[2])

	sort.SliceStable(n, func(i, j int) bool { return n[i] < n[j] })

	sum := 0
	for _, x := range n {
		sum += x
	}

	r := "NO"
	if sum == 17 && n[0] == 5 && n[1] == 5 {
		r = "YES"
	}
	fmt.Println(r)
}
