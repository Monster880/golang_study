import "sort"

func combinationSum2(candidates []int, target int) (res [][]int) {
	sort.Ints(candidates)
	var dfs func(index int, temp []int, target int)
	dfs = func(index int, temp []int, target int) {
		if target <= 0 {
			if target == 0 {
				arr := make([]int, len(temp))
				copy(arr, temp)
				res = append(res, arr)
			}
			return
		}
		for i := index; i < len(candidates); i++ {
			if i-1 >= index && candidates[i-1] == candidates[i] {
				continue
			}
			temp = append(temp, candidates[i])
			dfs(i+1, temp, target-candidates[i])
			temp = temp[:len(temp)-1]
		}
	}
	dfs(0, []int{}, target)
	return
}
