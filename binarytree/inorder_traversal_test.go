package binarytree

import (
	"fmt"
	"testing"
)

func TestInorderTraversalRecursive(t *testing.T) {
	tree := GetTree()
	ans := inorderTraverse(tree)
	fmt.Println(ans)
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
