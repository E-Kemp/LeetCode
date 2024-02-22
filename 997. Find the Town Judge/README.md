### Solution Retrospective

**Date:** 07/02/2024  
**Language:** Go  
**Difficulty:** Medium

Today's problem involved two calculations over an array of number pairs (to determine who the judge is). I went for the solution that first came to mind, and used one set of calculations to narrow down the number of second calculations rather than run both calculations in the same loop. My method used far less memory as a result but took about 30% more time to complete than the faster solutions.

My solution:

```go
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

  // test whether the candidate is trusted by everyone
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
```

Looking back, I could have made a record for each candidate: How many people trust them, and how many people they _don't_ trust, then return that candidate if the numbers equal _n_, but anyway enjoy my quick and dirty solution!
