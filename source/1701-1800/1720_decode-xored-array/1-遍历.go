package main

func main() {

}

// leetcode1720_解码异或后的数组
func decode(encoded []int, first int) []int {
	res := make([]int, 0)
	res = append(res, first)
	for i := 0; i < len(encoded); i++ {
		res = append(res, encoded[i]^res[len(res)-1])
	}
	return res
}
