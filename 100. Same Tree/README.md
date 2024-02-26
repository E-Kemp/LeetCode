### Solution Retrospective

**Date:** 26/02/2024  
**Language:** Go  
**Difficulty:** Easy

A lovely DFS search to compare 2 binary trees, this solution may not be extremely elegant, but I believe it is readable and it ran just as fast as the shorthand solutions! Yipee for languages that compile and run stupidly fast!

```go
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
```

A pretty quick and dirty DFS search, maybe it could have been BFS but they both have advantages and disadvantages... Have a good day!
