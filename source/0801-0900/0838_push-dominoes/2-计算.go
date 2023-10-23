package main

import "fmt"

func main() {
	//fmt.Println(pushDominoes(".L.R...LR..L.."))
	fmt.Println(pushDominoes(".......LR....."))
}

func pushDominoes(dominoes string) string {
	n := len(dominoes)
	arr := make([]int, n)
	value := 0
	for i := 0; i < n; i++ { // 1、从左往右
		// 计算左边的受力，R=n，L=0，.的时候随距离减1
		if dominoes[i] == 'R' {
			value = n
		} else if dominoes[i] == 'L' {
			value = 0
		} else {
			value = max(0, value-1)
		}
		arr[i] = arr[i] + value
	}
	value = 0
	for i := n - 1; i >= 0; i-- { // 2、从右往左
		// 计算右边的受力，R=0，L=R，.的时候随距离减1
		if dominoes[i] == 'L' {
			value = n
		} else if dominoes[i] == 'R' {
			value = 0
		} else {
			value = max(0, value-1)
		}
		arr[i] = arr[i] - value
	}
	res := []byte(dominoes)
	for i := 0; i < n; i++ {
		// 哪边受力大，结果随哪边
		if arr[i] > 0 {
			res[i] = 'R'
		} else if arr[i] < 0 {
			res[i] = 'L'
		} else {
			res[i] = '.'
		}
	}
	return string(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
