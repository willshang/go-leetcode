package main

func main() {

}

// leetcode1688_比赛中的配对次数
func numberOfMatches(n int) int {
	res := 0
	for n > 1 {
		res = res + n/2
		n = n/2 + n%2
	}
	return res
}
