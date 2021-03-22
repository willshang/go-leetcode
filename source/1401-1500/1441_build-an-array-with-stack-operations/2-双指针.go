package main

import "fmt"

func main() {
	fmt.Println(buildArray([]int{1, 3}, 3))
}

func buildArray(target []int, n int) []string {
	res := make([]string, 0)
	j := 1
	for i := 0; i < len(target); i++ {
		for ; j < target[i]; j++ {
			res = append(res, "Push")
			res = append(res, "Pop")
		}
		res = append(res, "Push")
		j++
	}
	return res
}
