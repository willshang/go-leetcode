package main

import "fmt"

func main() {
	arr := [][]int{
		{1},
		{2},
		{3},
		{},
	}
	fmt.Println(canVisitAllRooms(arr))
}

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	total := 0
	visited := make([]bool, n)
	visited[0] = true
	queue := make([]int, 0)
	queue = append(queue, 0)
	for len(queue) > 0 {
		start := queue[0]
		queue = queue[1:]
		total++
		for _, room := range rooms[start] {
			if visited[room] == false {
				visited[room] = true
				queue = append(queue, room)
			}
		}
	}
	return total == n
}
