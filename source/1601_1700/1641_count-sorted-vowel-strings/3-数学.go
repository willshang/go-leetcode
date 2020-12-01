package main

func main() {

}

// 将n个小球放到m个盒子里，盒子可以空：C(n+m-1, m-1) m=5 => C(n+4,4)
// 在n+4中选择4个整数(4*3*2)
func countVowelStrings(n int) int {
	return (n + 4) * (n + 3) * (n + 2) * (n + 1) / 24
}
