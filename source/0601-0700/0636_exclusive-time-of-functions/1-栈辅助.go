package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(exclusiveTime(1, []string{
		"0:start:0", "0:start:2", "0:end:5", "0:start:6", "0:end:6", "0:end:7",
	}))
}

type Node struct {
	Id        int
	StartTime int
	Wait      int
}

func exclusiveTime(n int, logs []string) []int {
	res := make([]int, n)
	stack := make([]Node, 0)
	for i := 0; i < len(logs); i++ {
		arr := strings.Split(logs[i], ":")
		id, _ := strconv.Atoi(arr[0])
		if arr[1] == "start" {
			start, _ := strconv.Atoi(arr[2])
			stack = append(stack, Node{
				Id:        id,
				StartTime: start,
				Wait:      0,
			})
		} else {
			end, _ := strconv.Atoi(arr[2])
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			total := end - node.StartTime + 1 - node.Wait
			res[node.Id] = res[node.Id] + total
			if len(stack) > 0 {
				wait := end - node.StartTime + 1
				stack[len(stack)-1].Wait = stack[len(stack)-1].Wait + wait
			}
		}
	}
	return res
}
