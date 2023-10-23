package main

func main() {

}

// Floyd
func Floyd(arr [][]int) [][]int {
	n := len(arr)
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if arr[i][k]+arr[k][j] < arr[i][j] {
					arr[i][j] = arr[i][k] + arr[k][j]
				}
			}
		}
	}
	return arr
}
