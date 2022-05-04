package main

import "sort"

func canAttendMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	max := 0
	tmp := -1
	for _, v := range intervals {
		if v[0] >= tmp {
			tmp = v[1]
			max++
		}
	}

	return max == len(intervals)
}
