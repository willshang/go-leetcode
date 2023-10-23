package main

func main() {

}

// leetcode2310_个位数字为K的整数之和
func minimumNumbers(num int, k int) int {
	if num == 0 {
		return 0
	}
	for i := 1; i <= 10; i++ {
		// 只考虑个位数
		if i*k <= num && (i*k)%10 == num%10 {
			return i
		}
	}
	return -1
}
