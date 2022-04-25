package binarytree

import (
	"fmt"
	"testing"
)

func TestBinaryTreePreorderTraversalRecursive(t *testing.T) {
	tree := GetTree()
	var ans []int
	traverse(tree, &ans)
	fmt.Println(ans)
}

func traverse(root *TreeNode, a *[]int) {
	if root != nil {
		*a = append(*a, root.Val)
		traverse(root.Left, a)
		traverse(root.Right, a)
	}
}
