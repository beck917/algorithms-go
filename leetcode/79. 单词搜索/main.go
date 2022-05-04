package main

func exist(board [][]byte, word string) bool {

	h, w := len(board), len(board[0])
	visited := make([][]bool, h) // 只需要初始化一次,后面回溯的时候所有的visted都会重置为false
	for k := range visited {
		visited[k] = make([]bool, w)
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {

			if board[i][j] == word[0] {

				ret := dfs(board, word, 0, i, j, visited)
				if ret {
					return true
				}
			}
		}
	}

	return false
}

func dfs(board [][]byte, word string, k, i, j int, visited [][]bool) bool {
	if i > len(board)-1 || j > len(board[0])-1 || i < 0 || j < 0 {
		return false
	}

	//fmt.Println(string(board[i][j]), k, visited[i][j],i,j)
	if visited[i][j] {
		return false
	}

	if board[i][j] != word[k] {
		return false
	}

	if k == len(word)-1 {
		return true
	}

	visited[i][j] = true
	// 为什么要再visited改为true之后呢,是因为如果放在前面,visted==true直接返回了,可能会改变其他的递归的visted值
	defer func() { visited[i][j] = false }() //注意回溯的位置,一定要在所有返回之后

	return dfs(board, word, k+1, i+1, j, visited) || dfs(board, word, k+1, i, j+1, visited) ||
		dfs(board, word, k+1, i-1, j, visited) || dfs(board, word, k+1, i, j-1, visited)
}
