package main

import (
	"fmt"
)

func main() {
	fmt.Println(hasAlternatingBits(5))
	fmt.Println(hasAlternatingBits(7))
	fmt.Println(hasAlternatingBits(4))
}

/*举几个例子就可以发现，对于0和1交错出现的数字网右移一位
比如1010101会变成101010，二者相加变成1111111，在加上1就会变10000000，
那么1111111&10000000=0；
如果n=11,二进制为1011，n>>1=101,101+1011+1=10001
这个数字和101取“与”是得不到0的*/
func hasAlternatingBits(n int) bool {
	return ((n + (n >> 1)) & (n + (n >> 1) + 1)) == 0
}

/*func hasAlternatingBits(n int) bool {
	std := n & 3

	fmt.Println(std)
	if std != 1 && std != 2{
		return false
	}

	for n > 0{
		if n&3 != std {
			return false
		}
		n = n >> 2
	}
	return true
}*/
