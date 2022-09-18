package main

import (
	"fmt"
	"sort"
)

func MostAppearItem(items []string) string {
	var res string
	buat := map[string]int{}

	for i := 0; i < len(items); i++ {
		if _, found := buat[items[i]]; found {
			buat[items[i]]++
		} else {
			buat[items[i]] = 1
		}
	}

	sortVal := []int{}
	for _, val := range buat {
		sortVal = append(sortVal, val)
	}
	sort.Ints(sortVal)

	for j := 0; j < len(sortVal); j++ {
		for index, value := range buat {
			if value == sortVal[j] {
				res += fmt.Sprintf("%s -> %d ", index, value)
				buat[index] = -1
			}
		}

	}
	return res

}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
}
