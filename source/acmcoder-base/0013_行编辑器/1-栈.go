package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var str string
		fmt.Scan(&str)
		res := getResult(str)
		fmt.Println(res)
	}
}

func getResult(str string) string {
	res := make([]byte, 0)
	for i := 0; i < len(str); i++ {
		if str[i] == '@' {
			res = make([]byte, 0)
		} else if str[i] == '#' {
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
		} else {
			res = append(res, str[i])
		}
	}
	return string(res)
}
