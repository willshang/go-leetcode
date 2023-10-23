package main

func main() {

}

// 程序员面试金典17.22_单词转换
func findLadders(beginWord string, endWord string, wordList []string) []string {
	m, preMap := make(map[string]int), make(map[string][]string)
	for i := 0; i < len(wordList); i++ {
		m[wordList[i]] = 1
	}
	if m[endWord] == 0 {
		return nil
	}
	for i := 0; i < len(wordList); i++ {
		for j := 0; j < len(wordList[i]); j++ {
			newStr := wordList[i][:j] + "*" + wordList[i][j+1:]
			preMap[newStr] = append(preMap[newStr], wordList[i])
		}
	}
	visited := make(map[string]bool)
	queue, path := make([]string, 0), make([][]string, 0)
	queue, path = append(queue, beginWord), append(path, []string{beginWord})
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			for j := 0; j < len(beginWord); j++ {
				newStr := queue[i][:j] + "*" + queue[i][j+1:]
				temp := make([]string, len(path[i]))
				copy(temp, path[i])
				for _, word := range preMap[newStr] {
					if word == endWord {
						return append(temp, endWord)
					} else if visited[word] == false {
						visited[word] = true
						queue, path = append(queue, word), append(path, append(temp, word))
					}
				}
			}
		}
		queue, path = queue[length:], path[length:]
	}
	return nil
}
