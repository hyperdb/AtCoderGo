package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	n := 0
	l := 0

	fmt.Scan(&n, &l)

	s := make([]string, l, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}

	sort.SliceStable(s, func(i, j int) bool { return s[i] < s[j] })

	r := strings.Join(s, "")

	fmt.Println(r)
}
