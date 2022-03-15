package main

import "fmt"

func main() {
	s := "("
	fmt.Println(isValid(s))
}

/**
这个题乍看比较简单,思路就是栈,但其实没那么简单,一个点比较难想到,想到了就不难
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
就是这个,需要建立一个map, 注意这个map是反的!
如果字符不在这个map里,则入栈
如果在,则出栈,查看是否能够和map里的数据匹配上

最后如果len(stack) == 0,栈为空,则说明能够匹配
*/
func isValid(s string) bool {
	var stack []byte
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for i := 0; i < len(s); i++ {
		if _, ok := pairs[s[i]]; !ok {
			// 入栈
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			// 出栈
			tmpv := stack[len(stack)-1]
			if tmpv != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
