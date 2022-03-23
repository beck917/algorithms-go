package shopee

var visited map[int]bool

func countComponents(n int, edges [][]int) int {
	visited = make(map[int]bool, 0)

	edgesMap := make(map[int][]int, 0)
	for i := 0; i < len(edges); i++ {
		if _, ok := edgesMap[edges[i][0]]; !ok {
			edgesMap[edges[i][0]] = make([]int, 0)
		}
		edgesMap[edges[i][0]] = append(edgesMap[edges[i][0]], edges[i][1])
	}

	num := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(i, edgesMap)
			num++
		}
	}

	return num
}

func dfs(node int, edgesMap map[int][]int) {
	if len(edgesMap[node]) == 0 {
		return
	}

	visited[node] = true

	for i := 0; i < len(edgesMap[node]); i++ {
		if !visited[edgesMap[node][i]] {
			dfs(edgesMap[node][i], edgesMap)
		}
	}

	return
}
