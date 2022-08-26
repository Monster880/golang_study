import "sort"

func permuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	temp := []int{}
	visited := make([]bool, n)
	var backtrack func(int)
	backtrack = func(index int) {
		if index == n {
			ans = append(ans, append([]int{}, temp...))
			return
		}
		for i, value := range nums {
			if visited[i] || i > 0 && !visited[i-1] && value == nums[i-1] {
				continue
			}
			temp = append(temp, value)
			visited[i] = true
			backtrack(index + 1)
			visited[i] = false
			temp = temp[:len(temp)-1]
		}
	}
	backtrack(0)
	return
}
