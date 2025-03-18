package main

import "fmt"

// 双向链表 Doubly Linked List

type Node struct {
	val  int
	prev *Node // 指向前驱节点
	next *Node // 指向后继节点
}
type NodeList struct {
	head *Node
}

// Insert 按照升序插入节点
func (n *NodeList) Insert(val int) {
	cur := n.head
	// 找到需要插入的位置
	for cur.next != nil && cur.next.val < val {
		cur = cur.next
	}
	// 插入新节点
	newNode := &Node{
		val:  val,
		prev: cur,
		next: cur.next,
	}
	// 更新后继节点的prev指针
	if cur.next != nil {
		cur.next.prev = newNode
	}
	cur.next = newNode
}

// Remove 删除节点
func (n *NodeList) Remove(val int) {
	cur := n.head
	for cur.next != nil && cur.next.val != val {
		cur = cur.next
	}
	if cur.next == nil {
		fmt.Println("没有找到该节点")
		return
	}
	// 更新后继节点的prev指针
	if cur.next.next != nil {
		cur.next.next.prev = cur
	}
	cur.next = cur.next.next
}

// Search 查找某个节点
func (n *NodeList) Search(val int) {
	cur := n.head
	for cur.next != nil && cur.next.val != val {
		cur = cur.next
	}
	if cur.next == nil {
		fmt.Println("没有找到该节点")
		return
	}
	fmt.Println("找到该节点")
}

// Print 打印双链表
func (n *NodeList) Print() {
	cur := n.head
	for cur != nil {
		fmt.Print(cur.val, "->")
		cur = cur.next
	}
	fmt.Println("结束")
}

func main() {
	list := &NodeList{
		// 定义一个头部伪节点，不包含实际值
		head: &Node{
			val:  -1,
			next: nil,
			prev: nil,
		},
	}
	list.Insert(1)
	list.Insert(100)
	list.Insert(2)
	list.Insert(3)
	list.Print()

	list.Remove(100)
	list.Search(1)

	list.Print()

}
