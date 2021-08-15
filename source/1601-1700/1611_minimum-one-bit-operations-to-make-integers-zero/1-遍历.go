package main

func main() {

}

func minimumOneBitOperations(n int) int {
	// 依次将高位的1翻转为0
	// 操作2：00..110000..00 => 00..010000..00
	// 操作2+操作1：00..010000..00 => 00..011000..00
	// 操作2：00..011000..00 => 00..001000..00
	res := 0
	if n == 0 {
		return 0
	}
	for i := 0; (1 << i) <= n; i++ {
		if (1<<i)&n > 0 { // 第i位>0
			total := 1<<(1+i) - 1 // 把(1+i)位数变为100..00的次数
			res = total - res     // 交替加减
		}
	}
	return res
}
