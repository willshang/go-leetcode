package main

func main() {

}

type Node struct {
	Name  string
	Child []*Node
}

type ThroneInheritance struct {
	isDead map[string]bool
	m      map[string]*Node
	king   *Node
}

func Constructor(kingName string) ThroneInheritance {
	node := &Node{
		Name:  kingName,
		Child: make([]*Node, 0),
	}
	res := ThroneInheritance{
		king:   node,
		m:      map[string]*Node{},
		isDead: map[string]bool{},
	}
	res.m[kingName] = node
	return res
}

func (this *ThroneInheritance) Birth(parentName string, childName string) {
	node := this.m[parentName]
	child := &Node{
		Name:  childName,
		Child: make([]*Node, 0),
	}
	this.m[childName] = child
	node.Child = append(node.Child, child)
}

func (this *ThroneInheritance) Death(name string) {
	this.isDead[name] = true
}

func (this *ThroneInheritance) GetInheritanceOrder() []string {
	res := make([]string, 0)
	root := this.king
	stack := make([]*Node, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if this.isDead[node.Name] == false {
			res = append(res, node.Name)
		}
		if len(node.Child) > 0 {
			for i := len(node.Child) - 1; i >= 0; i-- {
				stack = append(stack, node.Child[i])
			}
		}
	}
	return res
}
