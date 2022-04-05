package main

func main() {

}

// leetcode2177_找到和为给定整数的三个连续整数
func sumOfThree(num int64) []int64 {
	if num%3 == 0 {
		target := num / 3
		return []int64{target - 1, target, target + 1}
	}
	return nil
}
