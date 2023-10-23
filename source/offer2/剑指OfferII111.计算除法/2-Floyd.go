package main

func main() {

}

// 剑指OfferII111.计算除法
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	m := make(map[string]int) // 计算对应的id
	for i := 0; i < len(equations); i++ {
		a, b := equations[i][0], equations[i][1]
		if _, ok := m[a]; ok == false {
			m[a] = len(m)
		}
		if _, ok := m[b]; ok == false {
			m[b] = len(m)
		}
	}
	arr := make([][]float64, len(m)) // 邻接矩阵
	for i := 0; i < len(m); i++ {
		arr[i] = make([]float64, len(m))
	}
	for i := 0; i < len(equations); i++ {
		a, b := m[equations[i][0]], m[equations[i][1]]
		arr[a][b] = values[i]
		arr[b][a] = 1 / values[i]
	}
	for k := 0; k < len(arr); k++ { // Floyd
		for i := 0; i < len(arr); i++ {
			for j := 0; j < len(arr); j++ {
				if arr[i][k] > 0 && arr[k][j] > 0 {
					arr[i][j] = arr[i][k] * arr[k][j]
				}
			}
		}
	}
	res := make([]float64, len(queries))
	for i := 0; i < len(queries); i++ {
		a, okA := m[queries[i][0]]
		b, okB := m[queries[i][1]]
		if okA == false || okB == false || arr[a][b] == 0 {
			res[i] = -1
		} else {
			res[i] = arr[a][b]
		}
	}
	return res
}
