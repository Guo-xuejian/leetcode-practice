package main

import "fmt"

// 二叉树层次遍历，
type Tree struct {
	Val   int
	Left  *Tree
	Right *Tree
}

// 队列
func layerTravel(root *Tree) {
	if root == nil {
		fmt.Println("The Tree is nil")
		return
	}
	queue := []*Tree{}
	queue = append(queue, root)
	for len(queue) > 0 {
		temp := make([]*Tree, len(queue))
		copy(queue, temp)
		queue := []*Tree{}
		for i := 0; i < len(temp); i++ {
			fmt.Println(temp[i].Val)
			if temp[i].Left != nil {
				queue = append(queue, temp[i].Left)
			}
			if temp[i].Right != nil {
				queue = append(queue, temp[i].Right)
			}
		}
	}
}

// func main() {

// }
