package main

import (
	"fmt"
)

func main() {
	arr := []int{1,0,0,0,1,0,1}
	fmt.Println(maxDistToClosest(arr))
	fmt.Println(maxDistToClosest([]int{1,0,0,0}))
}
func maxDistToClosest(seats []int) int {
	size := len(seats)
	var lists []int
	for i := 0; i < size;i++{
		if seats[i] == 1{
			lists = append(lists,i)
		}
	}
	max := -1
	for i := 0; i < len(lists)-1;i++{
		if abs(lists[i],lists[i+1]) > max{
			max = abs(lists[i],lists[i+1])
		}
	}
	max = max / 2

	if lists[0] > max{
		max = lists[0]
	}
	if size - lists[len(lists)-1]-1 > max{
		max = size - lists[len(lists)-1]-1
	}
	return max
}

func abs(a,b int)int  {
	if a > b {
		return a - b
	}
	return b - a
}
/*func maxDistToClosest(seats []int) int {
	size := len(seats)
	maxDis := 0

	e := 0

	for i := 0; i < size; i++{
		if e == i{
			maxDis = e
		}else {
			maxDis = max(maxDis, (e+e%2)/2)
		}
		if seats[i] == 1{
			e = 0
		}else {
			e++
		}
	}
	return max(maxDis,e)
}
func max(a,b int)int  {
	if a > b{
		return a
	}
	return b
}*/