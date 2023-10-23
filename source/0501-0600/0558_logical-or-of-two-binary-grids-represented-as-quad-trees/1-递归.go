package main

func main() {

}

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

// leetcode558_四叉树交集
func intersect(quadTree1 *Node, quadTree2 *Node) *Node {
	res := &Node{}
	if quadTree1.IsLeaf == true { // 叶子节点
		if quadTree1.Val == true {
			return quadTree1
		}
		return quadTree2
	}
	if quadTree2.IsLeaf == true { // 叶子节点
		if quadTree2.Val == true {
			return quadTree2
		}
		return quadTree1
	}
	tL := intersect(quadTree1.TopLeft, quadTree2.TopLeft)
	tR := intersect(quadTree1.TopRight, quadTree2.TopRight)
	bL := intersect(quadTree1.BottomLeft, quadTree2.BottomLeft)
	bR := intersect(quadTree1.BottomRight, quadTree2.BottomRight)
	// 叶子节点判断
	if tL.IsLeaf == true && tR.IsLeaf == true && bL.IsLeaf == true && bR.IsLeaf == true &&
		tL.Val == tR.Val && tR.Val == bL.Val && bL.Val == bR.Val {
		res.IsLeaf = true
		res.Val = tL.Val // 4个值都相同
		return res
	}
	res.TopLeft = tL
	res.TopRight = tR
	res.BottomLeft = bL
	res.BottomRight = bR
	res.Val = false
	res.IsLeaf = false
	return res
}
