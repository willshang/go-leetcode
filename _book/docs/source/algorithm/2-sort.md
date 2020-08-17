# 排序算法
- 参考 
- https://www.cnblogs.com/onepixel/articles/7674659.html
## 总结

| No.  | 排序方法 | 时间复杂度(平均) | 时间复杂度(最坏) | 时间复杂度(最好) | 空间复杂度 | 稳定性 |
| ---- | -------- | ---------------- | ---------------- | ---------------- | ---------- | ------ |
| 01   | 冒泡排序 | O(n^2)           | O(n^2)           | O(n)             | O(1)       | 稳定   |
| 02   | 选择排序 | O(n^2)           | O(n^2)           | O(n^2)           | O(1)       | 不稳定 |
| 03   | 插入排序 | O(n^2)           | O(n^2)           | O(n)             | O(1)       | 稳定   |
| 04   | 希尔排序 | O(n^1.3)         | O(n^2)           | O(n)             | O(1)       | 不稳定 |
| 05   | 归并排序 | O(nlog(n))       | P                | O(nlog(n))       | O(n)       | 稳定   |
| 06   | 快速排序 | O(nlog(n))       | O(n^2)           | O(nlog(n))       | O(log(n))  | 不稳定 |
| 07   | 堆排序   | O(nlog(n))       | O(nlog(n))       | O(nlog(n))       | O(1)       | 不稳定 |
| 08   | 计数排序 | O(n+k)           | O(n+k)           | O(n+k)           | O(n+k)     | 稳定   |
| 09   | 桶排序   | O(n+k)           | O(n^2)           | O(n)             | O(n+k)     | 稳定   |
| 10   | 基数排序 | O(n*k)           | O(n*k)           | O(n*k)           | O(n+k)     | 稳定   |

## 1.冒泡排序
```go
// 依次比较和交换，把最大值或者最小值交换到后面已经排好序数组之前
func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
```
## 2.选择排序
```go
// 每次选择未排好序数组的最小值，添加到已经排好序后面
// 序列5 8 5 2 9，第一遍选择第1个元素5会和2交换，破坏稳定性
func selectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
```

## 3.插入排序

```go
// 依次选择一个数据，插入到前面已经排好序中，需要数据后移
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		pos := i - 1 // 已经有序的最后一个下标
		cur := arr[i]
		for pos >= 0 && arr[pos] > cur {
			arr[pos+1] = arr[pos] // 后移
			pos--
		}
		arr[pos+1] = cur
	}
}
```

## 4.希尔排序

```go
// 选定间隔
func shellSort(arr []int) {
	n := len(arr)
	for gap := n / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < n; i++ {
			j := i
			cur := arr[i]
			for j-gap >= 0 && cur < arr[j-gap] {
				arr[j] = arr[j-gap]
				j = j - gap
			}
			arr[j] = cur
		}
	}
}
```

## 5.归并排序

```go
func mergeSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}
	return merge(mergeSort(arr[:len(arr)/2]), mergeSort(arr[len(arr)/2:]))
}

func merge(left, right []int) []int {
	res := make([]int, 0)
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}
	if len(left) > 0 {
		res = append(res, left...)
	}
	if len(right) > 0 {
		res = append(res, right...)
	}
	return res
}
```

## 6.快排排序

```go
func quickSort(arr []int) {
	quick(arr, 0, len(arr)-1)
}

func quick(arr []int, left, right int) {
	if left >= right {
		return
	}
	index := partition(arr, left, right)
	quick(arr, left, index-1)
	quick(arr, index+1, right)
}

func partition(arr []int, left, right int) int {
	baseValue := arr[left] // 基准值
	for left < right {
		for baseValue <= arr[right] && left < right {
			right-- // 依次查找大于基准值的位置
		}
		arr[left] = arr[right]
		for arr[left] <= baseValue && left < right {
			left++ // 依次查找小于基准值的位置
		}
		arr[right] = arr[left]
	}
	arr[right] = baseValue
	return right
}
```

## 7.堆排序

- 堆排序

```go
func heapSort(arr []int) {
	buildHeap(arr, len(arr))
	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapModify(arr, 0, i-1)
	}
}

func buildHeap(arr []int, length int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		heapModify(arr, i, length-1)
	}
}

func heapModify(arr []int, start, end int) {
	temp := arr[start]
	for left := 2*start + 1; left <= end; left = 2*left + 1 {
		if left < end && arr[left] < arr[left+1] {
			left++
		}
		if arr[start] < arr[left] {
			arr[start] = arr[left]
			start = left
		}
		arr[start] = temp
	}
}
```

- 内置heap

```go
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

// 小根堆<,大根堆变换方向>
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}

func heapSort(arr []int) []int {
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	for i := 0; i < len(arr); i++ {
		heap.Push(&intHeap, arr[i])
	}
	res := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		value := heap.Pop(&intHeap).(int)
		res = append(res, value)
	}
	return res
}
```

## 8.计数排序

```go
func countingSort(arr []int) []int {
	maxValue := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxValue {
			maxValue = arr[i]
		}
	}
	bucket := make([]int, maxValue+1)
	for i := 0; i < len(arr); i++ {
		bucket[arr[i]]++
	}
	index := 0
	for i := 0; i < len(bucket); i++ {
		for bucket[i] > 0 {
			arr[index] = i
			index++
			bucket[i]--
		}
	}
	return arr
}
```

## 9.桶排序

```go

```

## 10.基数排序

```go

```

