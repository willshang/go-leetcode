package main

func main() {

}

// leetcode2076_处理含限制条件的好友请求
func friendRequests(n int, restrictions [][]int, requests [][]int) []bool {
	fa = Init(n)
	res := make([]bool, len(requests))
	for i := 0; i < len(requests); i++ {
		a, b := requests[i][0], requests[i][1]
		temp := make([]int, n)
		copy(temp, fa) // 备份当前的结果
		union(a, b)    // 尝试连接
		flag := true
		for j := 0; j < len(restrictions); j++ {
			c, d := restrictions[j][0], restrictions[j][1]
			if query(c, d) == true { // 查看是否有限制
				flag = false
				break
			}
		}
		if flag == true {
			res[i] = true
		} else {
			fa = make([]int, n) // 不满足条件，恢复回去
			copy(fa, temp)
		}
	}
	return res
}

var fa []int

// 初始化
func Init(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return arr
}

// 查询
func find(x int) int {
	if fa[x] != x {
		fa[x] = find(fa[x])
	}
	return fa[x]
}

// 合并
func union(i, j int) {
	fa[find(i)] = find(j)
}

func query(i, j int) bool {
	return find(i) == find(j)
}
