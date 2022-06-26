package main

import (
	"sort"
)

func main() {

}

// leetcode2300_咒语和药水的成功对数
func successfulPairs(spells []int, potions []int, success int64) []int {
	n := len(spells)
	res := make([]int, n)
	sort.Ints(potions)
	for i := 0; i < n; i++ {
		index := sort.Search(len(potions), func(j int) bool {
			return int64(potions[j])*int64(spells[i]) >= success
		})
		res[i] = len(potions) - index
	}
	return res
}
