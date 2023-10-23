package main

func main() {

}

func getHappyString(n int, k int) string {
	per := 1 << (n - 1)
	if k > 3*per {
		return ""
	}
	res := make([]byte, n)
	next := [][]int{{1, 2}, {0, 2}, {0, 1}} // 3个分支
	k = k - 1
	var status int
	for i := 0; i < n; i++ {
		if i == 0 { // 确定第1个分支
			status = k / per
		} else {
			per = per / 2
			status = next[status][k/per] // 确定后面的分支
		}
		k = k % per
		res[i] = byte(status + 'a')
	}
	return string(res)
}
