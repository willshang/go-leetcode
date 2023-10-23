package main

import "strings"

func main() {

}

func finalValueAfterOperations(operations []string) int {
	res := 0
	for i := 0; i < len(operations); i++ {
		if strings.Contains(operations[i], "+") {
			res++
		} else {
			res--
		}
	}
	return res
}
