package binarytree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostorderTraversalRecursive(t *testing.T) {
	tree := GetTree()
	expect := []int{7, 4, 5, 2, 6, 3, 1}
	actual := postorderTraverse(tree)
	assert.ElementsMatch(t, expect, actual)
}

func postorderTraverse(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	left := postorderTraverse(root.Left)
	right := postorderTraverse(root.Right)

	ans := make([]int, 0)
	ans = append(ans, left...)
	ans = append(ans, right...)
	ans = append(ans, root.Val)

	return ans
}
