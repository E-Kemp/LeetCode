package main

type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

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

func main() {
}

