package grind75

// Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST.

// According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

// So nearest common parent
// we are traversing from top to bottom
// we already have the node's address for p and q
// .left or .right = p or q
// Once that is done , we will backtrack? how ?
// we just push the data into a new data structure and then  thats it
// So p or q itself can be the common parent .

// Last common value ???
// 1,2,3,4,6,7,9,100
// 5,6,7,8,11,12

type TreeNode struct {
	Val   int
	Left  TreeNode
	Right TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}
	// basically we are iterating till we can iterate toward p and q together
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root

}
