package main

var visited [][]int

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	visited = make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]int, len(grid[i]))
	}

	max := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 && visited[i][j] == 0 {
				//fmt.Println(i, j)
				num := dfs(grid, i, j)
				if num > max {
					max = num
				}
			}
		}
	}

	return max
}

func dfs(grid [][]int, i, j int) int {
	if grid[i][j] == 0 { // || visited[i][j] == 1
		return 0
	}

	visited[i][j] = 1

	var left, right, up, down int
	if i-1 >= 0 && visited[i-1][j] == 0 {
		left = dfs(grid, i-1, j)
	}
	if i+1 <= len(grid)-1 && visited[i+1][j] == 0 {
		right = dfs(grid, i+1, j)
	}

	if j-1 >= 0 && visited[i][j-1] == 0 {
		up = dfs(grid, i, j-1)
	}

	if j+1 <= len(grid[i])-1 && visited[i][j+1] == 0 {
		down = dfs(grid, i, j+1)
	}

	//fmt.Println(left,right,up,down)

	return left + right + up + down + 1
}
