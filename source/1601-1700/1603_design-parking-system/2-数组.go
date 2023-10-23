package main

func main() {

}

// leetcode1603_设计停车系统
type ParkingSystem struct {
	arr [3]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		arr: [3]int{big, medium, small},
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.arr[carType-1] > 0 {
		this.arr[carType-1]--
		return true
	}
	return false
}
