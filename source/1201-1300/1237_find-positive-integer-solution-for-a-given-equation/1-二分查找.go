package main

func main() {

}

func findSolution(customFunction func(int, int) int, z int) [][]int {
	res := make([][]int, 0)
	for i := 1; i <= 1000; i++ {
		left := 1
		right := 1000
		v1, v2 := customFunction(i, left), customFunction(i, right)
		if z < v1 || z > v2 {
			continue
		}
		for left <= right {
			mid := left + (right-left)/2
			v := customFunction(i, mid)
			if z == v {
				res = append(res, []int{i, mid})
				break
			} else if z > v {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return res
}
