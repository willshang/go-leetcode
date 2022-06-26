package main

import (
	"sort"
)

func main() {

}

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	for i := 0; i < len(spells); i++ {
		// xy>= success => y >= success/x => y > floor((success-1)/x)
		spells[i] = len(potions) - sort.SearchInts(potions, int(success-1)/spells[i]+1)
	}
	return spells
}
