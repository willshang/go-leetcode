package main

import "fmt"

func main() {
	fmt.Println(fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16))
}

// leetcode68_文本左右对齐
func fullJustify(words []string, maxWidth int) []string {
	res := make([]string, 0)
	count := 0
	start := 0
	for i := 0; i < len(words); i++ {
		count = count + len(words[i])
		if count > maxWidth {
			temp := justify(words, start, i-1, maxWidth)
			res = append(res, temp)
			start = i
			if i == len(words)-1 {
				count = 0
				i--
			} else {
				count = len(words[i]) + 1
			}
		} else if i == len(words)-1 {
			temp := justify(words, start, i, maxWidth)
			res = append(res, temp)
		} else {
			count++
		}
	}
	return res
}

func justify(words []string, start, end int, maxWidth int) string {
	arr := make([]byte, maxWidth)
	for i := 0; i < len(arr); i++ {
		arr[i] = byte(' ')
	}
	index := 0
	// 文本的最后一行应为左对齐，且单词之间不插入额外的空格。
	if start == end || end == len(words)-1 {
		for i := start; i <= end; i++ {
			copy(arr[index:], words[i])
			index = index + len(words[i]) + 1
		}
	} else {
		// 要求尽可能均匀分配单词间的空格数量。
		// 如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。
		count := end - start + 1
		left := maxWidth - count + 1
		for i := start; i <= end; i++ {
			left = left - len(words[i])
		}
		space := left / (count - 1) // 均分
		mod := left % (count - 1)   // 多的放左边
		for i := start; i <= end; i++ {
			copy(arr[index:], words[i])
			index = index + len(words[i]) + 1 + space
			if mod > 0 {
				index++
				mod--
			}
		}
	}
	return string(arr)
}
