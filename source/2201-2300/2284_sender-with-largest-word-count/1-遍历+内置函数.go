package main

import "strings"

func main() {

}

// leetcode2284_最多单词数的发件人
func largestWordCount(messages []string, senders []string) string {
	res := ""
	maxValue := 0
	m := make(map[string]int)
	for i := 0; i < len(senders); i++ {
		count := len(strings.Fields(messages[i]))
		m[senders[i]] = m[senders[i]] + count
		if m[senders[i]] > maxValue {
			maxValue = m[senders[i]]
			res = senders[i]
		} else if m[senders[i]] == maxValue && senders[i] > res {
			res = senders[i]
		}
	}
	return res
}
