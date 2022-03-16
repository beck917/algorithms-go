package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

var results []string

func generateParenthesis(n int) []string {
	results = make([]string, 0)
	dfs(n, 0, 0, "")

	return results
}

func dfs(n int, left, right int, path string) {
	if left > n || right > left {
		return
	}

	if len(path) >= 2*n {
		results = append(results, path)
		return
	}

	dfs(n, left+1, right, path+"(")
	dfs(n, left, right+1, path+")")
}
