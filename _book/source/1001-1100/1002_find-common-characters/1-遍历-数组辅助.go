package main

import "fmt"

func main() {
	fmt.Println(commonChars([]string{"bella", "label", "roller"}))
}

func commonChars(A []string) []string {
	arr := [26]int{}
	for _, v := range A[0] {
		arr[v-'a']++
	}
	for i := 1; i < len(A); i++ {
		temp := [26]int{}
		for _, v := range A[i] {
			temp[v-'a']++
		}
		for i := 0; i < len(arr); i++ {
			arr[i] = min(arr[i], temp[i])
		}
	}
	res := make([]string, 0)
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			for j := 0; j < arr[i]; j++ {
				res = append(res, string('a'+i))
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
