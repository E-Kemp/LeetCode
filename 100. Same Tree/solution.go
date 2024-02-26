package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// check inputs
	if p == nil || q == nil {
		return p == nil && q == nil
	}

	// check int
	if p.Val != q.Val {
		return false
	}

	// check left side
	leftCheck := false
	if p.Left != nil && q.Left != nil {
		leftCheck = isSameTree(p.Left, q.Left)
	} else if p.Left == nil && q.Left == nil {
		leftCheck = true
	}

	// check left side
	rightCheck := false
	if p.Right != nil && q.Right != nil {
		rightCheck = isSameTree(p.Right, q.Right)
	} else if p.Right == nil && q.Right == nil {
		rightCheck = true
	}

	return leftCheck && rightCheck
}

func main() {
}

