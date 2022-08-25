var res [][]int

func permute(nums []int) [][]int {
	res = [][]int{}
	backTrack(nums, []int{}, len(nums))
	return res
}

func backTrack(nums []int, temp []int, length int) {
	if len(nums) == 0 {
		arr := make([]int, len(temp))
		copy(arr, temp)
		res = append(res, arr)
	}

	for i := 0; i < length; i++ {
		cur := nums[i]
		temp = append(temp, cur)
		nums = append(nums[:i], nums[i+1:]...)
		backTrack(nums, temp, len(nums))
		nums = append(nums[:i], append([]int{cur}, nums[i:]...)...)
		temp = temp[:len(temp)-1]
	}
}

