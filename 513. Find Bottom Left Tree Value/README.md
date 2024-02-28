### Solution Retrospective

**Date:** 28/02/2024  
**Language:** Go  
**Difficulty:** Easy

This time we're doing a bit of breadth-first searching, utilizing a queue instead of a recursive function to search for the deepest left-most value:

```go
func findBottomLeftValue(root *TreeNode) int {
  row, queue := []*TreeNode{root}, []*TreeNode{}
  for true {
    queue = nil
    for _, v := range row {
      if v.Left != nil { queue = append(queue, v.Left) }
      if v.Right != nil { queue = append(queue, v.Right) }
    }
    if len(queue) == 0 {
      return row[0].Val
    } else {
      row = queue
    }
  }
  return 0
}
```

This took about 20 minutes to complete, and I feel like my Go skills are becoming more elegant as I complete these challenges. Perhaps it's time for me to start a proper project with Go? 👀
