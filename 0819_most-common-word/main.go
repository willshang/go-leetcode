package main

import (
	"fmt"
	"strings"
)

func main() {
	//paragraph := "Bob hit a ball, the hit BALL flew far after it was hit."
	paragraph := "Bob. hIt, baLl"
	banned := []string{"bob", "hit"}
	fmt.Println(mostCommonWord(paragraph,banned))
}
func mostCommonWord(paragraph string, banned []string) string {
	isBanned := make(map[string]bool)
	for _, b := range banned{
		isBanned[b] = true
	}


	chars := []string{"!","?",",","'",";","."}
	for _, c := range chars{
		paragraph = strings.Replace(paragraph,c,"",-1)
	}

	p := strings.ToLower(paragraph)
	ss := strings.Split(p," ")
	count := make(map[string]int,len(ss))
	for _, v := range ss{
		if isBanned[v]{
			continue
		}
		fmt.Println(v)
		count[v]++
	}

	res := ""
	max := 0

	for s, c := range count{
		if max < c{
			max = c
			res = s
		}
	}
	return res
}

