package main

import (
	"fmt"
)

func main() {
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(10))
}

func fib(N int) int {
	if N == 0 {
		return 0
	}
	/*
		ans = [Fn+1 Fn
			   Fn Fn-1]
			= [ 1 0
		 		0 1]
	*/
	ans := matrix{
		a: 1,
		b: 0,
		c: 0,
		d: 1,
	}
	m := matrix{
		a: 1,
		b: 1,
		c: 1,
		d: 0,
	}
	for N > 0 {
		if N%2 == 1 {
			ans = multi(ans, m)
		}
		m = multi(m, m)
		N = N >> 1
	}
	return ans.b
}

/*
a b
c d
*/
type matrix struct {
	a, b, c, d int
}

// 矩阵乘法
func multi(x, y matrix) matrix {
	newA := x.a*y.a + x.b*y.c
	newB := x.a*y.b + x.b*y.d
	newC := x.c*y.a + x.d*y.c
	newD := x.c*y.b + x.d*y.d
	return matrix{
		a: newA,
		b: newB,
		c: newC,
		d: newD,
	}
}
