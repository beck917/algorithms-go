package main

import "fmt"

func main() {
	grid := make([][]byte, 0)
	grid = append(grid, []byte{'1', '0', '1', '1', '1'})
	grid = append(grid, []byte{'1', '0', '1', '0', '1'})
	grid = append(grid, []byte{'1', '1', '1', '0', '1'})
	fmt.Println(numIslands(grid))
}

/**
输入：grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["1","1","0","1","1"]
]
[["1","1","1"],["0","1","0"],["1","1","1"]]
[['1','0','1','1','1'],
 ['1','0','1','0','1'],
 ['1','1','1','0','1']]
输出：3
*/
func numIslands(grid [][]byte) int {

	parentNodes := make([][]int, 0)

	num := 0
	visted := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visted[i] = make([]bool, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			visted[i][j] = false
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if visted[i][j] == false && grid[i][j] == '1' {
				fmt.Println(i, j)
				parentNodes = append(parentNodes, []int{i, j})
				addIsland(grid, parentNodes, visted)
				num++
			}
		}
	}
	return num
}

func addIsland(grid [][]byte, parentNodes [][]int, visted [][]bool) [][]int {
	var otherNodes [][]int
	for len(parentNodes) != 0 {
		tmpNodes := make([][]int, 0)
		otherNodes = make([][]int, 0)
		//visted := make(map[string]bool)
		for _, v := range parentNodes {
			i := v[0]
			j := v[1]
			if visted[i][j] == true {
				continue
			}
			visted[i][j] = true
			if i+1 < len(grid) {
				if grid[i+1][j] == '1' {
					tmpNodes = append(tmpNodes, []int{i + 1, j})
				} else {
					otherNodes = append(otherNodes, []int{i + 1, j})
				}

			}
			if j+1 < len(grid[0]) {
				if grid[i][j+1] == '1' {
					tmpNodes = append(tmpNodes, []int{i, j + 1})
				} else {
					otherNodes = append(otherNodes, []int{i, j + 1})
				}

			}
			if i-1 >= 0 {
				if grid[i-1][j] == '1' && visted[i-1][j] == false {
					tmpNodes = append(tmpNodes, []int{i - 1, j})
				}
				//visted[i-1][j] = true
			}
			if j-1 >= 0 {
				if grid[i][j-1] == '1' && visted[i][j-1] == false {
					tmpNodes = append(tmpNodes, []int{i, j - 1})
				}
			}
		}
		parentNodes = tmpNodes
	}

	return otherNodes
}
