package main

func main() {

}

var mod = 1000000007

func countHousePlacements(n int) int {
	var a, b, c, d int
	a = 1 // 2侧没有房子
	b = 1 // 1侧有房子：上
	c = 1 // 1侧有房子：下
	d = 1 // 2侧都有房子
	for i := 2; i <= n; i++ {
		a, b, c, d = (a+b+c+d)%mod, (a+c)%mod, (a+b)%mod, a%mod
	}
	return (a + b + c + d) % mod
}
