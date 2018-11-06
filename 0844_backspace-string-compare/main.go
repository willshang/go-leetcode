package main

import (
	"fmt"
	"strings"
)

func main() {
	S := "a##c"
	T := "#a#c"
	fmt.Println(backspaceCompare(S,T))
}
func backspaceCompare(S string, T string) bool {
	s := check(S)
	t := check(T)
	return s == t
}

func check(str string)string  {
	res := []string{}
	for _,v:= range str{
		if string(v) == "#"{
			if len(res) != 0{
				res = res[:len(res)-1]
			}
		}else {
			res = append(res,string(v))
		}
	}
	return strings.Join(res,"")
}