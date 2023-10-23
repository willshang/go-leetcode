package main

func main() {

}

// leetcode1837_K进制表示下的各位数字总和
func sumBase(n int, k int) int {
	res := 0
	for n > 0 {
		res = res + n%k
		n = n / k
	}
	return res
}
