package main

import "strings"

func main() {

}

var m map[byte]bool = map[byte]bool{
	'!': true,
	'.': true,
	',': true,
}

func countValidWords(sentence string) int {
	res := 0
	arr := strings.Fields(sentence)
	for i := 0; i < len(arr); i++ {
		flag := true
		countA := 0
		countB := 0
		for j := 0; j < len(arr[i]); j++ {
			if '0' <= arr[i][j] && arr[i][j] <= '9' {
				flag = false
				break
			}
			if m[arr[i][j]] == true {
				countA++
				if j != len(arr[i])-1 {
					flag = false
					break
				}
			}
			if arr[i][j] == '-' {
				countB++
				if j == 0 || j == len(arr[i])-1 {
					flag = false
					break
				}
				if (('a' <= arr[i][j-1] && arr[i][j-1] <= 'z') &&
					('a' <= arr[i][j+1] && arr[i][j+1] <= 'z')) == false {
					flag = false
					break
				}
			}
		}
		if flag == true && countA <= 1 && countB <= 1 {
			res++
		}
	}
	return res
}
