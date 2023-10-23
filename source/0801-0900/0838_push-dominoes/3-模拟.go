package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(pushDominoes(".L.R...LR..L.."))
	fmt.Println(pushDominoes(".......LR....."))
}

func pushDominoes(dominoes string) string {
	s := ""
	for s != dominoes { // 模拟每一秒：当没有变化的时候退出
		s = dominoes
		dominoes = strings.ReplaceAll(dominoes, "R.L", "X") // 把R.L这种不变的临时替换为1个不影响的X
		dominoes = strings.ReplaceAll(dominoes, ".L", "LL") // 向左倒
		dominoes = strings.ReplaceAll(dominoes, "R.", "RR") // 向右倒
		dominoes = strings.ReplaceAll(dominoes, "X", "R.L") // 替换回来
	}
	return dominoes
}
