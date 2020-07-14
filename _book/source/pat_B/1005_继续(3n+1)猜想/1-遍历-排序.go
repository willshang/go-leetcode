package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	arr := make([]int, 10001)
	maps := make(map[int]int, n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scanf("%d", &num)
		maps[num] = num // map保存输入的值
		if num == 1 {
			continue
		}
		for {
			if num%2 != 0 {
				num = 3*num + 1
			}
			num = num / 2
			if num == 1 {
				break
			}
			//fmt.Println("num: 出现:\t",num,arr[num] == 1)
			arr[num] = 1
		}
	}
	tempArr := make([]int, 0)
	for k := range maps {
		tempArr = append(tempArr, maps[k])
	}
	sort.Ints(tempArr)
	var flag = false
	for i := len(tempArr) - 1; i >= 0; i-- {
		if arr[tempArr[i]] == 0 {
			if flag == true {
				fmt.Print(" ")
			}
			fmt.Print(tempArr[i])
			flag = true
		}
	}
}
