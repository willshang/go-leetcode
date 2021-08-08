package main

import "fmt"

func main() {
	//fmt.Println(isTransformable("84532", "34852"))
	fmt.Println(isTransformable("23", "32"))
}

// leetcode1585_检查字符串是否可以通过排序子字符串得到另一个字符串
func isTransformable(s string, t string) bool {
	n := len(s)
	arr := [10][]int{}
	for i := 0; i < n; i++ {
		arr[int(s[i]-'0')] = append(arr[int(s[i]-'0')], i)
	}
	for i := 0; i < n; i++ {
		index := int(t[i] - '0')
		if len(arr[index]) == 0 {
			return false
		}
		// 看当前位置s中等于t[i]的第x个元素，能不能移动到前面
		// 一次交换前后2个，目标值往前移（冒泡排序）
		for j := 0; j < index; j++ {
			// 前面不能存在比当前值小的，因为2个数排序后当前值会在后面
			if len(arr[j]) > 0 && arr[j][0] < arr[index][0] {
				return false
			}
		}
		arr[index] = arr[index][1:] // 当前数字匹配到了，往后移
	}
	return true
}
