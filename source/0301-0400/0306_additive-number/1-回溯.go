package main

import "fmt"

func main() {
	fmt.Println(isAdditiveNumber("112358"))
}

func isAdditiveNumber(num string) bool {
	if len(num) < 3 {
		return false
	}
	return dfs(num)
}

func dfs(str string) bool {

}
