package main

func main() {

}

func numSpecial(mat [][]int) int {
	res := 0
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 && judge(mat, i, j) == true {
				res++
			}
		}
	}
	return res
}

func judge(mat [][]int, i, j int) bool {
	for a := 0; a < len(mat[i]); a++ {
		if mat[i][a] == 1 && a != j {
			return false
		}
	}
	for a := 0; a < len(mat); a++ {
		if mat[a][j] == 1 && a != i {
			return false
		}
	}
	return true
}
