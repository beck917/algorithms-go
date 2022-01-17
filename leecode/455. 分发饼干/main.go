package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	j := 0
	num := 0
	// 小孩找饼干
	for i := 0; i <= len(g)-1; {
		if j > len(s)-1 {
			break
		}
		if g[i] <= s[j] {
			i++
			num++
		}
		j++
	}

	return num
}
