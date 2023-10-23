package main

func main() {

}

func findLadders(beginWord string, endWord string, wordList []string) []string {
	m := make(map[string]int)
	for i := 0; i < len(wordList); i++ {
		m[wordList[i]] = 1
	}
	if m[endWord] == 0 {
		return nil
	}
	preMap := make(map[string][]string)
	for i := 0; i < len(wordList); i++ {
		for j := 0; j < len(wordList[i]); j++ {
			newStr := wordList[i][:j] + "*" + wordList[i][j+1:]
			if _, ok := preMap[newStr]; !ok {
				preMap[newStr] = make([]string, 0)
			}
			preMap[newStr] = append(preMap[newStr], wordList[i])
		}
	}
	visited := make(map[string]bool)
	count := 0
	queue := make([]string, 0)
	queue = append(queue, beginWord)
	path := make([][]string, 0)
	path = append(path, []string{beginWord})
	for len(queue) > 0 {
		count++
		node := queue[0]
		queue = queue[1:]
		arr := path[0]
		path = path[1:]
		for j := 0; j < len(beginWord); j++ {
			newStr := node[:j] + "*" + node[j+1:]
			temp := make([]string, len(arr))
			copy(temp, arr)
			for _, word := range preMap[newStr] {
				if word == endWord {
					return append(temp, endWord)
				}
				if visited[word] == false {
					visited[word] = true
					queue = append(queue, word)
					path = append(path, append(temp, word))
				}
			}
		}
	}
	return nil
}
