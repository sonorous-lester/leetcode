package binarytree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryTreePreorderTraversalRecursive(t *testing.T) {
	expect := []int{1, 2, 4, 7, 5, 3, 6}
	tree := GetTree()
	var ans []int
	traverse(tree, &ans)
	assert.ElementsMatch(t, expect, ans, "answer should be: %v", expect)
}

func traverse(root *TreeNode, a *[]int) {
	if root != nil {
		*a = append(*a, root.Val)
		traverse(root.Left, a)
		traverse(root.Right, a)
	}
}
