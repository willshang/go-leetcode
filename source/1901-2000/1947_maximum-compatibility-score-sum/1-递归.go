package main

func main() {

}

var arr [][]int
var res int

func maxCompatibilitySum(students [][]int, mentors [][]int) int {
	n := len(students)
	arr = make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
		for j := 0; j < n; j++ {
			arr[i][j] = calculate(students[i], mentors[j])
		}
	}
	res = 0
	for i := 0; i < n; i++ {
		visited1, visited2 := make([]int, n), make([]int, n)
		visited1[0], visited2[i] = 1, 1
		dfs(n, visited1, visited2, arr[0][i])
	}
	return res
}

func dfs(n int, visited1, visited2 []int, sum int) {
	if sum > res {
		res = sum
	}
	index := -1
	for i := 0; i < n; i++ {
		if visited1[i] == 0 {
			index = i
			break
		}
	}
	if index == -1 {
		return
	}
	for i := 0; i < n; i++ {
		if visited2[i] == 1 {
			continue
		}
		temp1 := make([]int, n)
		temp2 := make([]int, n)
		copy(temp1, visited1)
		copy(temp2, visited2)
		temp1[index] = 1
		temp2[i] = 1
		dfs(n, temp1, temp2, sum+arr[index][i])
	}
}

func calculate(a, b []int) int {
	res := 0
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			res++
		}
	}
	return res
}
