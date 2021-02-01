package main

func main() {

}

// leetcode869_重新排序得到2的幂
func reorderedPowerOf2(N int) bool {
	arr := getCount(N)
	for i := 0; i < 31; i++ {
		if arr == getCount(1<<i) {
			return true
		}
	}
	return false
}

func getCount(n int) [10]int {
	arr := [10]int{}
	for n > 0 {
		arr[n%10]++
		n = n / 10
	}
	return arr
}
