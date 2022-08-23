func findCircleNum(isConnected [][]int) (res int) {
	visited := make([]bool, len(isConnected))
	var dfs func(int)
	dfs = func(start int) {
		visited[start] = true
		for end, connection := range isConnected[start] {
			if connection == 1 && !visited[end] {
				dfs(end)
			}
		}
	}
	for start, visit := range visited {
		if !visit {
			res++
			dfs(start)
		}
	}
	return res
}