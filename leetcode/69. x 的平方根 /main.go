package main

func mySqrt(x int) int {
	left := 0
	right := x
	res := 0
	for left <= right {
		mid := (left + right) >> 1
		mid2 := mid * mid
		if mid2 > x {
			right = mid - 1
		} else if mid2 < x {
			// 首次<x的时候记录下mid
			res = mid
			left = mid + 1
		} else {
			res = mid
			break
		}
	}

	return res
}
