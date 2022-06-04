package main

import "strings"

func main() {

}

func largestWordCount(messages []string, senders []string) string {
	res := ""
	maxValue := 0
	m := make(map[string]int)
	for i := 0; i < len(senders); i++ {
		count := strings.Count(messages[i], " ") + 1
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
