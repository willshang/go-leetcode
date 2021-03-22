package main

func main() {

}

// leetcode1780_判断一个数字是否可以表示成三的幂的和
func checkPowersOfThree(n int) bool {
	arr := make([]int, 0)
	arr = append(arr, 1)
	sum := 1
	for i := 0; i < 15; i++ {
		sum = sum * 3
		arr = append(arr, sum)
	}
	for i := len(arr) - 1; i >= 0; i-- {
		if n > arr[i] {
			n = n - arr[i]
		} else if n == arr[i] {
			return true
		}
	}
	return false
}
