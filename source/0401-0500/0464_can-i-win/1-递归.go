package main

func main() {

}

var m []bool

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	a := maxChoosableInteger
	if a*(a+1)/2 < desiredTotal {
		return false
	}
	m = make([]bool, 1<<a)
	return dfs(a, desiredTotal, 0)
}

func dfs(a, b int, status int) bool {
	if m[status] == true {
		return true
	}
	for i := 1; i <= a; i++ {
		cur := 1 << (i - 1)
		if cur&status > 0 { // 当前位(i-1)为1
			continue
		}
		if b <= i { // 当前选的数可以赢
			m[status] = true
			return true
		}
		next := status | cur            // 按位或运算：status第(i-1)变为1
		if dfs(a, b-i, next) == false { // 如果下个人要输的话，当前人要赢
			m[status] = true
			return true
		}
	}
	m[status] = false
	return false
}
