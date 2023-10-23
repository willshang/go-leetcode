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

func getResult(str string) int {
	arr := [12]int{}
	for i := 0; i < len(str); i++ {
		arr[i] = int(str[i] - 'a')
	}
	res := 1
	for i := 0; i < len(arr); i++ {

	}
	return res
}

func getValue() {

}
