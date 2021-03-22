package main

func main() {

}

type ParkingSystem struct {
	arr   [3]int
	total [3]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		arr:   [3]int{big, medium, small},
		total: [3]int{0, 0, 0},
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	index := carType - 1
	if this.total[index] < this.arr[index] {
		this.total[index]++
		return true
	}
	return false
}
