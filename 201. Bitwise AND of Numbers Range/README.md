### Solution Retrospective

**Date:** 07/02/2024  
**Language:** Go  
**Difficulty:** Medium

I don't think I could've more royally misunderstood the capabilities of Go 😂

Looking at it fomr my academia glasses, this is a really simple problem where you shift the bits of two numbers until they stop matching, and return the result as a decimal. I didn't realise the `>>` operator **would do most of the heavy lifting for me**, so I instead converted the numbers into string representations of their binrary forms, and did some insane string manipulation to get the same result - a hacky workaround but it somehow was fast enough to beat ~80% of submissions in LeetCode...

My solution (hilariously wrong!):

```go
func rangeBitwiseAnd(left int, right int) int {
  leftBinary := strconv.FormatInt(int64(left), 2)
  rightBinary := strconv.FormatInt(int64(right), 2)
  binaryLen := len(leftBinary)

  if binaryLen < len(rightBinary) {
    leftBinary = strings.Repeat("0", len(rightBinary)-binaryLen) + leftBinary
    binaryLen = len(rightBinary)
  } else {
    rightBinary = strings.Repeat("0", len(leftBinary)-binaryLen) + rightBinary
  }

  result := ""
  fin := false
  for i := 0; i < binaryLen; i++ {
    if !fin && leftBinary[i] == rightBinary[i] {
      result = result + string(leftBinary[i])
    } else {
      result = result + "0"
      fin = true
    }
  }
  numResult, err := strconv.ParseInt(result, 2, 64)
  if err != nil {
    fmt.Println("Something went wrong")
    return 0
  }
  
  return int(numResult)
}
```

The better solution..:

```go
func fasterRangeBitwiseAnd(left int, right int) int {
  shift := 0
  for left != right {
    left >>= 1
    right >>= 1
    shift++
  }
  return left << shift
}
```
