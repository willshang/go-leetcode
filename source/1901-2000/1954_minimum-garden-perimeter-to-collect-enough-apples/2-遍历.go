package main

func main() {

}

func minimumPerimeter(neededApples int64) int64 {
	var i int64
	for i = 1; ; i = i + 1 {
		total := 2 * i * (i + 1) * (2*i + 1)
		if neededApples <= total {
			return i * 8
		}
	}
	return 0
}
