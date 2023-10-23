package main

func main() {

}

// leetcode1654_到家的最少跳跃次数
func minimumJumps(forbidden []int, a int, b int, x int) int {
	m := make([]bool, 6001)
	for i := 0; i < len(forbidden); i++ {
		m[forbidden[i]] = true
	}
	queue := make([][2]int, 0)
	queue = append(queue, [2]int{0, 0})
	res := -1
	for len(queue) > 0 {
		length := len(queue)
		res++
		for i := 0; i < length; i++ {
			index, dir := queue[i][0], queue[i][1]
			if index == x {
				return res
			}
			if dir == 0 && index-b > 0 && m[index-b] == false { // 向后跳-b
				m[index-b] = true
				queue = append(queue, [2]int{index - b, 1})
			}
			if index+a < len(m) && m[index+a] == false { // 向前跳+a
				m[index+a] = true
				queue = append(queue, [2]int{index + a, 0})
			}
		}
		queue = queue[length:]
	}
	return -1
}
