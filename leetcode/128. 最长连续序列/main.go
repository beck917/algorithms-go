package main

/**
这个题目思路很巧妙,先用map去重

然后如果这个数字-1的值不在map中,就往后遍历
如果在就停止不进去循环,后面的数字可以进如循环

比如 4,3,2,1,前面三个都不会进入,到最后1才会进入,然后依次遍历出最大值
*/
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
				cur++
				tmpmax++
			}

			if tmpmax > max {
				max = tmpmax
			}
		}
	}

	return max
}
