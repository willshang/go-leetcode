# 排序算法
- 参考 
- https://www.cnblogs.com/onepixel/articles/7674659.html
## 总结

## 1.冒泡排序
```go
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

```