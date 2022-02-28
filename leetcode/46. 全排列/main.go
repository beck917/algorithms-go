package main

import "fmt"

func main() {
	nums := []int{5, 4, 6, 2}
	fmt.Println(permute(nums))
}

var results [][]int

/**
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
*/
func permute(nums []int) [][]int {
	n := len(nums)
	results = make([][]int, 0)
	used := make([]bool, n)
	for i, _ := range nums {
		used[i] = false
	}
	path := make([]int, 0)
	backTrack(nums, used, path)

	return results
}

func backTrack(nums []int, used []bool, path []int) {
	if len(path) == len(nums) {
		// 因为我们要输出最终的结果集,path是个数组,在回溯过程中,会被修改,这里我们需要将path拷贝一份
		p := make([]int, len(path))
		copy(p, path)
		results = append(results, p)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] == false {
			//fmt.Println(used, path, nums[i])
			used[i] = true

			path = append(path, nums[i])

			backTrack(nums, used, path)
			// 为什么输出[1,2,3]后, 回溯到[1]之后, 这里的循环还在继续, 运行了[1,2]之后, 这时轮到3了
			// path 记录1,2,3后, [3]第一个backtrack退出,另外循环也会退出(因为只有一个[3]),这时会将use[3]标记为false, 因为3的递归退出了
			// 那么循环到[2,3]这一层, 2这里的backtrack也会退出[因为[3]的循环退出了],use[2]也标记为false.[1,2,3]=>[2,3]=>[3]
			// 这时边回到了[1]这个循环, 由于2已经退出,那么循环继续到3
			used[i] = false
			// 这里需要注意和dfs图和树的区别
			// 有两点区别, 1.在dfs的时候我们记录路径,用的是链表结构,链表结构传递的是指针,后续循环的修改,不会对当前指针造成影响
			// 这里最主要的就是链表和数组的区别
			// 假设是链表,那么进行到[2,3]的时候,每次链表是1->2,1->3,因为在循环的过程中.next会被覆盖掉
			// 但如果是数组,那么[2,3]循环,就会变成[1,2],然后是[1,2,3], 所以如果用数组,一定要回溯
			path = path[0 : len(path)-1]

			// 另外注意一点,golang是值传递的,对于回溯内部对于path数组的变更,外部是感知不到的,这里只是对循环的回溯处理
		}
	}
}
