package main

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	results := make([][]int, 0)

	var tmp []int
	for i := 0; i < 2; i++ {
		tmp = append(tmp, -1)
	}
	for _, v := range intervals {
		if tmp[1] >= v[0] {
			// 可以合并
			if v[1] > tmp[1] {
				tmp[1] = v[1]
			}
			if v[0] < tmp[0] {
				tmp[0] = v[0]
			}
		} else {
			if tmp[0] != -1 {
				p := make([]int, 2)
				copy(p, tmp)
				results = append(results, p)
			}
			tmp[0] = v[0]
			tmp[1] = v[1]
		}
	}
	results = append(results, tmp)
	return results
}
