package main

func main() {

}

func construct2DArray(original []int, m int, n int) [][]int {
	total := len(original)
	if n*m != total {
		return nil
	}
	res := make([][]int, 0)
	index := 0
	for i := 0; i < m; i++ {
		temp := make([]int, 0)
		for j := 0; j < n; j++ {
			temp = append(temp, original[index])
			index++
		}
		res = append(res, temp)
	}
	return res
}
