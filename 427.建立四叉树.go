/*
 * @lc app=leetcode.cn id=427 lang=golang
 *
 * [427] 建立四叉树
 */

// @lc code=start
/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

func construct(grid [][]int) *Node {
	if len(grid) == 0 {
		return nil
	}
	n := len(grid)
	var help func(x, y, off int) *Node
	help = func(x, y, off int) *Node {
		for i := x; i < x+off; i++ {
			for j := y; j < y+off; j++ {
				if grid[i][j] != grid[x][y] {
					return &Node{
						true,
						false,
						help(x, y, off/2),
						help(x, y+off/2, off/2),
						help(x+off/2, y, off/2),
						help(x+off/2, y+off/2, off/2),
					}
				}
			}
		}
		if grid[x][y] == 0 {
			return &Node{false, true, nil, nil, nil, nil}
		}
		return &Node{true, true, nil, nil, nil, nil}
	}
	return help(0, 0, n)
}

// @lc code=end

