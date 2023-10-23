package main

import (
	"math"
	"sort"
)

func main() {

}

func minimizedMaximum(n int, quantities []int) int {
	return sort.Search(100000, func(i int) bool {
		return judge(quantities, i) <= n
	})
}

func judge(quantities []int, per int) int {
	if per == 0 { // 注意除0错误
		return math.MaxInt32
	}
	res := 0
	for i := 0; i < len(quantities); i++ {
		// res = res + (quantities[i]+per-1)/per
		res = res + quantities[i]/per
		if quantities[i]%per != 0 {
			res++
		}
	}
	return res
}
