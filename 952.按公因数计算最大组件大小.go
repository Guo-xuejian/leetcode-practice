/*
 * @lc app=leetcode.cn id=952 lang=golang
 *
 * [952] 按公因数计算最大组件大小
 *
 * https://leetcode.cn/problems/largest-component-size-by-common-factor/description/
 *
 * algorithms
 * Hard (37.60%)
 * Likes:    78
 * Dislikes: 0
 * Total Accepted:    4.4K
 * Total Submissions: 11.3K
 * Testcase Example:  '[4,6,15,35]'
 *
 * 给定一个由不同正整数的组成的非空数组 nums ，考虑下面的图：
 *
 *
 * 有 nums.length 个节点，按从 nums[0] 到 nums[nums.length - 1] 标记；
 * 只有当 nums[i] 和 nums[j] 共用一个大于 1 的公因数时，nums[i] 和 nums[j]之间才有一条边。
 *
 *
 * 返回 图中最大连通组件的大小 。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：nums = [4,6,15,35]
 * 输出：4
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：nums = [20,50,9,63]
 * 输出：2
 *
 *
 * 示例 3：
 *
 *
 *
 *
 * 输入：nums = [2,3,6,7,4,12,21,39]
 * 输出：8
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 2 * 10^4
 * 1 <= nums[i] <= 10^5
 * nums 中所有值都 不同
 *
 *
 */

// @lc code=start
type unionFind struct {
	parent, rank []int
}

func newUnionFind(n int) unionFind {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return unionFind{parent, make([]int, n)}
}

func (uf unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf unionFind) merge(x, y int) {
	x, y = uf.find(x), uf.find(y)
	if x == y {
		return
	}
	if uf.rank[x] > uf.rank[y] {
		uf.parent[y] = x
	} else if uf.rank[x] < uf.rank[y] {
		uf.parent[x] = y
	} else {
		uf.parent[y] = x
		uf.rank[x]++
	}
}

func largestComponentSize(nums []int) (ans int) {
	m := 0
	for _, num := range nums {
		m = max(m, num)
	}
	uf := newUnionFind(m + 1)
	for _, num := range nums {
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				uf.merge(num, i)
				uf.merge(num, num/i)
			}
		}
	}
	cnt := make([]int, m+1)
	for _, num := range nums {
		rt := uf.find(num)
		cnt[rt]++
		ans = max(ans, cnt[rt])
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// @lc code=end

