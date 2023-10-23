package main

func main() {

}

type SubrectangleQueries struct {
	arr    [][]int
	record [][]int // 更改记录
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{arr: rectangle}
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	this.record = append(this.record, []int{row1, col1, row2, col2, newValue})
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
	for i := len(this.record) - 1; i >= 0; i-- {
		row1, col1 := this.record[i][0], this.record[i][1]
		row2, col2 := this.record[i][2], this.record[i][3]
		newValue := this.record[i][4]
		if row1 <= row && row <= row2 &&
			col1 <= col && col <= col2 {
			return newValue
		}
	}
	return this.arr[row][col]
}
