package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	fmt.Println(calPoints(arr))
}
func calPoints(ops []string) int {
	stacks := make([]int, 0, len(ops))
	for i := range ops {
		switch ops[i] {
		case "+":
			r1 := stacks[len(stacks)-1]
			r2 := stacks[len(stacks)-2]
			stacks = append(stacks, r1+r2)
		case "D":
			r1 := stacks[len(stacks)-1]
			stacks = append(stacks, 2*r1)
		case "C":
			stacks = stacks[:len(stacks)-1]
		default:
			tempInt, _ := strconv.Atoi(ops[i])
			stacks = append(stacks, tempInt)
		}
	}
	res := 0
	for _, p := range stacks {
		res = res + p
	}
	return res
}
