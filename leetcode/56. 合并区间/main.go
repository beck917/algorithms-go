package main

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	results := make([][]int, 0)

	var tmp []int
	for _, v := range intervals {
		if len(tmp) != 0 && tmp[1] >= v[0] {
			// 可以合并
			if v[1] > tmp[1] {
				tmp[1] = v[1]
			}
			if v[0] < tmp[0] {
				//tmp[0] = v[0] 这里不需要,已经排序过了
			}
		} else {
			if len(tmp) != 0 {
				results = append(results, tmp)
			}
			tmp = make([]int, 2)

			tmp[0] = v[0]
			tmp[1] = v[1]
		}
	}
	results = append(results, tmp)
	return results
}
