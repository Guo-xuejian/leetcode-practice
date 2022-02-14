// 剑指 Offer 33. 二叉搜索树的后序遍历序列
// 输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。

 

// 参考以下这颗二叉搜索树：

//      5
//     / \
//    2   6
//   / \
//  1   3
// 示例 1：

// 输入: [1,6,3,2,5]
// 输出: false
// 示例 2：

// 输入: [1,3,2,6,5]
// 输出: true
 

// 提示：

// 数组长度 <= 1000


// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了100.00%的用户
func verifyPostorder(postorder []int) bool {
     // 始终将最后一个节点当作根节点，递归判定左右子树和当前根节点的上下限关系
     var buildBST func(up, down int)
     buildBST = func(up, down int) {
         currLength := len(postorder)
         if currLength == 0 {    // 已无节点则构建 BST 成功，退出
             return
         }
         currVal := postorder[currLength - 1]
         if !(currVal > down && currVal < up) {  // 不符合上下限关系退出
             return
         }
         postorder = postorder[:currLength - 1]  // 根节点结束使命
         buildBST(up, currVal)       // 根节点是右子节点的下限
         buildBST(currVal, down)     // 左子节点的上限
     }
     buildBST(math.MaxInt32, -math.MaxInt32)
     if len(postorder) == 0 {
         return true
     }
     return false
 }