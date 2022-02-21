package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(compressString(""))
}

/**
字符串压缩。利用字符重复出现的次数，编写一种方法，实现基本的字符串压缩功能。比如，字符串aabcccccaaa会变为a2b1c5a3。若“压缩”后的字符串没有变短，则返回原先的字符串。你可以假设字符串中只包含大小写英文字母（a至z）。

示例1:

 输入："aabcccccaaa"
 输出："a2b1c5a3"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/compress-string-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func compressString(S string) string {
	output := ""
	tmp := ""
	count := 0
	for i := 0; i < len(S); i++ {
		vs := string(S[i])
		if tmp != vs {
			output += strconv.Itoa(count) + vs
			tmp = vs
			count = 1
		} else {
			count++
		}
	}
	output += strconv.Itoa(count)

	outputNew := output[1:]
	if len(outputNew) >= len(S) {
		outputNew = S
	}
	return outputNew
}
