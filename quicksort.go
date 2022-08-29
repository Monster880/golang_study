package main

func main() {
	list := []int{1, 2, 5, 4, 3}
	quickSort(list, 0, len(list)-1)
}

func quickSort(list []int, low, high int) {
	if high > low {
		pivot := partition(list, low, high)
		quickSort(list, low, pivot-1)
		quickSort(list, pivot+1, high)
	}
}

func partition(list []int, low, high int) {
	pivot := list[low]
	for low < high {
		for low < high && pivot <= list[high] {
			high--
		}
		list[low] = list[high]
		for low < high && pivot >= list[low] {
			low++
		}
		list[high] = list[low]
	}
	list[low] = pivot
	return low
}
