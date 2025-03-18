package main

import "fmt"

// 循环链表，设置头节点为哨兵节点

type Node struct {
	val  int
	next *Node
}
type NodeList struct {
	head *Node
}

func (n *NodeList) Insert(val int) {
	cur := n.head
	for cur.next != n.head && cur.next.val < val {
		cur = cur.next
	}
	newNode := &Node{
		val:  val,
		next: cur.next,
	}
	cur.next = newNode
}

func (n *NodeList) Search(val int) bool {
	cur := n.head
	for cur.next != n.head && cur.next.val < val {
		cur = cur.next
	}
	if cur.next != n.head && cur.next.val == val {
		return true
	} else {
		return false
	}
}

func (n *NodeList) Remove(val int) {
	cur := n.head
	for cur.next != n.head && cur.next.val < val {
		cur = cur.next
	}
	if cur.next != n.head && cur.next.val == val {
		cur.next = cur.next.next
	} else {
		fmt.Println("没有找到要删除的节点")
	}
}

func (n *NodeList) Print() {
	cur := n.head
	for cur.next != n.head {
		cur = cur.next
		fmt.Print(cur.val, "->")
	}
	fmt.Println("结束")
}

func main() {
	// 创建一个哨兵头节点，并让他指向自己
	sentinel := &Node{
		val:  -1,
		next: nil,
	}
	sentinel.next = sentinel
	list := &NodeList{
		head: sentinel,
	}
	list.Insert(1)
	list.Insert(2)
	list.Insert(5)
	list.Insert(3)
	list.Remove(1)

	a := list.Search(5)
	fmt.Println(a)

	list.Print()
}
