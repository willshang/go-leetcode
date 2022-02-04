package main

func main() {

}

// leetcode_lcp23_魔术排列
func isMagic(target []int) bool {
	n := len(target)
	arr := make([]int, n) // 构建初始数组
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}
	arr = Shuffle(arr)
	k := 0 // 找到k：可以证明只能正好匹配到k个；其中1~N不重复
	for ; k < n && arr[k] == target[k]; k++ {
	}
	if k == 0 { // 不满足
		return false
	}
	for { // 按规则模拟
		if k >= len(arr) {
			return judge(arr, target, len(arr))
		}
		if judge(arr, target, k) == false {
			return false
		}
		arr = arr[k:]
		arr = Shuffle(arr)
		target = target[k:]
	}
	return false
}

func Shuffle(arr []int) []int {
	temp := make([]int, 0)
	for i := 1; i < len(arr); i = i + 2 {
		temp = append(temp, arr[i])
	}
	for i := 0; i < len(arr); i = i + 2 {
		temp = append(temp, arr[i])
	}
	return temp
}

func judge(a, b []int, k int) bool {
	for i := 0; i < k; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
