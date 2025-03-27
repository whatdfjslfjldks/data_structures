package main

import "fmt"

type TreeNode struct {
	val    int
	left   *TreeNode
	right  *TreeNode
	height int
}

func (r *TreeNode) Insert(val int) *TreeNode {
	if r == nil {
		return &TreeNode{
			val:    val,
			height: 1,
		}
	}
	if val < r.val {
		r.left = r.left.Insert(val)
	} else {
		r.right = r.right.Insert(val)
	}
	// 更新当前树的高度
	r.height = 1 + max(r.left.CalHeight(), r.right.CalHeight())
	// 计算当前的平衡因子
	fac := r.CalBalanceFac()
	// LL操作
	if fac > 1 && val < r.left.val {
		return r.RotateRight()
	}
	// RR
	if fac < -1 && val > r.right.val {
		return r.RotateLeft()
	}
	// LR
	if fac > 1 && val > r.left.val {
		r.left = r.left.RotateLeft()
		return r.RotateRight()
	}
	// RL
	if fac < -1 && val < r.right.val {
		r.right = r.right.RotateRight()
		return r.RotateLeft()
	}

	return r
}
func (r *TreeNode) Remove(val int) *TreeNode {
	if r == nil {
		return nil
	}
	if val < r.val {
		r.left = r.left.Remove(val)
	} else if val > r.val {
		r.right = r.right.Remove(val)
	} else {
		// 情况1：没有子节点或只有一个子节点
		if r.left == nil {
			return r.right
		} else if r.right == nil {
			return r.left
		}
		// 情况2：有两个子节点，找到右子树的最小节点
		minNode := r.right
		for minNode.left != nil {
			minNode = minNode.left
		}
		// 替换当前节点的值为最小节点的值
		r.val = minNode.val
		// 删除右子树的最小节点
		r.right = r.right.Remove(minNode.val)
	}

	// 删除节点后更新当前树的高度
	r.height = 1 + max(r.left.CalHeight(), r.right.CalHeight())
	// 计算当前的平衡因子
	fac := r.CalBalanceFac()
	// LL操作
	if fac > 1 && r.left.CalBalanceFac() >= 0 {
		return r.RotateRight()
	}
	// RR
	if fac < -1 && r.right.CalBalanceFac() <= 0 {
		return r.RotateLeft()
	}
	// LR
	if fac > 1 && r.left.CalBalanceFac() < 0 {
		r.left = r.left.RotateLeft()
		return r.RotateRight()
	}
	// RL
	if fac < -1 && r.right.CalBalanceFac() > 0 {
		r.right = r.right.RotateRight()
		return r.RotateLeft()
	}

	return r
}

// RotateRight 右旋
func (r *TreeNode) RotateRight() *TreeNode {
	a := r.left
	b := a.right

	// 执行旋转
	a.right = r
	r.left = b

	a.height = max(a.left.CalHeight(), a.right.CalHeight()) + 1
	r.height = max(r.left.CalHeight(), r.right.CalHeight()) + 1

	return a
}

// RotateLeft 左旋
func (r *TreeNode) RotateLeft() *TreeNode {
	a := r.right
	b := a.left

	a.left = r
	r.right = b

	a.height = max(a.left.CalHeight(), a.right.CalHeight()) + 1
	r.height = max(r.left.CalHeight(), r.right.CalHeight()) + 1

	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (r *TreeNode) CalBalanceFac() int {
	if r == nil {
		return 0
	}
	return r.left.CalHeight() - r.right.CalHeight()
}
func (r *TreeNode) CalHeight() int {
	if r == nil {
		return 0
	}
	return r.height
}

func (r *TreeNode) InOrder() {
	if r != nil {
		r.left.InOrder()
		fmt.Println(r.val)
		r.right.InOrder()
	}
}
func (r *TreeNode) PreOrder() {
	if r != nil {
		fmt.Println(r.val)
		r.left.PreOrder()
		r.right.PreOrder()
	}
}
func main() {
	r := &TreeNode{
		val:    1,
		height: 1,
	}
	r = r.Insert(2)
	r = r.Insert(3)

	r.PreOrder()
}
