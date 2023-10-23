package main

func main() {

}

// leetcode2059_转化数字的最小运算数
func minimumOperations(nums []int, start int, goal int) int {
	m := make(map[int]bool)
	m[start] = true
	queue := make([]int, 0)
	queue = append(queue, start)
	count := 1
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			for j := 0; j < len(nums); j++ {
				temp := []int{queue[i] + nums[j], queue[i] - nums[j], queue[i] ^ nums[j]}
				for k := 0; k < len(temp); k++ {
					if temp[k] == goal {
						return count
					}
					if 0 <= temp[k] && temp[k] <= 1000 && m[temp[k]] == false {
						m[temp[k]] = true
						queue = append(queue, temp[k])
					}
				}
			}
		}
		count++
		queue = queue[length:]
	}
	return -1
}
