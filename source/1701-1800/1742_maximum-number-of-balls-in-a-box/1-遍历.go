package main

func main() {

}

// leetcode1742_盒子中小球的最大数量
func countBalls(lowLimit int, highLimit int) int {
	res := 0
	m := make(map[int]int)
	for i := lowLimit; i <= highLimit; i++ {
		sum := getSum(i)
		m[sum]++
		if m[sum] > res {
			res = m[sum]
		}
	}
	return res
}

func getSum(i int) int {
	sum := 0
	for i > 0 {
		sum = sum + i%10
		i = i / 10
	}
	return sum
}
