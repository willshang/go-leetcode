package main

import (
	"fmt"
	"strings"
)

func main() {
	strs := []string{"Hello", "Alaska", "Dad", "Peace"}
	fmt.Println(findWords(strs))
}
func findWords(words []string) []string {
	m := make(map[byte]int)
	m['q'] = 1
	m['w'] = 1
	m['e'] = 1
	m['r'] = 1
	m['t'] = 1
	m['y'] = 1
	m['u'] = 1
	m['i'] = 1
	m['o'] = 1
	m['p'] = 1
	m['a'] = 2
	m['s'] = 2
	m['d'] = 2
	m['f'] = 2
	m['g'] = 2
	m['h'] = 2
	m['j'] = 2
	m['k'] = 2
	m['l'] = 2
	m['z'] = 3
	m['x'] = 3
	m['c'] = 3
	m['v'] = 3
	m['b'] = 3
	m['n'] = 3
	m['m'] = 3

	res := []string{}

	for i := 0; i < len(words); i++ {
		b := []byte(strings.ToLower(words[i]))
		level := m[b[0]]
		flag := true
		for j := 1; j < len(b); j++ {
			if m[b[j]] != level {
				flag = false
				break
			}
		}
		if flag {
			res = append(res, words[i])
		}
	}
	return res
}

/*
var qRow = map[byte]bool{
	'q' : true,
	'w' : true,
	'e' : true,
	'r' : true,
	't' : true,
	'y' : true,
	'u' : true,
	'i' : true,
	'o' : true,
	'p' : true,
}

var aRow = map[byte]bool{
	'a' : true,
	's' : true,
	'd' : true,
	'f' : true,
	'g' : true,
	'h' : true,
	'j' : true,
	'k' : true,
	'l' : true,
}

var zRow = map[byte]bool{
	'z' : true,
	'x' : true,
	'c' : true,
	'v' : true,
	'b' : true,
	'n' : true,
	'm' : true,
}
func findWords(words []string) []string {
	res := make([]string, 0, len(words))

	for _, word := range words{
		w := strings.ToLower(word)
		if isAllIn(w, qRow) || isAllIn(w, aRow) || isAllIn(w,zRow){
			res = append(res,word)
		}
	}
	return res
}

func isAllIn(s string, Row map[byte]bool)bool  {
	for i := range s{
		if !Row[s[i]]{
			return false
		}
	}
	return true
}*/
