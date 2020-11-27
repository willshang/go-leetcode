package main

import "fmt"

func main() {
	arr := []int{66, 33, 55, 22, 11, 99, 88, 77}
	buildHeap(arr) // 构建堆
	fmt.Println(arr)
	// 插入使用上浮
	arr = append(arr, 1)
	up(arr)
	fmt.Println(arr)
	// 删除使用下沉
	value := pop(arr)
	fmt.Println(value, arr)
}

func pop(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	res := arr[0]
	arr[0] = arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	down(arr, 0, len(arr))
	return res
}

// 构建堆[最小堆为例]
func buildHeap(arr []int) {
	// 从最后一个非叶子节点开始
	for i := (len(arr) - 2) / 2; i >= 0; i-- {
		down(arr, i, len(arr))
	}
}

// 下沉[最小堆为例]
func down(arr []int, parentId int, length int) {
	temp := arr[parentId]
	childId := 2*parentId + 1 // 左节点
	for childId < length {
		// 有右节点且右节点小于左节点，childId+1
		if childId+1 < length && arr[childId+1] < arr[childId] { // 小根堆
			// if childId+1 < length && arr[childId+1] > arr[childId] { // 大根堆
			childId = childId + 1
		}
		if temp <= arr[childId] { // 小根堆
			// if temp >= arr[childId] { // 大根堆
			break
		}
		arr[parentId] = arr[childId]
		parentId = childId
		childId = 2*childId + 1
	}
	arr[parentId] = temp
}

// 上浮[最小堆为例]
func up(arr []int) {
	childId := len(arr) - 1                   // 最后一个节点下标（子节点）
	last := arr[childId]                      // 最后一个节点
	parentId := (childId - 1) / 2             // 父节点
	for childId > 0 && last < arr[parentId] { // 小根堆
		// for childId > 0 && last > arr[parentId] { // 大根堆
		arr[childId] = arr[parentId] //
		childId = parentId           //
		parentId = (childId - 1) / 2 //
	}
	arr[childId] = last
}
