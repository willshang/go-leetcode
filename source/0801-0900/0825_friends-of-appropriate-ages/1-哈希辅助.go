package main

func main() {

}

// leetcode825_适龄的朋友
func numFriendRequests(ages []int) int {
	res := 0
	m := make(map[int]int)
	for i := 0; i < len(ages); i++ {
		m[ages[i]]++
	}
	for k, v := range m {
		for key, value := range m {
			if key <= k/2+7 || key > k || (key > 100 && k < 100) {
				continue
			} else if k == key {
				res = res + v*(value-1)
			} else {
				res = res + v*value
			}
		}
	}
	return res
}
