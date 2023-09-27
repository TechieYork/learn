package main

import "fmt"

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

type TreeNodeStack struct {
	Node  *TreeNode
	Level int
}

func NewTreeNodeStack(node *TreeNode, level int) *TreeNodeStack {
	return &TreeNodeStack{
		Node:  node,
		Level: level,
	}
}

func main() {
	demoRoot := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Println("====== Depth first traversal using stack ======")
	depthFirstTraverse(demoRoot, 0)
	fmt.Println("====== Depth first traversal using recursive ======")
	depthFirstTraverseRecursive(demoRoot, 0)
	fmt.Println("====== Width first traversal using stack ======")
	widthFirstTraverse(demoRoot, 0)
}

func depthFirstTraverseRecursive(node *TreeNode, level int) {
	if node == nil {
		return
	}
	fmt.Printf("level=%v, value=%v\r\n", level, node.Val)
	depthFirstTraverseRecursive(node.Right, level+1)
	depthFirstTraverseRecursive(node.Left, level+1)
}

func depthFirstTraverse(node *TreeNode, level int) {
	if node == nil {
		return
	}
	stack := NewStack(NewTreeNodeStack(node, level))
	for !stack.IsEmpty() {
		item := stack.PopBack().(*TreeNodeStack)
		if item.Node.Left != nil {
			stack.PushBack(NewTreeNodeStack(item.Node.Left, item.Level+1))
		}
		if item.Node.Right != nil {
			stack.PushBack(NewTreeNodeStack(item.Node.Right, item.Level+1))
		}
		fmt.Printf("level=%v, value=%v\r\n", item.Level, item.Node.Val)
	}
}

func widthFirstTraverse(node *TreeNode, level int) {
	if node == nil {
		return
	}
	stack := NewStack(NewTreeNodeStack(node, level))
	for !stack.IsEmpty() {
		item := stack.PopFront().(*TreeNodeStack)
		if item.Node.Left != nil {
			stack.PushBack(NewTreeNodeStack(item.Node.Left, item.Level+1))
		}
		if item.Node.Right != nil {
			stack.PushBack(NewTreeNodeStack(item.Node.Right, item.Level+1))
		}
		fmt.Printf("level=%v, value=%v\r\n", item.Level, item.Node.Val)
	}
}

type Stack struct {
	data []interface{}
}

func NewStack(data ...interface{}) *Stack {
	if len(data) > 0 {
		return &Stack{
			data: data,
		}
	}
	return &Stack{
		data: make([]interface{}, 0),
	}
}

func (q *Stack) PushBack(item interface{}) {
	q.data = append(q.data, item)
}

func (q *Stack) PopFront() interface{} {
	item := q.data[0]
	q.data = q.data[1:]
	return item
}

func (q *Stack) PopBack() interface{} {
	item := q.data[len(q.data)-1]
	q.data = q.data[:len(q.data)-1]
	return item
}

func (q *Stack) IsEmpty() bool {
	return len(q.data) == 0
}
