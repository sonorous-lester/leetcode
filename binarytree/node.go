package binarytree

/*
This tree be like:
					1
				2		3
			  4   5   X   6
			7
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GetTree() *TreeNode {
	root := &TreeNode{1, nil, nil}
	root.Left = &TreeNode{2, nil, nil}
	root.Right = &TreeNode{3, nil, nil}
	root.Left.Left = &TreeNode{4, nil, nil}
	root.Left.Right = &TreeNode{5, nil, nil}
	root.Right.Right = &TreeNode{6, nil, nil}
	root.Left.Left.Left = &TreeNode{7, nil, nil}
	return root
}
