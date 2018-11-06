package main

import (
	"fmt"
	"strings"
)

func main() {
	pattern := "abba"
	str := "dog cat cat dog"
	fmt.Println(wordPattern(pattern,str))
}

func wordPattern(pattern string, str string) bool {
	pa := strings.Split(pattern,"")
	sa := strings.Split(str," ")
	if len(pa) != len(sa){
		return false
	}
	length := len(pa)
	m := make(map[string]string,length)
	n := make(map[string]string,length)

	for i := 0; i < length; i++{
		if w, ok := m[pa[i]]; ok{
			if y, ok := n[sa[i]]; ok{
				if w != sa[i] && y != pa[i]{
					return false
				}
			}else {
				return false
			}
		}else {
			if _,ok := m[pa[i]];!ok && n[sa[i]] != ""{
				return false
			}
			m[pa[i]] = sa[i]
			n[sa[i]] = pa[i]
		}
	}
	return true
}
/*func wordPattern(pattern string, str string) bool {
	pa := strings.Split(pattern,"")
	sa := strings.Split(str," ")
	if len(pa) != len(sa){
		return false
	}
	return isMatch(pa,sa) && isMatch(sa,pa)
}

func isMatch(pa, sa []string) bool {
	length := len(pa)
	m := make(map[string]string,length)

	for i := 0; i < length; i++{
		if w, ok := m[pa[i]]; ok{
			if w != sa[i]{
				return false
			}
		}else {
			m[pa[i]] = sa[i]
		}
	}
	return true
}*/