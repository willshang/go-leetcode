package main

import "fmt"

func main() {
	fmt.Println(countSubstrings("abc"))
}

func countSubstrings(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	str := add(s)
	length := len(str)
	res := 0
	for i := 0; i < length; i++ {
		curLength := search(str, i)
		res = res + curLength/2 + curLength%2
	}
	return res
}

func add(s string) string {
	var res []rune
	for _, v := range s {
		res = append(res, '#')
		res = append(res, v)
	}
	res = append(res, '#')
	return string(res)
}

func search(s string, center int) int {
	i := center - 1
	j := center + 1
	step := 0
	for ; i >= 0 && j < len(s) && s[i] == s[j]; i, j = i-1, j+1 {
		step++
	}
	return step
}
