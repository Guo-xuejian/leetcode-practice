/*
 * @lc app=leetcode.cn id=331 lang=golang
 *
 * [331] 验证二叉树的前序序列化
 */

// @lc code=start
func isValidSerialization(preorder string) bool {
	length := len(preorder)
	// 每一个非空节点都有两个空节点的空间
	// 根节点特殊，只有一个，因为它自己就占了一个
	slots := 1
	for i := 0; i < length; {
		if slots == 0 { // 循环内无空节点空间，退出
			return false
		}
		if preorder[i] == ',' { // 连接符
			i++
		} else if preorder[i] == '#' { // 空节点占用
			i++
			slots--
		} else { // 非空节点出现
			// 由于数字可能不止一位
			for i < length && preorder[i] != ',' {
				i++
			}
			// 非空节点自带的空节点空间加入
			slots++ // slots = slots - 1 + 2
		}
	}
	// 最终的空节点不能有剩余
	return slots == 0
}

// @lc code=end

