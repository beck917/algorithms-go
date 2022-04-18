package main

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		sx := strconv.Itoa(x)
		sy := strconv.Itoa(y)

		return sx+sy > sy+sx
	})

	if nums[0] == 0 {
		return "0"
	}

	var ret string
	for i := 0; i < len(nums); i++ {
		ret += strconv.Itoa(nums[i])
	}

	return ret
}
