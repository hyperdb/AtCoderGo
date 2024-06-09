package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	n := []int{0, 0, 0}

	fmt.Scan(&n[0], &n[1], &n[2])

	sort.SliceStable(n, func(i, j int) bool { return n[i] < n[j] })

	t := []int{5, 5, 7}
	r := "NO"
	if reflect.DeepEqual(n, t) {
		r = "YES"
	}
	fmt.Println(r)
}
