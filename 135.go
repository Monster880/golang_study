func candy(ratings []int) (res int) {
	n := len(ratings)
	left := make([]int, n)
	right := make([]int, n)
	for i, r := range ratings {
		if i > 0 && r > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		} else {
			right[i] = 1
		}
		res += max(left[i], right[i])
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}