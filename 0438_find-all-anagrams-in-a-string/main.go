package main

import "fmt"

func main() {
	s := "abcdddacbeee"
	p := "bcd"
	fmt.Println(findAnagrams(s,p))
}
func findAnagrams(s string, p string) []int {
	ret := []int{}
	var target, window  [26]int
	for i := 0; i < len(p); i++{
		target[p[i]-'a']++
	}

	check := func(i int) {
		if window == target{
			ret = append(ret,i)
		}
	}

	for i := 0; i < len(s); i++{
		window[s[i]-'a']++
		if i == len(p)-1{
			check(0)
		}else if len(p) <= i{
			window[s[i-len(p)] -'a']--
			check(i-len(p)+1)
		}
	}
	return ret
}