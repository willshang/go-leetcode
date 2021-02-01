package main

func main() {

}

// leetcode1024_视频拼接
func videoStitching(clips [][]int, T int) int {
	arr := make([]int, T)
	for i := 0; i < len(clips); i++ {
		a, b := clips[i][0], clips[i][1]
		if a < T && arr[a] < b {
			arr[a] = b // 更新当前位置能到达最远的位置
		}
	}
	last := 0
	prev := 0
	res := 0
	// 变成leetcode45.跳跃游戏II的变形
	for i := 0; i < len(arr); i++ {
		if arr[i] > last {
			last = arr[i]
		}
		if i == last { // 无法达到目标
			return -1
		}
		if i == prev {
			res++
			prev = last
		}
	}
	return res
}
