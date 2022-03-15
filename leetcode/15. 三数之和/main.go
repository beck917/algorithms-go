package main

import (
	"fmt"
	"sort"
)

func main() {
	//nums := []int{-1, 0, 1, 2, -1, -4}
	nums := []int{0, 0, 0}

	fmt.Println(threeSum(nums))
}

/**
知道思路很容易写出来,但问题是思路并不好想...
首先进行排序, 目的是三个数之和=0. 所以需要排序, 确保左边<右边

// 三指针，左右指针从最左和最右，逼近，直到相遇
从i=0开始,然后left=i+1, right=len(nums)-1

for left < right
然后tmpNum := nums[i] + nums[left] + nums[right]三数相加
<0 left++ >= right--
==0 就讲数组加入results

但这里需要注意去重,
第一个去重
1. 循环left, 对比left和left+是否数字一样
循环right,对比right和right-1数字是否一样
如果数字一样,意味着移动right或者left, 会导致重复的数字组合
2. i的去重,当i>0,并且nums[i]和前一个i-1的数据一样的时候, continue
因为这时,两个数据一样,也会导致重复的数据组合

总体来说,这题思路还是非常复杂的,而且需要去重,不是正常人思路,需要背诵阶梯方式
*/
func threeSum(nums []int) [][]int {
	// 先进行排序
	sort.Ints(nums)
	//fmt.Println(nums)

	results := make([][]int, 0)

	// 三指针，左右指针从最左和最右，逼近，直到相遇
	n := len(nums)
	for i := 0; i < n; i++ {
		left := i + 1
		right := len(nums) - 1

		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for left < right {
			tmpNum := nums[i] + nums[left] + nums[right]
			if tmpNum < 0 {
				left++
			} else if tmpNum > 0 {
				right--
			} else {
				results = append(results, []int{nums[i], nums[left], nums[right]})

				// 去重
				for left+1 <= n-1 && nums[left] == nums[left+1] {
					left++
				}
				for right-1 >= 0 && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}

	return results
}
