package main

func main() {

}

// leetcode464_我能赢吗
var m map[int]bool

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	a := maxChoosableInteger
	if a*(a+1)/2 < desiredTotal {
		return false
	}
	m = make(map[int]bool)
	return dfs(a, desiredTotal, 0)
}

func dfs(a, b int, status int) bool {
	if v, ok := m[status]; ok {
		return v
	}
	for i := 1; i <= a; i++ {
		cur := 1 << (i - 1)
		next := status | cur                                           // 按位或运算：status第(i-1)变为1
		if cur&status == 0 && (b <= i || dfs(a, b-i, next) == false) { // 当前位(i-1)为1
			m[status] = true
			return true
		}
	}
	m[status] = false
	return false
}
