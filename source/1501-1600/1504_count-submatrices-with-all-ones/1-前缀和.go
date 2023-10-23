package main

func main() {

}

func numSubmat(mat [][]int) int {
	res := 0
	n, m := len(mat), len(mat[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j > 0 && mat[i][j] == 1 {
				mat[i][j] = mat[i][j-1] + 1 // 到左边连续1的个数
			}
			target := mat[i][j]       // 底：当前行以mat[i][j]=1结尾的连续1的个数
			for k := i; k >= 0; k-- { // 遍历高
				if target > mat[k][j] { // 上一层连续1的个数小
					target = mat[k][j] // 缩小底
				}
				// 示例1：右下角一行3个1
				// 1(1) 1(2) 1(3) => 有3种组成矩形的情况
				// 第1种情况.1(3)
				// 第2种情况.1(2)+1(3)
				// 第3种情况.1(1)+1(2)+1(3)
				res = res + target // 加上该高度的个数
			}
		}
	}
	return res
}
