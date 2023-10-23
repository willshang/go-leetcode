package main

import "fmt"

func main() {
	fmt.Println(countSubstrings("abc"))
}

func countSubstrings(s string) int {
	var res []rune
	res = append(res, '$')
	for _, v := range s {
		res = append(res, '#')
		res = append(res, v)
	}
	res = append(res, '#')
	res = append(res, '!')
	str := string(res)
	n := len(str) - 1
	arr := make([]int, n)
	leftMax, rightMax, result := 0, 0, 0
	for i := 1; i < n; i++ {
		if i <= rightMax {
			arr[i] = min(rightMax-i+1, arr[2*leftMax-i])
		} else {
			arr[i] = 1
		}
		for str[i+arr[i]] == str[i-arr[i]] {
			arr[i]++
		}
		if i+arr[i]-1 > rightMax {
			leftMax = i
			rightMax = i + arr[i] - 1
		}
		result = result + arr[i]/2
	}
	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
