package main

import (
	"fmt"
	"math/rand"
)

const (
	MaxLevel    = 5
	Probability = 0.5
)

type Node struct {
	val  int
	next *Node
	down *Node
}

type SkipList struct {
	head  *Node
	level int // 当前所在最大层级
}

// RandomLevel 获取插入元素的最大层级
func RandomLevel() int {
	level := 1
	for level < MaxLevel && rand.Float64() < Probability {
		level++
	}
	return level
}

// Insert 向跳表中插入一个元素
func (s *SkipList) Insert(val int) {
	level := RandomLevel()
	// 如果当前元素的最大层级大于当前跳表的最大层级，更新跳表的最大层级
	if level > s.level {
		for i := s.level; i < level; i++ {
			// 创建一个新的头节点
			newHead := &Node{
				val:  -1,
				next: nil,
				down: s.head,
			}
			s.head = newHead
		}
		s.level = level
	}
	// 从跳表的顶部开始，准备插入新的节点
	cur := s.head
	var update []*Node
	for i := 0; i < s.level; i++ {
		for cur.next != nil && cur.next.val < val {
			cur = cur.next
		}
		update = append([]*Node{cur}, update...)
		if cur.down != nil {
			cur = cur.down
		}
	}
	// 从底部向上遍历跳表，添加新的元素
	var down *Node
	for i := 0; i < level; i++ {
		// 定义新节点
		newNode := &Node{
			val:  val,
			next: update[i].next,
			down: down,
		}
		update[i].next = newNode
		down = newNode
	}
}

func (s *SkipList) Search(val int) bool {
	cur := s.head
	for cur != nil {
		for cur.next != nil && cur.next.val < val {
			cur = cur.next
		}
		if cur.next != nil && cur.next.val == val {
			return true
		}
		cur = cur.down
	}
	return false
}

func (s *SkipList) Remove(val int) {
	cur := s.head
	var remove []*Node

	for cur != nil {
		for cur.next != nil && cur.next.val < val {
			cur = cur.next
		}
		remove = append([]*Node{cur}, remove...)
		cur = cur.down
	}

	// 删除操作
	for i := 0; i < len(remove); i++ {
		if remove[i].next != nil && remove[i].next.val == val {
			remove[i].next = remove[i].next.next
		}
	}
	// 跳表的最大层级进行更新
	for s.head != nil && s.head.next == nil {
		s.head = s.head.down
		s.level--
	}

}

func (s *SkipList) Print() {
	cur := s.head
	for i := s.level; i > 0; i-- {
		fmt.Print("当前层级： ", i)
		tmp := cur
		for tmp.next != nil {
			fmt.Print(" ", tmp.next.val, "->")
			tmp = tmp.next
		}
		fmt.Println("nil")
		cur = cur.down
	}
}

func main() {
	list := SkipList{
		head: &Node{
			val:  -1,
			next: nil,
			down: nil,
		},
		level: 1,
	}
	list.Insert(1)
	list.Insert(8)
	list.Insert(2)
	list.Insert(3)

	result := list.Search(1)
	fmt.Println("第一次跳表结构")
	list.Print()
	fmt.Println("第一次查找结构", result)
	list.Remove(1)
	fmt.Println("第二次跳表结构")
	list.Print()
	result = list.Search(1)
	fmt.Println("第二次查找结构", result)

}
