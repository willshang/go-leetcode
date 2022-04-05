package main

func main() {

}

// leetcode2180_统计各位数字之和为偶数的整数个数
func countEven(num int) int {
	res := 0
	for i := 2; i <= num; i++ {
		sum := 0
		temp := i
		for temp > 0 {
			sum = sum + temp%10
			temp = temp / 10
		}
		if sum%2 == 0 {
			res++
		}
	}
	return res
}
