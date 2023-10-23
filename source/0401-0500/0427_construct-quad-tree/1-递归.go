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

func construct(grid [][]int) *Node {
	n := len(grid)
	return dfs(grid, 0, 0, n, n)
}

func dfs(grid [][]int, x1, y1, x2, y2 int) *Node {
	if x1+1 == x2 {
		value := false
		if grid[x1][y1] == 1 {
			value = true
		}
		return &Node{
			Val:    value,
			IsLeaf: true,
		}
	}
	midX := (x1 + x2) / 2
	midY := (y1 + y2) / 2
	tL := dfs(grid, x1, y1, midX, midY)
	tR := dfs(grid, x1, midY, midX, y2)
	bL := dfs(grid, midX, y1, x2, midY)
	bR := dfs(grid, midX, midY, x2, y2)
	if tL.IsLeaf == true && tR.IsLeaf == true && bL.IsLeaf == true && bR.IsLeaf == true &&
		((tL.Val == true && tR.Val == true && bL.Val == true && bR.Val == true) ||
			(tL.Val == false && tR.Val == false && bL.Val == false && bR.Val == false)) {
		return &Node{
			Val:    tL.Val,
			IsLeaf: true,
		}
	}
	return &Node{
		Val:         false,
		IsLeaf:      false,
		TopLeft:     tL,
		TopRight:    tR,
		BottomLeft:  bL,
		BottomRight: bR,
	}
}
