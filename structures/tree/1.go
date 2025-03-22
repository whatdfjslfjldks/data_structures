package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Insert 插入一个元素
func (root *TreeNode) Insert(val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = root.Left.Insert(val)
	} else {
		root.Right = root.Right.Insert(val)
	}
	return root
}
func (root *TreeNode) Remove(val int) *TreeNode {
	if root == nil {
		return root
	}
	if val < root.Val {
		root.Left = root.Left.Remove(val)
	} else if val > root.Val {
		root.Right = root.Right.Remove(val)
	} else {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		minNode := root.Right
		for minNode.Left != nil {
			minNode = minNode.Left
		}
		root.Val = minNode.Val
	}
	return root
}
func (root *TreeNode) Search(val int) *TreeNode {
	if root == nil {
		return nil
	}
	if val < root.Val {
		return root.Left.Search(val)
	} else if val > root.Val {
		return root.Right.Search(val)
	} else {
		return root
	}
}

// InOrder 中序遍历
func (root *TreeNode) InOrder() {
	if root != nil {
		root.Left.InOrder() // 先遍历左子树
		fmt.Printf("%d ", root.Val)
		root.Right.InOrder() // 再遍历右子树
	}
}

// PreOrder 前序遍历
func (root *TreeNode) PreOrder() {
	if root != nil {
		fmt.Printf("%d ", root.Val) // 先访问根节点
		root.Left.PreOrder()        // 再遍历左子树
		root.Right.PreOrder()       // 最后遍历右子树
	}
}

// PostOrder 后序遍历
func (root *TreeNode) PostOrder() {
	if root != nil {
		root.Left.PostOrder()       // 先遍历左子树
		root.Right.PostOrder()      // 再遍历右子树
		fmt.Printf("%d ", root.Val) // 最后访问根节点
	}
}

func main() {
	var root *TreeNode = nil
	root = root.Insert(1)
	root = root.Insert(2)
	root = root.Insert(4)
	root = root.Remove(3)

	root.InOrder()
}
