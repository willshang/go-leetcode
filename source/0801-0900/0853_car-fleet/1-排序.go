package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(carFleet(12, []int{10,8,0,5,3},[]int{2,4,1,1,3}))
	fmt.Println(carFleet(10, []int{0, 4, 2}, []int{2, 1, 3}))
}

// leetcode853_车队
type Node struct {
	Position int
	Left     float64
}

func carFleet(target int, position []int, speed []int) int {
	if len(position) == 0 {
		return 0
	}
	arr := make([]Node, 0)
	for i := 0; i < len(position); i++ {
		arr = append(arr, Node{
			Position: position[i],
			Left:     float64(target-position[i]) / float64(speed[i]),
		})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Position > arr[j].Position
	})
	res := 1
	prev := arr[0].Left
	for i := 1; i < len(arr); i++ {
		if prev < arr[i].Left {
			res++
			prev = arr[i].Left
		}
	}
	return res
}
