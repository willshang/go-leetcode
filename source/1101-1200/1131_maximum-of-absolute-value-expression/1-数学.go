package main

func main() {

}

func maxAbsValExpr(arr1 []int, arr2 []int) int {
	/*
		|arr1[i] - arr1[j]| + |arr2[i] - arr2[j]| + |i - j|
		=  (arr1[i] + arr2[i] + i) - (arr1[j] + arr2[j] + j)
		=  (arr1[i] + arr2[i] - i) - (arr1[j] + arr2[j] - j)
		=  (arr1[i] - arr2[i] + i) - (arr1[j] - arr2[j] + j)
		=  (arr1[i] - arr2[i] - i) - (arr1[j] - arr2[j] - j)
		= -(arr1[i] + arr2[i] + i) + (arr1[j] + arr2[j] + j)
		= -(arr1[i] + arr2[i] - i) + (arr1[j] + arr2[j] - j)
		= -(arr1[i] - arr2[i] + i) + (arr1[j] - arr2[j] + j)
		= -(arr1[i] - arr2[i] - i) + (arr1[j] - arr2[j] - j)
		其中：
		A = arr1[i] + arr2[i] + i
		B = arr1[i] + arr2[i] - i
		C = arr1[i] - arr2[i] + i
		D = arr1[i] - arr2[i] - i
		结果：
		max( |arr1[i] - arr1[j]| + |arr2[i] - arr2[j]| + |i - j|)
		= max(max(A) - min(A),max(B) - min(B),max(C) - min(C), max(D) - min(D))
	*/
	arr := make([][]int, 4)
	for i := 0; i < len(arr1); i++ {
		a, b := arr1[i], arr2[i]
		arr[0] = append(arr[0], a+b+i)
		arr[1] = append(arr[1], a+b-i)
		arr[2] = append(arr[2], a-b+i)
		arr[3] = append(arr[3], a-b-i)
	}
	a, b, c, d := getValue(arr[0]), getValue(arr[1]), getValue(arr[2]), getValue(arr[3])
	return max(a, max(b, max(c, d)))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getValue(arr []int) int {
	minValue, maxValue := arr[0], arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxValue {
			maxValue = arr[i]
		}
		if arr[i] < minValue {
			minValue = arr[i]
		}
	}
	return maxValue - minValue
}
