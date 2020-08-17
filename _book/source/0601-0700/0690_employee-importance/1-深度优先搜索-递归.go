package main

import "fmt"

func main() {
	arr := make([]*Employee, 0)
	first := &Employee{
		Id:           1,
		Importance:   5,
		Subordinates: []int{2, 3},
	}
	second := &Employee{
		Id:           2,
		Importance:   3,
		Subordinates: []int{},
	}
	third := &Employee{
		Id:           3,
		Importance:   3,
		Subordinates: []int{},
	}
	arr = append(arr, []*Employee{first, second, third}...)
	fmt.Println(getImportance(arr, 1))
}

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

// leetcode690_员工的重要性
func getImportance(employees []*Employee, id int) int {
	if len(employees) == 0 {
		return 0
	}
	var root *Employee
	for i := 0; i < len(employees); i++ {
		if employees[i].Id == id {
			root = employees[i]
		}
	}
	if root == nil {
		return 0
	}
	res := root.Importance
	for i := range root.Subordinates {
		res = res + getImportance(employees, root.Subordinates[i])
	}
	return res
}
