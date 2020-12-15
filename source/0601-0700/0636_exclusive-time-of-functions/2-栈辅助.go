package main

import (
	"strconv"
	"strings"
)

func main() {

}

// leetcode636_函数的独占时间
func exclusiveTime(n int, logs []string) []int {
	res := make([]int, n)
	stack := make([]int, 0)
	var prev int
	for i := 0; i < len(logs); i++ {
		arr := strings.Split(logs[i], ":")
		id, _ := strconv.Atoi(arr[0])
		if arr[1] == "start" {
			start, _ := strconv.Atoi(arr[2])
			if len(stack) > 0 {
				lastId := stack[len(stack)-1]
				res[lastId] = res[lastId] + start - prev
			}
			stack = append(stack, id)
			prev = start
		} else {
			end, _ := strconv.Atoi(arr[2])
			lastId := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[lastId] = res[lastId] + end - prev + 1
			prev = end + 1
		}
	}
	return res
}
