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

// leetcode427_建立四叉树
func construct(grid [][]int) *Node {
	n := len(grid)
	return dfs(grid, 0, 0, n, n)
}

func dfs(grid [][]int, x1, y1, x2, y2 int) *Node {
	isLeaf := true
	for i := x1; i < x2; i++ {
		for j := y1; j < y2; j++ {
			if grid[i][j] != grid[x1][y1] {
				isLeaf = false
				break
			}
		}
	}
	if isLeaf == true {
		return &Node{
			Val:    grid[x1][y1] == 1,
			IsLeaf: true,
		}
	}
	midX := (x1 + x2) / 2
	midY := (y1 + y2) / 2
	tL := dfs(grid, x1, y1, midX, midY)
	tR := dfs(grid, x1, midY, midX, y2)
	bL := dfs(grid, midX, y1, x2, midY)
	bR := dfs(grid, midX, midY, x2, y2)
	return &Node{
		Val:         false,
		IsLeaf:      false,
		TopLeft:     tL,
		TopRight:    tR,
		BottomLeft:  bL,
		BottomRight: bR,
	}
}
