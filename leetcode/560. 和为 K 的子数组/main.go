package main

func subarraySum(nums []int, k int) int {
	preSum := 0 //前缀和, 所有前面的数字和

	flagMap := make(map[int]int, 0) // 记录前缀和是否有的map
	flagMap[0] = 1                  //注意初始化为1

	max := 0 //返回值,总数

	for i := 0; i < len(nums); i++ {
		preSum += nums[i]

		if v, ok := flagMap[preSum-k]; ok {
			max += v
		}

		flagMap[preSum] += 1 // 这里要注意负数的情况,是有可能产生多个值的
	}

	return max
}
