package main

import "strings"

func main() {

}

func winnerOfGame(colors string) bool {
	arrA := strings.Split(colors, "B")
	arrB := strings.Split(colors, "A")
	countA, countB := 0, 0
	for i := 0; i < len(arrA); i++ {
		if len(arrA[i]) >= 3 {
			countA = countA + len(arrA[i]) - 2
		}
	}
	for i := 0; i < len(arrB); i++ {
		if len(arrB[i]) >= 3 {
			countB = countB + len(arrB[i]) - 2
		}
	}
	if countA > countB {
		return true
	}
	return false
}
