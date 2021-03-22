package main

func main() {

}

// leetcode1310_子数组异或查询
func xorQueries(arr []int, queries [][]int) []int {
	temp := make([]int, len(arr)+1)
	for i := 0; i < len(arr); i++ {
		temp[i+1] = temp[i] ^ arr[i]
	}
	res := make([]int, 0)
	for i := 0; i < len(queries); i++ {
		a, b := queries[i][0], queries[i][1]+1
		res = append(res, temp[a]^temp[b])
	}
	return res
}
