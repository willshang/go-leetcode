package main

import "fmt"

func main() {
	//fmt.Println(pushDominoes(".L.R...LR..L.."))
	fmt.Println(pushDominoes(".......LR....."))
}

// leetcode838_推多米诺
func pushDominoes(dominoes string) string {
	n := len(dominoes)
	arr := append([]byte{'L'}, dominoes...)
	arr = append(arr, 'R') // 前面补1个L，后面补1个R
	temp := make([]byte, n+2)
	for string(temp) != string(arr) { // 模拟每一秒：当没有变化的时候退出
		copy(temp, arr) // 存储之前1秒的情况
		for i := 1; i <= n; i++ {
			if arr[i] == '.' { // 当前位置为.进行判断2边情况
				// 讨论2种变的情况
				if temp[i-1] == 'R' && temp[i+1] != 'L' {
					arr[i] = 'R' // 1、左边向右边倒，右边为R或者.
				} else if temp[i+1] == 'L' && temp[i-1] != 'R' {
					arr[i] = 'L' // 2、右边向左边倒，左边为L或者.
				}
			}
		}
	}
	return string(arr[1 : n+1])
}
