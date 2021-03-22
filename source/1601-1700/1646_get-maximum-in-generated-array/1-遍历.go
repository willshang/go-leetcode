package main

func main() {

}

// leetcode1646_获取生成数组中的最大值
func getMaximumGenerated(n int) int {
	arr := make([]int, n+3)
	arr[0] = 0
	arr[1] = 1
	res := 0
	for i := 0; i <= n; i++ {
		if i%2 == 0 {
			arr[i] = arr[i/2]
		} else {
			arr[i] = arr[i/2] + arr[(i+1)/2]
		}
		if arr[i] > res {
			res = arr[i]
		}
	}
	return res
}
