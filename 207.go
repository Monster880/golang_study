func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges    = make([][]int, numCourses)
		indegree = make([]int, numCourses)
		result   []int
	)

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
		indegree[info[0]]++
	}

	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		temp := queue[0]
		queue = queue[1:]
		result = append(result, temp)
		for _, v := range edges[temp] {
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	return len(result) == numCourses
}
