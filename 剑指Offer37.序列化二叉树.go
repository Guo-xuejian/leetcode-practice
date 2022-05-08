// 剑指 Offer 37. 序列化二叉树
// 请实现两个函数，分别用来序列化和反序列化二叉树。

// 你需要设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

// 提示：输入输出格式与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。

// 示例：

// 输入：root = [1,2,3,null,null,4,5]
// 输出：[1,2,3,null,null,4,5]

// package main

// import (
// 	. "nc_tools"
// 	"strconv"
// 	"strings"
// )

// import "fmt"

// import "unsafe"
/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param root TreeNode类
 * @return TreeNode类
 */
func dfs(root *TreeNode, str *string) {
	if root == nil {
		*str += "null,"
		return
	}
	*str += strconv.Itoa(root.Val) + ","
	dfs(root.Left, str)
	dfs(root.Right, str)
	return
}
func Serialize(root *TreeNode) string {
	// write code her
	if root == nil {
		return ""
	}
	ans := ""
	dfs(root, &ans)
	return ans
}
func Deserialize(s string) *TreeNode {
	// write code here
	if len(s) == 0 {
		return nil
	}
	strs := strings.Split(s, ",")
	return de_dfs(&strs)
}
func de_dfs(str *[]string) *TreeNode {
	if (*str)[0] == "null" {
		*str = (*str)[1:] //没办法go是值传递，你需要改变切片的长度，让构造可以进行下去
		return nil
	}
	nums, _ := strconv.Atoi((*str)[0])
	*str = (*str)[1:]
	root := &TreeNode{Val: nums}
	if len(*str) != 0 {
		root.Left = de_dfs(str)
	}
	if len(*str) != 0 {
		root.Right = de_dfs(str)
	}
	return root
}
