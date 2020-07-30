package main

type TreeNode struct {
	    Val int
	    Left *TreeNode
	    Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}else{
		max(maxDepth(root.Left), maxDepth(root.right)) + 1
	}
}

