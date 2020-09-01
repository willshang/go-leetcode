package main

import "container/list"

func main() {

}

// 程序员面试金典03.06_动物收容所
type AnimalShelf struct {
	arr [2]*list.List
}

func Constructor() AnimalShelf {
	return AnimalShelf{
		arr: [2]*list.List{list.New(), list.New()},
	}
}

func (this *AnimalShelf) Enqueue(animal []int) {
	this.arr[animal[1]].PushBack(animal[0])
}

func (this *AnimalShelf) DequeueAny() []int {
	if this.arr[0].Len() == 0 && this.arr[1].Len() == 0 {
		return []int{-1, -1}
	}
	if this.arr[1].Len() > 0 &&
		(this.arr[0].Len() == 0 || this.arr[1].Front().Value.(int) < this.arr[0].Front().Value.(int)) {
		return []int{this.arr[1].Remove(this.arr[1].Front()).(int), 1}
	}
	return []int{this.arr[0].Remove(this.arr[0].Front()).(int), 0}
}

func (this *AnimalShelf) DequeueDog() []int {
	if this.arr[1].Len() > 0 {
		return []int{this.arr[1].Remove(this.arr[1].Front()).(int), 1}
	}
	return []int{-1, -1}
}

func (this *AnimalShelf) DequeueCat() []int {
	if this.arr[0].Len() > 0 {
		return []int{this.arr[0].Remove(this.arr[0].Front()).(int), 0}
	}
	return []int{-1, -1}
}
