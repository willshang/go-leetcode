package main

func main() {

}

// leetcode335_路径交叉
func isSelfCrossing(distance []int) bool {
	n := len(distance)
	if n < 4 {
		return false
	}
	for i := 3; i < n; i++ {
		// 长度为4相交
		if i >= 3 && distance[i] >= distance[i-2] &&
			distance[i-1] <= distance[i-3] {
			return true
		}
		// 长度为5相交
		if i >= 4 && distance[i]+distance[i-4] >= distance[i-2] &&
			distance[i-1] == distance[i-3] {
			return true
		}
		// 长度为6相交
		if i >= 5 && distance[i]+distance[i-4] >= distance[i-2] &&
			distance[i-1]+distance[i-5] >= distance[i-3] &&
			distance[i-2] > distance[i-4] &&
			distance[i-1] < distance[i-3] {
			return true
		}
	}
	return false
}
