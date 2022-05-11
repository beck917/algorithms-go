package main

func trap(height []int) int {
	var leftMax, rightMax []int
	size := len(height)
	leftMax = make([]int, size)
	rightMax = make([]int, size)

	leftMax[0] = height[0]

	for i := 1; i < size; i++ {
		leftMax[i] = max(height[i], height[i-1])
	}

	rightMax[size-1] = height[size-1]
	for i := size - 2; i >= 0; i-- {
		rightMax[i] = max(height[i], rightMax[i-1])
	}

	result := 0
	for i := 0; i < size; i++ {
		result += max(0, max(leftMax[i], rightMax[i])-height[i])
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
