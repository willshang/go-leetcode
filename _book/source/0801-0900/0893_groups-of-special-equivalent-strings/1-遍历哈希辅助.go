package main

import "fmt"

func main() {
	arr := []string{"abc", "acb", "bac", "bca", "cab", "cba"}
	fmt.Println(numSpecialEquivGroups(arr))
}

// leetcode893_特殊等价字符串组
func numSpecialEquivGroups(A []string) int {
	groups := make(map[[26]int]bool)
	for _, a := range A {
		count := [26]int{}
		i := 0
		for i = 0; i < len(A[0]); i++ {
			count[a[i]-'a']++
			if i%2 == 0 {
				count[a[i]-'a'] += 1000
			}
		}
		groups[count] = true
	}
	return len(groups)
}
