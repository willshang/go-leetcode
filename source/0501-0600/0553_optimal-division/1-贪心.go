package main

import (
	"strconv"
	"strings"
)

func main() {

}

// leetcode553_最优除法
func optimalDivision(nums []int) string {
	res := make([]string, 0)
	for i := 0; i < len(nums); i++ {
		res = append(res, strconv.Itoa(nums[i]))
	}
	if len(res) < 3 {
		return strings.Join(res, "/")
	}
	return res[0] + "/(" + strings.Join(res[1:], "/") + ")"
}
