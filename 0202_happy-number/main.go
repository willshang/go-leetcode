package main

import "fmt"

func main() {
	fmt.Println(isHappy(19))
}

func isHappy(n int) bool {
	now, next := n, nextValue(n)
	m := make(map[int]int)
	m[now] = 1
	for {
		if next == 1{
			break
		}
		if _, ok := m[next]; ok{
			break
		}else {
			m[next] = 1
		}
		next = nextValue(next)
	}

	if next == 1{
		return true
	}

	return false
}

func nextValue(n int) int  {
	ret := 0
	for n != 0{
		ret = ret + (n%10) * (n%10)
		n = n / 10
	}
	return ret
}
/*func isHappy(n int) bool {
	now, next := n, nextValue(n)
	for now != next{
		now = nextValue(now)
		next = nextValue(nextValue(next))
	}
	if now == 1{
		return true
	}
	return false
}

func nextValue(n int) int  {
	ret := 0
	for n != 0{
		ret = ret + (n%10) * (n%10)
		n = n / 10
	}
	return ret
}*/