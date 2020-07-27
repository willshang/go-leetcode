package main

import "fmt"

func main() {
	fmt.Println(movingCount(1, 2, 1))
}

func movingCount(m int, n int, k int) int {
	if k < 0 || m <= 0 || n <= 0 {
		return 0
	}
	res := 0
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	visited[0][0] = true
	queue := make([][2]int, 0)
	queue = append(queue, [2]int{0, 0})
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		x := node[0]
		y := node[1]
		if getDigiSum(x)+getDigiSum(y) <= k {
			res++
		}
		if x+1 < m && getDigiSum(x+1)+getDigiSum(y) <= k && visited[x+1][y] == false {
			queue = append(queue, [2]int{x + 1, y})
			visited[x+1][y] = true
		}
		if x-1 >= 0 && getDigiSum(x-1)+getDigiSum(y) <= k && visited[x-1][y] == false {
			queue = append(queue, [2]int{x - 1, y})
			visited[x-1][y] = true
		}
		if y+1 < n && getDigiSum(x)+getDigiSum(y+1) <= k && visited[x][y+1] == false {
			queue = append(queue, [2]int{x, y + 1})
			visited[x][y+1] = true
		}
		if y-1 >= 0 && getDigiSum(x)+getDigiSum(y-1) <= k && visited[x][y-1] == false {
			queue = append(queue, [2]int{x, y - 1})
			visited[x][y-1] = true
		}
	}
	return res
}

func getDigiSum(num int) int {
	sum := 0
	for num > 0 {
		sum = sum + num%10
		num = num / 10
	}
	return sum
}
