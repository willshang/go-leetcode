package main

import (
	"fmt"
	"strings"
)

func main() {
	J := "aA"
	S := "aAAbbbb"
	fmt.Println(numJewelsInStones(J,S))
}
func numJewelsInStones(J string, S string) int {
	res := 0
	for _,v := range J{
		res = res + strings.Count(S,string(v))
	}
	return res
}
/*func numJewelsInStones(J string, S string) int {
	Jewel := make(map[byte]bool,len(J))
	for i := range J{
		Jewel[J[i]] = true
	}

	res := 0
	for i := range S{
		if Jewel[S[i]]{
			res++
		}
	}
	return res
}*/