package main

func main() {

}

// leetcode2022_将一维数组转变成二维数组
func construct2DArray(original []int, m int, n int) [][]int {
	total := len(original)
	if n*m != total {
		return nil
	}
	res := make([][]int, 0)
	for i := 0; i < total; i = i + n {
		res = append(res, original[i:i+n])
	}
	return res
}
