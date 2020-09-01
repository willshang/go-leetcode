package main

import "fmt"

func main() {
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	return dfs(s1, s2, s3, 0, 0, 0)
}

func dfs(s1, s2, s3 string, i, j, k int) bool {
	if k == len(s3) && i == len(s1) && j == len(s2) {
		return true
	}
	if k >= len(s3) {
		return false
	}
	if i < len(s1) {
		if s1[i] == s3[k] {
			if dfs(s1, s2, s3, i+1, j, k+1) {
				return true
			}
		}
	}
	if j < len(s2) {
		if s2[j] == s3[k] {
			if dfs(s1, s2, s3, i, j+1, k+1) {
				return true
			}
		}
	}
	return false
}
