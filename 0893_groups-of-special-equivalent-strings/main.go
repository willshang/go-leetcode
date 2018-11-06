package main

import "fmt"

func main() {
	arr := []string{"abc","acb","bac","bca","cab","cba"}
	fmt.Println(numSpecialEquivGroups(arr))
}

//map[1 1 1000 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]:true
//   [1000 1 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]:true
//   [1 1000 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]:true
func numSpecialEquivGroups(A []string) int {
	strSize := len(A[0])
	groups := make(map[[26]int]bool,len(A))

	for _, a := range A{
		count := [26]int{}
		i := 0
		for i = 0; i+1 < strSize; i += 2{
			count[a[i]-'a']++
			count[a[i+1]-'a'] += 1000
		}

		if i < strSize{
			count[a[i]-'a']++
		}
		groups[count] = true
	}
	return len(groups)
}