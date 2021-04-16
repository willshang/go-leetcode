package main

func main() {

}

// leetcode1359_有效的快递序列数目
var mod = 1000000007

func countOrders(n int) int {
	if n == 1 {
		return 1
	}
	res := 1
	// 前面有2(i-1)个数
	// 把Pi和Di当成1个插入，有2(i-1)+1=2i-1种插法
	// 把Pi和Di当前2个依次插入，有(2i-1)*(2i-2)/2=(2i-1)*(i-1)
	// 加起来 => (2*i-1)*i
	for i := 2; i <= n; i++ {
		res = res * (i*2 - 1) * i % mod
	}
	return res
}
