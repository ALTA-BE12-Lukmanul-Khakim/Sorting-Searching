package main

import (
	"fmt"
	"sort"
)

func MaximumBuyProduct(money int, productPrice []int) int {
	sort.Ints(productPrice)
	res := 0
	//var res1 int
	var sisa int

	if money < productPrice[0] {
		res += 0
	} else if money > productPrice[0] {
		sisa = money - productPrice[0]
		res += 1
	}

	for i := 1; i < len(productPrice); i++ {
		if sisa >= productPrice[i] {
			sisa = sisa - productPrice[i]
			res = res + 1
		}
		if sisa < productPrice[i] {
			res = res + 0
		}

	}
	//fmt.Println(res)
	return res
}

func main() {
	fmt.Println(MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000}))
	fmt.Println(MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000}))
	fmt.Println(MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000}))
	fmt.Println(MaximumBuyProduct(4000, []int{7500, 3000, 2500, 2000}))
	fmt.Println(MaximumBuyProduct(0, []int{10000, 30000}))

}
