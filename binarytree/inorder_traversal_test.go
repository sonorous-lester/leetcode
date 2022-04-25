package binarytree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInorderTraversalRecursive(t *testing.T) {
	expect := []int{7, 4, 2, 5, 1, 3, 6}
	tree := GetTree()
	ans := inorderTraverse(tree)
	assert.ElementsMatchf(t, expect, ans, "answer should be: %v", expect)
}

func inorderTraverse(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}

	ans := make([]int, 0)
	left := inorderTraverse(root.Left)
	right := inorderTraverse(root.Right)

	ans = append(ans, left...)
	ans = append(ans, root.Val)
	ans = append(ans, right...)

	return ans
}
