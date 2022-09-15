package main

import "fmt"

func FindMaxAndMin(arr []int) string {
	var res1, res2 string
	max := 0
	min := 1000
	i := 1
	j := 1
	for i, _ = range arr {
		if arr[i] > max {
			max = arr[i]
			res1 = fmt.Sprintf("max: %d  index: %d", arr[i], i)
		}
	}

	for j, _ = range arr {
		if arr[j] < min {
			min = arr[j]
			res2 = fmt.Sprintf("min: %d  index: %d", arr[j], j)
		} else {
			min = arr[0]
		}

	}
	return res1 + " " + res2

	// for i := 1; i < len(arr); i++ {
	// 	if arr[i] > max {
	// 		max = arr[i]
	// 	}
	// 	if arr[i] < min {
	// 		min = arr[i]
	// 	}
	// 	res += "min"
	// }
}

func main() {
	fmt.Println(FindMaxAndMin([]int{5, 7, 4, -2, -1, 8}))
	fmt.Println(FindMaxAndMin([]int{2, -5, -4, 22, 7, 7}))
	fmt.Println(FindMaxAndMin([]int{4, 3, 9, 4, -21, 7}))
	fmt.Println(FindMaxAndMin([]int{-1, 5, 6, 4, 2, 18}))
	fmt.Println(FindMaxAndMin([]int{-2, 5, -7, 4, 7, -20}))
}
