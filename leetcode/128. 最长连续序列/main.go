package main

func longestConsecutive(nums []int) int {
	n := len(nums)
	numsMap := make(map[int]bool, n)

	for _, v := range nums {
		numsMap[v] = true
	}

	max := 0
	for i := 0; i < n; i++ {
		if !numsMap[nums[i]-1] {
			cur := nums[i]
			tmpmax := 0
			for numsMap[cur] {
				cur += 1
				tmpmax++
			}

			if tmpmax > max {
				max = tmpmax
			}
		}
	}

	return max
}
