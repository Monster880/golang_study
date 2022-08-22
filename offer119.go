import (
    "sort"
    "math"
    "fmt"
)
func longestConsecutive(nums []int) int {
    tempArr := uniqueArr(nums)
    // fmt.Println(tempArr[0])
    sort.Ints(tempArr)
    if len(tempArr) < 2{
        return len(tempArr)
    }
    res := 1
    left := 0
    n := len(tempArr)
    for right := 1; right < n; right++{
        // fmt.Println(tempArr[right])
        // fmt.Println(tempArr[right-1])
        if tempArr[right]-1 != tempArr[right-1]{
            res = int(math.Max(float64(res), float64(right - left)))
            left = right
            // fmt.Println(res)
        }
        res = int(math.Max(float64(res), float64(right - left + 1)))
    }
    return res
}

func uniqueArr(nums []int) []int{
    newArr := make([]int, 0)
    tempMap := make(map[int]bool, len(nums))
    for _, v := range nums{
        if tempMap[v] == false{
            tempMap[v] = true
            newArr = append(newArr,v)
        }
    }
    return newArr
}

// func uniqueArr(nums []int) []int{
//     newArr := make([]int, 0)
//     tempMap := make(map[int]bool, len(nums))
//     for _, v := range nums{
//         if tempMap[v] == false{
//             tempMap[v] = true
//             newArr = append(newArr, v)
//         }
//     }
//     return newArr
// }