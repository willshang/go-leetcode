package main

import "fmt"

func main() {
	fmt.Println(minMutation("AAAAAAAA", "CCCCCCCC", []string{"AAAAAAAA", "AAAAAAAC", "AAAAAACC", "AAAAACCC",
		"AAAACCCC", "AACACCCC", "ACCACCCC", "ACCCCCCC", "CCCCCCCA", "CCCCCCCC"}))
}

func minMutation(start string, end string, bank []string) int {
	arr := []byte{'A', 'T', 'C', 'G'}
	m := make(map[string]bool)
	for i := 0; i < len(bank); i++ {
		m[bank[i]] = true
	}
	if _, ok := m[end]; ok == false {
		return -1
	}
	res := 0
	queueA := make([]string, 0)
	queueA = append(queueA, start)
	queueB := make([]string, 0)
	queueB = append(queueB, end)
	for len(queueA) > 0 {
		res++
		if len(queueA) > len(queueB) {
			queueA, queueB = queueB, queueA
		}
		length := len(queueA)
		for i := 0; i < length; i++ {
			str := queueA[i]
			for j := 0; j < len(str); j++ {
				for k := 0; k < len(arr); k++ {
					if arr[k] != str[j] {
						newStr := str[:j] + string(arr[k]) + str[j+1:]
						if _, ok := m[newStr]; ok {
							queueA = append(queueA, newStr)
							delete(m, newStr)
						}
						for l := 0; l < len(queueB); l++ {
							if newStr == queueB[l] {
								return res
							}
						}
					}
				}
			}
		}
		queueA = queueA[length:]
	}
	return -1
}
