package main

import (
	"fmt"
)

func main() {
	fmt.Println(distributeCandies(7, 4))
}

func distributeCandies(candies int, num_people int) []int {
	res := make([]int, num_people)
	times := 1
	for times*(times+1)/2 <= candies {
		times++
	}
	// 计算出当前糖果最多可以发给多少个人,剩下最后一个人多少颗糖
	times--
	last := candies - times*(times+1)/2
	for i := 0; i < num_people; i++ {
		n := times / num_people
		if times%num_people > i {
			n = n + 1
		}
		// 等差数列{an}的通项公式为：an=a1+(n-1)d。
		// 前n项和公式为：Sn=n*a1+n(n-1)d/2或Sn=n(a1+an)/2
		// Sn=n(a1+a1+(n-1)d)/2=n(2a1+(n-1)d)/2
		// (i+1)为首项，num_people为公差，n为数列长度，的等差数列的和
		res[i] = n * (2*(i+1) + (n-1)*num_people) / 2
		if times%num_people == i {
			res[i] = res[i] + last
		}
	}
	return res
}
