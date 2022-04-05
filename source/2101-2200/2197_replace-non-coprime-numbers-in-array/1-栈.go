package main

func main() {

}

// leetcode2197_替换数组中的非互质数
func replaceNonCoprimes(nums []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		for len(res) > 0 {
			if gcd(v, res[len(res)-1]) > 1 {
				v = lcm(v, res[len(res)-1])
				res = res[:len(res)-1]
			} else {
				break
			}
		}
		res = append(res, v)
	}
	return res
}

func lcm(x, y int) int {
	return x * y / gcd(x, y)
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
