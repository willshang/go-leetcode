package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(watchedVideosByFriends([][]string{
		{"m"},
		{"xaqqqqqq"},
		{"lhvh"},
	}, [][]int{{1}, {0}}, 1, 1))
}

// leetcode1311_获取你好友已观看的视频
func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	m := make(map[string]int)
	visited := make(map[int]bool)
	visited[id] = true
	queue := make([]int, 0)
	queue = append(queue, id)
	for len(queue) > 0 {
		level--
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[i]
			for j := 0; j < len(friends[node]); j++ {
				ID := friends[node][j]
				if visited[ID] == false {
					visited[ID] = true
					queue = append(queue, ID)
					if level == 0 {
						for _, v := range watchedVideos[ID] {
							m[v]++
						}
					}
				}
			}
		}
		if level == 0 {
			break
		}
		queue = queue[length:]
	}
	res := make([]string, 0)
	for k := range m {
		res = append(res, k)
	}
	sort.Slice(res, func(i, j int) bool {
		if m[res[i]] == m[res[j]] {
			return res[i] < res[j]
		}
		return m[res[i]] < m[res[j]]
	})
	return res
}
