package main

func main() {

}

func removeElement(nums []int, val int) int {
	j := 0
	count := len(nums)
	for i := 0; i < count; i++ {
		if nums[i] != val {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}

	return j
}
