sort.Slice(intSlice, func(i, j int) bool {
	a := strconv.Itoa(intSlice[i])
	b := strconv.Itoa(intSlice[j])
	return a+b > b+a
})
