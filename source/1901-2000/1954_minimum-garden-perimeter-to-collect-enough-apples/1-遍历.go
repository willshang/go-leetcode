package main

func main() {

}

// leetcode1954_收集足够苹果的最小花园周长
func minimumPerimeter(neededApples int64) int64 {
	var total int64
	var i int64
	var sum int64
	for i = 0; ; i = i + 2 {
		sum = 3 * i * i // 边长为i（偶数）的正方形的外围苹果总个数
		total = total + sum
		if neededApples <= total {
			return int64(i) * 4
		}
	}
	return 0
}
