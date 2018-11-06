package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"}))
	fmt.Println(longestCommonPrefix([]string{"a"}))
}

//执行用时：8 ms
/*func longestCommonPrefix(strs []string) string {
	if len(strs) == 0{
		return ""
	}
	if len(strs) == 1{
		return strs[0]
	}

	short := strs[0]
	for _, s := range strs{
		if len(short) > len(s){
			short = s
		}
	}

	for i, _ := range short{
		shortest := short[:i+1]
		for _,str := range strs{
			if strings.Index(str,shortest) != 0{
				return short[:i]
			}
		}
	}
	return short
}*/

//执行用时：0 ms
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0{
		return ""
	}
	if len(strs) == 1{
		return strs[0]
	}

	length := 0

	for i := 0; i < len(strs[0]); i++{
		char := strs[0][i]
		for j := 1; j < len(strs); j++{
			if i >= len(strs[j]) || char != strs[j][i]{
				return strs[0][:length]
			}
		}
		length++
	}
	return strs[0][:length]
}