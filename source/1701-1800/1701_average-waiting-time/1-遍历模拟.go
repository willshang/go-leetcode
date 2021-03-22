package main

func main() {

}

// leetcode1701_平均等待时间
func averageWaitingTime(customers [][]int) float64 {
	sum := 0
	cur := customers[0][0]
	for i := 0; i < len(customers); i++ {
		if cur < customers[i][0] {
			cur = customers[i][0] + customers[i][1]
		} else {
			cur = cur + customers[i][1] // 做菜时间
		}
		wait := cur - customers[i][0] // 等待时间
		sum = sum + wait
	}
	return float64(sum) / float64(len(customers))
}
