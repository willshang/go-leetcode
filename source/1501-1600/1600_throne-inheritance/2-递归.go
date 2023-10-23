package main

func main() {

}

// leetcode1600_皇位继承顺序
type ThroneInheritance struct {
	isDead   map[string]bool
	children map[string][]string
	king     string
}

func Constructor(kingName string) ThroneInheritance {
	return ThroneInheritance{
		isDead:   make(map[string]bool),
		children: make(map[string][]string),
		king:     kingName,
	}
}

func (this *ThroneInheritance) Birth(parentName string, childName string) {
	this.children[parentName] = append(this.children[parentName], childName)
}

func (this *ThroneInheritance) Death(name string) {
	this.isDead[name] = true
}

func (this *ThroneInheritance) GetInheritanceOrder() []string {
	return dfs(this, this.king)
}

func dfs(this *ThroneInheritance, name string) []string {
	res := make([]string, 0)
	if this.isDead[name] == false {
		res = append(res, name)
	}
	for _, child := range this.children[name] {
		res = append(res, dfs(this, child)...)
	}
	return res
}
