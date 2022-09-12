package bst

import (
	"math"
)

// InitBST creates a bst from an array
func InitBST(values []int) *TreeNode {
	if len(values) == 0 {
		return nil
	}
	var root *TreeNode
	for i := 0; i < len(values); i++ {
		root = Insert(root, values[i])
	}
	return root
}

// Insert inserts a new node to bst
func Insert(root *TreeNode, val int) *TreeNode {
	// exists, return
	if Exist(root, val) {
		return root
	}

	n := &TreeNode{Val: val}

	// empty tree
	if root == nil {
		return n
	}

	p := root
	for p != nil {
		if val < p.Val {
			// search in left
			if p.Left == nil {
				p.Left = n
				break
			}
			p = p.Left
		} else {
			// search in right
			if p.Right == nil {
				p.Right = n
				break
			}
			p = p.Right
		}
	}
	return root
}

// InorderTraversal outputs the binary search tree in ascending order
func InorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int

	res = append(res, InorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, InorderTraversal(root.Right)...)

	return res
}

// Exist checks whether a node is in BST
func Exist(root *TreeNode, val int) bool {
	if root == nil {
		return false
	}
	if root.Val == val {
		return true
	}

	if val < root.Val {
		return Exist(root.Left, val)
	}

	return Exist(root.Right, val)
}

// Delete deletes a node from bst if exists
func Delete(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		// case 1: root is leaf root
		if root.Left == nil && root.Right == nil {
			return nil
		}

		// case 2: root has a node
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		// case 3: root has two nodes
		// the minimum of Right -> root	(choose it in this code)
		// or
		// the maximum of Left -> root
		pre := root
		cur := root.Right
		for cur.Left != nil {
			pre = cur
			cur = cur.Left
		}
		pre.Right = cur.Right
		cur.Left = root.Left
		cur.Right = root.Right
		root = cur
		return root
	}

	if val < root.Val {
		root.Left = Delete(root.Left, val)
	}
	if val > root.Val {
		root.Right = Delete(root.Right, val)
	}
	return root
}

// IsBST checks whether a tree is a bst
func IsBST(root *TreeNode) bool {
	return isBST(root, math.MinInt64, math.MaxInt64)
}

func isBST(root *TreeNode, lower, upper int64) bool {
	if root == nil {
		return true
	}

	if int64(root.Val) <= lower || int64(root.Val) >= upper {
		return false
	}

	return isBST(root.Left, lower, int64(root.Val)) && isBST(root.Right, int64(root.Val), upper)
}
