/*
 * @lc app=leetcode.cn id=587 lang=golang
 *
 * [587] 安装栅栏
 */

// @lc code=start
func cross(p, q, r []int) int {
	return (q[0]-p[0])*(r[1]-q[1]) - (q[1]-p[1])*(r[0]-q[0])
}

func outerTrees(trees [][]int) (ans [][]int) {
	n := len(trees)
	if n < 4 {
		return trees
	}

	leftMost := 0
	for i, tree := range trees {
		if tree[0] < trees[leftMost][0] {
			leftMost = i
		}
	}

	vis := make([]bool, n)
	p := leftMost
	for {
		q := (p + 1) % n
		for r, tree := range trees {
			// 如果 r 在 pq 的右侧，则 q = r
			if cross(trees[p], trees[q], tree) < 0 {
				q = r
			}
		}
		// 是否存在点 i, 使得 p q i 在同一条直线上
		for i, b := range vis {
			if !b && i != p && i != q && cross(trees[p], trees[q], trees[i]) == 0 {
				ans = append(ans, trees[i])
				vis[i] = true
			}
		}
		if !vis[q] {
			ans = append(ans, trees[q])
			vis[q] = true
		}
		p = q
		if p == leftMost {
			break
		}
	}
	return
}

// @lc code=end

