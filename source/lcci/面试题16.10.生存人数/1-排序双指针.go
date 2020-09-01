package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxAliveYear([]int{1900, 1901, 1950}, []int{1948, 1951, 2000}))
}

func maxAliveYear(birth []int, death []int) int {
	sort.Ints(birth)
	sort.Ints(death)
	res := birth[0]
	max := 0
	j := 0
	count := 0
	for i := 0; i < len(birth); i++ {
		count++
		for birth[i] > death[j] {
			count--
			j++
		}
		if count > max {
			max = count
			res = birth[i]
		}
	}
	return res
}
