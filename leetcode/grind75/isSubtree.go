package grind75

import (
	"errors"

	"algorithms.com/m/leetcode/data_structures"
)

// Given the roots of two binary trees root and subRoot, return true if there is a subtree of root
// with the same structure and node values of subRoot and false otherwise.

// A subtree of a binary tree tree is a tree that consists of a node in tree and all of this node's descendants.
//The tree tree could also be considered as a subtree of itself.

func IsSubtree(root *data_structures.TreeNode, subRoot *data_structures.TreeNode) bool {
	for root != nil {
		node, err := findNode(root, subRoot.Val)
		if err != nil {
			return IsSubtree(root.Left, subRoot) || IsSubtree(root.Right, subRoot)
		}

		return Traverse(node, subRoot)
	}

	return false

}

func Traverse(a *data_structures.TreeNode, b *data_structures.TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil && b != nil {
		return false
	}
	if a != nil && b == nil {
		return false
	}
	return a.Val == b.Val && Traverse(a.Left, b.Left) && Traverse(a.Right, b.Right)
}

func findNode(root *data_structures.TreeNode, val int) (*data_structures.TreeNode, error) {

	if root.Val == val {
		return root, nil
	}
	x, err := findNode(root.Left, val)

	y, err1 := findNode(root.Right, val)
	if err != nil && err1 != nil {
		return &data_structures.TreeNode{}, errors.New("error")
	}
	if err == nil {
		return x, nil
	}
	return y, nil
}
