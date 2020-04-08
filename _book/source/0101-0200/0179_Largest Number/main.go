package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	arr := []int{10, 2}
	fmt.Println(largestNumber(arr))

	arr2 := []int{3, 30, 34, 5, 9}
	fmt.Println(largestNumber(arr2))

	arr3 := []int{121, 12}
	fmt.Println(largestNumber(arr3))

	arr4 := []int{0, 0}
	fmt.Println(largestNumber(arr4))

	arr5 := []int{1, 0, 0}
	fmt.Println(largestNumber(arr5))
}

func largestNumber(nums []int) string {
	strArr := make([]string, 0)

	// 先转换为str
	for i := range nums {
		strArr = append(strArr, strconv.Itoa(nums[i]))
	}

	sort.Slice(strArr, func(i, j int) bool {
		if strings.Contains(strArr[i], strArr[j]) || strings.Contains(strArr[j], strArr[i]) {
			if strArr[i]+strArr[j] < strArr[j]+strArr[i] {
				return true
			}
			return false
		}

		return strArr[i] < strArr[j]
	})
	// fmt.Println(strArr)
	result := ""
	isZero := true
	for i := len(strArr) - 1; i >= 0; i-- {
		if i > 0 && strArr[i] != "0" {
			isZero = false
		}
		if isZero == true && i > 0 {
			continue
		}
		result = result + strArr[i]
	}
	return result
}
