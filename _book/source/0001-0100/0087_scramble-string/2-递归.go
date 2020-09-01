package main

import "fmt"

func main() {
	fmt.Println(isScramble("great", "rgeat"))
}

func isScramble(s1 string, s2 string) bool {
	return dfs([]byte(s1), []byte(s2))
}

func dfs(arr1, arr2 []byte) bool {
	if compare(arr1, arr2) == false {
		return false
	}
	if len(arr1) <= 2 {
		return (len(arr1) == 2 && ((arr1[0] == arr2[0] && arr1[1] == arr2[1]) ||
			(arr1[0] == arr2[1] && arr1[1] == arr2[0]))) ||
			(len(arr1) == 1 && arr1[0] == arr2[0])
	}
	for i := 1; i < len(arr1); i++ {
		leftA, rightA := arr1[:i], arr1[i:]
		leftB, rightB := arr2[:i], arr2[i:]
		LB, RB := arr2[len(arr1)-i:], arr2[:len(arr1)-i]
		if (dfs(leftA, leftB) && dfs(rightA, rightB)) || (dfs(leftA, LB) && dfs(rightA, RB)) {
			return true
		}
	}
	return false
}

func compare(arr1, arr2 []byte) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	arrA := make([]byte, 26)
	arrB := make([]byte, 26)
	for i := 0; i < len(arr1); i++ {
		arrA[arr1[i]-'a']++
		arrB[arr2[i]-'a']++
	}
	for i := 0; i < len(arrA); i++ {
		if arrA[i] != arrB[i] {
			return false
		}
	}
	return true
}
