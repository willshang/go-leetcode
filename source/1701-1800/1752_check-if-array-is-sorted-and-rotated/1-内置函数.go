package main

import (
	"sort"
	"strconv"
	"strings"
)

func main() {

}

func check(nums []int) bool {
	temp := make([]int, len(nums))
	copy(temp, nums)
	sort.Ints(temp)
	nums = append(nums, nums...)
	a := change(temp)
	b := change(nums)
	if strings.Contains(b, a) {
		return true
	}
	return false
}

func change(arr []int) string {
	res := ""
	for i := 0; i < len(arr); i++ {
		res = res + strconv.Itoa(arr[i]) + ","
	}
	res = strings.TrimRight(res, ",")
	return res
}
