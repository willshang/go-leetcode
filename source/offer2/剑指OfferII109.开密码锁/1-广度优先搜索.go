package main

import "fmt"

func main() {
	fmt.Println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0000"))
	//fmt.Println(openLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888"))
}

// 剑指OfferII109.开密码锁
func openLock(deadends []string, target string) int {
	m := make(map[string]int)
	m["0000"] = 0
	for i := 0; i < len(deadends); i++ {
		if deadends[i] == "0000" || deadends[i] == target {
			return -1
		}
		m[deadends[i]] = 0
	}
	if target == "0000" {
		return 0
	}
	queue := make([]string, 0)
	queue = append(queue, "0000")
	res := 0
	dir := []int{1, -1}
	for len(queue) > 0 {
		res++
		length := len(queue)
		for i := 0; i < length; i++ {
			str := queue[i]
			for j := 0; j < 4; j++ {
				for k := 0; k < len(dir); k++ {
					char := string((int(str[j]-'0')+10+dir[k])%10 + '0')
					newStr := str[:j] + char + str[j+1:]
					if _, ok := m[newStr]; ok {
						continue
					}
					queue = append(queue, newStr)
					m[newStr] = 1
					if newStr == target {
						return res
					}
				}
			}
		}
		queue = queue[length:]
	}
	return -1
}
