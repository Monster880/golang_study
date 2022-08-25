func combinationSum(candidates []int, target int) (res [][]int) {
	combination := []int{}
	var dfs func(target, index int)
	dfs = func(target, index int) {
		if index == len(candidates) {
			return
		}
		if target == 0 {
			res = append(res, append([]int{}, combination...))
			return
		}
		dfs(target, index+1)
		if target-candidates[index] >= 0 {
			combination = append(combination, candidates[index])
			dfs(target-candidates[index], index)
			combination = combination[:len(combination)-1]
		}
	}
	dfs(target, 0)
	return
}
