package main

func main() {

}

func numberOfMatches(n int) int {
	// 共有n个队伍,一个冠军,需要淘汰n-1个 队伍。
	// 每一场比赛淘汰一个队伍,因此进行了n-1场比赛。
	return n - 1
}
