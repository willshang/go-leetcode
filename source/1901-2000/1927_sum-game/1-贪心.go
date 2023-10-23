package main

import "fmt"

func main() {
	fmt.Println(sumGame("5023"))
}

// leetcode1927_求和游戏
func sumGame(num string) bool {
	n := len(num)
	leftSum, rightSum := 0, 0
	leftCount, rightCount := 0, 0
	for i := 0; i < n; i++ {
		if i < n/2 {
			if num[i] == '?' {
				leftCount++
			} else {
				leftSum = leftSum + int(num[i]-'0')
			}
		} else {
			if num[i] == '?' {
				rightCount++
			} else {
				rightSum = rightSum + int(num[i]-'0')
			}
		}
	}
	if (leftCount+rightCount)%2 == 1 { // ?总数为奇数个
		return true
	}
	if leftSum-rightSum != (rightCount-leftCount)*9/2 {
		return true
	}
	return false
}
