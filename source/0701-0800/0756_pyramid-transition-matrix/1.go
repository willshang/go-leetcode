package main

func main() {

}

func pyramidTransition(bottom string, allowed []string) bool {
	total := 1 << 7
	status := make([][]int, total)
	for i := 0; i < total; i++ {
		status[i] = make([]int, total)
	}
	for i := 0; i < len(allowed); i++ {
		a := 1 << int(allowed[i][0]-'A')
		b := 1 << int(allowed[i][1]-'A')
		c := 1 << int(allowed[i][2]-'A')
		for j := 0; j < total; j++ {
			if (a & j) > 0 {
				for k := 0; k < total; k++ {
					if (b & k) > 0 {
						status[j][k] = status[j][k] | c
					}
				}
			}
		}
	}

}
