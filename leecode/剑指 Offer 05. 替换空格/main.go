package main

import "fmt"

func main() {
	fmt.Println(replaceSpace("We are happy."))
}

// 1. 先扩容
// 2. 双指针法，注意，是从后往前的双指针。
// 3. 需要注意的是，i指针的数组长度的s，而j指针的数组长度为扩容厚的ss。这样i如果不为空，就直接改变j的值
func replaceSpace(s string) string {
	//扩容
	tmps := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			tmps += "  "
		}
	}
	ss := s + tmps

	bs := []byte(ss)
	j := len(bs) - 1
	for i := len(s) - 1; i >= 0; i-- {
		if bs[i] != ' ' {
			bs[j] = bs[i]
			j--
		} else {
			bs[j] = '0'
			bs[j-1] = '2'
			bs[j-2] = '%'
			j -= 3
		}
	}

	return string(bs)
}
