package main

func main() {

}

func countPairs(nums []int, k int) int64 {
	res := int64(0)
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[gcd(k, nums[i])]++ // k和nums[i]的最大公约数
	}
	for k1, v1 := range m {
		for k2, v2 := range m {
			// 核心：a*b是k的倍数 => gcd(a,k)*gcd(b,k)是k的倍数
			if k1*k2%k == 0 {
				res = res + int64(v1)*int64(v2) // 组合数
			}
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i]*nums[i]%k == 0 { // 本身x本身：不满足要求
			res--
		}
	}
	return res / 2 // 要满足有序要求
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
