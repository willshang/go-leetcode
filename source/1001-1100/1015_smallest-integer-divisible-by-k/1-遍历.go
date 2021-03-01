package main

func main() {

}

// leetcode1015_可被K整除的最小整数
func smallestRepunitDivByK(K int) int {
	if K%2 == 0 || K%5 == 0 {
		return -1
	}
	res := 1
	num := 1
	for {
		if num%K == 0 {
			return res
		}
		num = num % K // (n*10+1)%K = ((n%K)*10+1)%K
		num = 10*num + 1
		res++
	}
	return -1
}
