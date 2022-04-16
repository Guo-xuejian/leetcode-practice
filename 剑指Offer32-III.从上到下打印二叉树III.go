// 剑指 Offer 32 - III. 从上到下打印二叉树 III
// 请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。

 

// 例如:
// 给定二叉树: [3,9,20,null,null,15,7],

//     3
//    / \
//   9  20
//     /  \
//    15   7
// 返回其层次遍历结果：

// [
//   [3],
//   [20,9],
//   [15,7]
// ]
 

// 提示：

// 节点总数 <= 1000



/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func levelOrder(root *TreeNode) (res [][]int) {
  if root == nil {
      return
  }
  queue := []*TreeNode{root}
  flag := true
  for len(queue) > 0 {
      temp := make([]*TreeNode, len(queue))
      copy(temp, queue)
      queue = []*TreeNode{}
      curr := []int{}
      if flag {
          for _, node := range temp {
              curr = append(curr, node.Val)
              if node.Left != nil {
                  queue = append(queue, node.Left)
              }
              if node.Right != nil {
                  queue = append(queue, node.Right)
              }
          }
      } else {
          for i := len(temp) - 1; i >= 0; i-- {
              curr = append(curr, temp[i].Val)
              if temp[i].Right != nil {
                  queue = append(queue, temp[i].Right)
              }
              if temp[i].Left != nil {
                  queue = append(queue, temp[i].Left)
              }
          }
          for i := 0; i < len(queue) / 2; i++ {
              queue[i], queue[len(queue) - 1- i] = queue[len(queue) - 1- i], queue[i]
          }
      }
      flag = !flag
      res = append(res, curr)
  }
  return
}