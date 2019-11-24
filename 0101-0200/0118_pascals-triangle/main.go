package main

import "fmt"

func main() {
	res := generate(5)
	for _, v := range res {
		fmt.Println(v)
	}
}

func generate(numRows int) [][]int {
	var result [][]int
	for i := 0; i < numRows; i++ {
		var row []int
		for j := 0; j <= i; j++ {
			tmp := 1
			if j == 0 || j == i {

			} else {
				tmp = result[i-1][j-1] + result[i-1][j]
			}
			row = append(row, tmp)
		}
		result = append(result, row)
	}
	return result
}

/*func generate(numRows int) [][]int {
	res := [][]int{}
	if numRows == 0{
		return res
	}

	res = append(res,[]int{1})
	if numRows == 1{
		return res
	}

	for i := 1; i < numRows; i++{
		res = append(res,genNext(res[i-1]))
	}
	return res
}

func genNext(p []int) []int  {

	res := make([]int,1 ,len(p)+1)
	res = append(res,p...)

	for i := 0; i < len(res) - 1; i++{
		res[i] = res[i] + res[i+1]
	}
	return res
}
*/
