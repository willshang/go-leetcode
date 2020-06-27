package main

func main() {

}

func findSolution(customFunction func(int, int) int, z int) [][]int {
	res := make([][]int, 0)
	for i := 1; i <= 1000; i++ {
		for j := 1; j < 1000; j++ {
			if customFunction(i, j) == z {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}
