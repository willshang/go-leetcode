package main

func main() {

}

// 剑指OfferII114.外星文字典
func alienOrder(words []string) string {
	n := len(words)
	degree := make(map[int]int) // 入度
	arr := [26][]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < len(words[i]); j++ {
			degree[int(words[i][j]-'a')] = 0
		}
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			length := min(len(words[i]), len(words[j]))
			for k := 0; k < length; k++ {
				a, b := int(words[i][k]-'a'), int(words[j][k]-'a')
				if a == b {
					if k == length-1 && len(words[i]) > len(words[j]) { // 不合法
						return ""
					}
					continue
				}
				arr[a] = append(arr[a], b) // 有序关系：a < b => a=>b
				degree[b]++
				break // 已经有有序关系，退出
			}
		}
	}
	// 拓扑排序
	res := make([]byte, 0)
	queue := make([]int, 0) // 入度=0 入队
	for i := 0; i < 26; i++ {
		if v, ok := degree[i]; ok && v == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		res = append(res, byte('a'+node))
		for i := 0; i < len(arr[node]); i++ {
			next := arr[node][i]
			degree[next]--
			if degree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}
	/*
		// 通过判断长度来判断拓扑排序完成
		if len(res) != len(degree) {
			return ""
		}
	*/
	// 通过判断度数来判断拓扑排序完成
	for i := 0; i < 26; i++ {
		if degree[i] > 0 {
			return ""
		}
	}
	return string(res)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
