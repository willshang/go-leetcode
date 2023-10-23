package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []string{"abc", "acb", "bac", "bca", "cab", "cba"}
	fmt.Println(numSpecialEquivGroups(arr))
}

func numSpecialEquivGroups(A []string) int {
	groups := make(map[string]bool)
	for _, a := range A {
		odd := make([]byte, 0)
		even := make([]byte, 0)
		for i := 0; i < len(a); i++ {
			if i%2 == 0 {
				even = append(even, a[i]-'a')
			} else {
				odd = append(odd, a[i]-'a')
			}
		}
		sort.Slice(odd, func(i, j int) bool {
			return odd[i] < odd[j]
		})
		sort.Slice(even, func(i, j int) bool {
			return even[i] < even[j]
		})
		groups[string(odd)+string(even)] = true
	}
	return len(groups)
}
