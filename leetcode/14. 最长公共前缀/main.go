package main

/**
输入：strs = ["flower","flow","flight"]
输出："fl"
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := ""
LOOP:
	for i := 0; i < len(strs[0]); i++ {
		b := strs[0][i]

		for j := 1; j < len(strs); j++ {
			if strs[j] != b {
				break LOOP
			}
		}
		prefix += string(b)
	}

	return prefix
}
