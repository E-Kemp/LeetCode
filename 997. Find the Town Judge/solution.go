package main

import (
	"fmt"
	"slices"
)

type testCase struct {
  n int
  trust [][]int
}

func findJudge(n int, trust [][]int) int {
  candidates := []int{}

  if len(trust) == 0 && n == 1 {
    return 1
  }
  
  // get judge candidates
  for i := 0; i <= n; i++ {
    if !slices.ContainsFunc(trust, func(v []int) bool {
      return v[0] == i
    }) {
      candidates = append(candidates, i)
    }
  }

  if len(candidates) == 0 {
    return -1
  }

  for _, c := range candidates {
    trustCount := 0
    for _, t := range trust {
      if t[1] == c {
        trustCount++
      }
    }
    if trustCount == n-1 {
      return c
    }
  }

  return -1
}

func ShouldPassFindJudgeTest(testCase testCase) error {
  judge := findJudge(testCase.n, testCase.trust)
  if judge != testCase.n {
    return fmt.Errorf("findJudge(%d, %v) should be %d", testCase.n, testCase.trust, testCase.n)
  }
  return nil
}

func main() {
  case1 := testCase{n: 2, trust: [][]int{{1, 2}}}
  case2 := testCase{n: 3, trust: [][]int{{1, 3}, {2, 3}}}
  case3 := testCase{n: 3, trust: [][]int{{1, 3}, {2, 3}, {3, 1}}}

  errored := false
  if err := ShouldPassFindJudgeTest(case1); err != nil {
    fmt.Println(err.Error())
    errored = true
  }
  if err := ShouldPassFindJudgeTest(case2); err != nil {
    fmt.Println(err.Error())
    errored = true
  }
  if err := ShouldPassFindJudgeTest(case3); err != nil {
    fmt.Println(err.Error())
    errored = true
  }

  if !errored {
    fmt.Println("All tests passed")
  }
}