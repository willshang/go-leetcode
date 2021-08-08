package main

func main() {

}

// leetcode1953_你可以工作的最大周数
func numberOfWeeks(milestones []int) int64 {
	var maxCount int64
	var sum int64
	for i := 0; i < len(milestones); i++ {
		if int64(milestones[i]) > maxCount {
			maxCount = int64(milestones[i])
		}
		sum = sum + int64(milestones[i])
	}
	if maxCount > (sum+1)/2 {
		return (sum-maxCount)*2 + 1
	}
	return sum
}
