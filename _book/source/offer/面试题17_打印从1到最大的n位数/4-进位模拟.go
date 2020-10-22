package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(printNumbers(3))
}

func printNumbers(n int) []int {
	res := make([]int, 0)
	if n <= 0 {
		return nil
	}
	temp := make([]string, 0)
	arr := make([]rune, n)
	for i := 0; i < len(arr); i++ {
		arr[i] = '0'
	}
	for !increment(arr) {
		if printNum(arr) != "" {
			temp = append(temp, printNum(arr))
		}
	}
	for i := 0; i < len(temp); i++ {
		value, _ := strconv.Atoi(temp[i])
		res = append(res, value)
	}
	return res
}

func increment(arr []rune) bool {
	isOverflow := false
	nTakeOver := 0
	for i := len(arr) - 1; i >= 0; i-- {
		sum := int(arr[i]-'0') + nTakeOver
		if i == len(arr)-1 {
			sum++
		}
		if sum >= 10 {
			if i == 0 {
				isOverflow = true
			} else {
				sum = sum - 10
				nTakeOver = 1
				arr[i] = rune('0' + sum)
			}
		} else {
			arr[i] = rune('0' + sum)
			break
		}
	}
	return isOverflow
}

func printNum(arr []rune) string {
	res := ""
	isBeginning := true
	for i := 0; i < len(arr); i++ {
		if isBeginning && arr[i] != '0' {
			isBeginning = false
		}
		if !isBeginning {
			res = res + string(arr[i])
		}
	}
	return res
}
