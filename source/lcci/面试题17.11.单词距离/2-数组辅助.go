package main

import "fmt"

func main() {
	arr := []string{"I", "am", "a", "student", "from", "a", "university", "in", "a", "city"}
	fmt.Println(findClosest(arr, "a", "student"))
}

func findClosest(words []string, word1 string, word2 string) int {
	res := len(words) - 1
	arrA, arrB := make([]int, 0), make([]int, 0)
	for i := 0; i < len(words); i++ {
		if words[i] == word1 {
			arrA = append(arrA, i)
		}
		if words[i] == word2 {
			arrB = append(arrB, i)
		}
	}
	i, j := 0, 0
	for i < len(arrA) && j < len(arrB) {
		if abs(arrA[i], arrB[j]) < res {
			res = abs(arrA[i], arrB[j])
		}
		if arrA[i] < arrB[j] {
			i++
		} else {
			j++
		}
	}
	return res
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
