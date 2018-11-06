package main

import "fmt"

func main() {
	fmt.Println(getRow(3))
}
/*func getRow(rowIndex int) []int {
	res := make([]int,1,rowIndex+1)
	res[0] = 1
	if rowIndex == 0{
		return res
	}

	for i := 0; i < rowIndex; i++{
		res = append(res,1)
		fmt.Println(res)
		for j := len(res) -2 ; j > 0; j--{
			res[j] = res[j] + res[j-1]
		}

	}
	return res
}*/

func getRow(rowIndex int) []int {
	var result [][]int
	for i := 0; i < rowIndex+1; i++{
		var row []int
		for j := 0; j <= i; j++{
			tmp := 1
			if j == 0 || j == i{

			}else {
				tmp = result[i-1][j-1] + result[i-1][j]
			}
			row = append(row,tmp)
		}
		result = append(result,row)
	}
	return result[rowIndex]
}