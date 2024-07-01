package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1,
		Left: nil,
		Right: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
			Right: nil}}
	ans := levelOrder(root)
	fmt.Println("=======")
	fmt.Println(ans)
}

func inorderTraversal(root *TreeNode) []int {
	queue := make([]*TreeNode, 0)
	ans := make([]int, 0)
	queue = append(queue, root)
	fmt.Printf("queue len: %v \r\n", len(queue))
	for {
		// pop if left is nil
		cur := queue[len(queue)-1]
		if cur.Left == nil {
			queue = queue[:len(queue)-1]
			fmt.Printf("-- queue len: %v \r\n", len(queue))
			ans = append(ans, cur.Val)
			fmt.Println(ans)
			// push right
			if cur.Right != nil {
				queue = append(queue, cur.Right)
				cur.Right = nil
				fmt.Printf("** queue len: %v \r\n", len(queue))
			}
			if len(queue) == 0 {
				fmt.Println("should break")
				break
			}
		} else {
			// push left
			if cur.Left != nil {
				queue = append(queue, cur.Left)
				cur.Left = nil
				fmt.Printf("@@ queue len: %v \r\n", len(queue))
			}
		}

		if len(queue) == 0 {
			break
		}
	}
	return ans
}

func levelOrder(root *TreeNode) [][]int {
	stack := []*Node{NewNode(root, 0)}
	ans := make([][]int, 0)
	for len(stack) > 0 {
		// pop node from head
		cur := stack[0]
		stack = stack[1:]

		// add to ans
		if len(ans) < cur.Level+1 {
			ans = append(ans, []int{})
		}

		ans[cur.Level] = append(ans[cur.Level], cur.Tree.Val)
		fmt.Printf("level: %v, ans: %v\r\n", cur.Level, ans)

		// push childs
		if cur.Tree.Left != nil {
			stack = append(stack, NewNode(cur.Tree.Left, cur.Level+1))
		}
		if cur.Tree.Right != nil {
			stack = append(stack, NewNode(cur.Tree.Right, cur.Level+1))
		}
	}
	return ans
}

type Node struct {
	Tree  *TreeNode
	Level int
}

func NewNode(tree *TreeNode, level int) *Node {
	return &Node{
		Tree:  tree,
		Level: level,
	}
}
