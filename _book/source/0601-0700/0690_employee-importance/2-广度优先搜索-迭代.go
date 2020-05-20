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

func getImportance(employees []*Employee, id int) int {
	if len(employees) == 0 {
		return 0
	}
	m := make(map[int]*Employee)
	for i := 0; i < len(employees); i++ {
		m[employees[i].Id] = employees[i]
	}
	root := m[id]
	if root == nil {
		return 0
	}
	res := 0
	list := make([]*Employee, 0)
	list = append(list, root)
	for len(list) > 0 {
		node := list[0]
		list = list[1:]
		res = res + node.Importance
		for i := range node.Subordinates {
			if value, ok := m[node.Subordinates[i]]; ok {
				list = append(list, value)
			}
		}
	}
	return res
}
