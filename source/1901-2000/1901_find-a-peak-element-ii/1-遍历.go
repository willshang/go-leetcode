package main

func main() {

}

func findPeakGrid(mat [][]int) []int {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if (1 <= i && mat[i][j] < mat[i-1][j]) ||
				(i < len(mat)-1 && mat[i][j] < mat[i+1][j]) ||
				(1 <= j && mat[i][j] < mat[i][j-1]) ||
				(j < len(mat[i])-1 && mat[i][j] < mat[i][j+1]) {
				continue
			}
			return []int{i, j}
		}
	}
	return nil
}
