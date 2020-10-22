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

// leetcode841_钥匙和房间
var visited []bool
var total int

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	total = 0
	visited = make([]bool, n)
	dfs(rooms, 0)
	return total == n
}

func dfs(rooms [][]int, start int) {
	visited[start] = true
	total++
	for _, room := range rooms[start] {
		if visited[room] == false {
			dfs(rooms, room)
		}
	}
}
