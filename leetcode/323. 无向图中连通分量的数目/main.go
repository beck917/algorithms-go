package main

var visited map[int]bool

func countComponents(n int, edges [][]int) int {
	visited = make(map[int]bool, 0)

	edgesMap := make(map[int][]int, 0)
	for i := 0; i < len(edges); i++ {
		if _, ok := edgesMap[edges[i][0]]; !ok {
			edgesMap[edges[i][0]] = make([]int, 0)
		}
		edgesMap[edges[i][0]] = append(edgesMap[edges[i][0]], edges[i][1])

		if _, ok := edgesMap[edges[i][1]]; !ok {
			edgesMap[edges[i][1]] = make([]int, 0)
		}
		edgesMap[edges[i][1]] = append(edgesMap[edges[i][1]], edges[i][0])
	}

	// fmt.Println(edgesMap)

	num := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			//fmt.Println(i)
			dfs(i, edgesMap)
			num++
		}
	}

	return num
}

func dfs(node int, edgesMap map[int][]int) {
	visited[node] = true
	if len(edgesMap[node]) == 0 {
		return
	}
	//fmt.Println(node, edgesMap[node])

	for i := 0; i < len(edgesMap[node]); i++ {
		if !visited[edgesMap[node][i]] {
			dfs(edgesMap[node][i], edgesMap)
		}
	}

	return
}
