package main

import "fmt"

func main() {
	fmt.Println(permutation("qwe"))
}

func permutation(S string) []string {
	if len(S) == 1 {
		return []string{S}
	}
	res := make([]string, 0)
	for i := 0; i < len(S); i++ {
		str := S[:i] + S[i+1:]
		arr := permutation(str)
		for _, v := range arr {
			res = append(res, v+string(S[i]))
		}
	}
	return res
}
