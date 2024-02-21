package main

import (
	"fmt"
	"strconv"
	"strings"
)

// My convoluted version using some whacky string manipulation

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

func ShouldPassBitwiseTest(left int, right int, testCase int) error {
  if x := rangeBitwiseAnd(left, right); x != testCase {
    return fmt.Errorf("fail %d, %d. Expectd %d, got %d", left, right, testCase, x)
  }
  return nil
}

// Actual solution, using nerdy bitwise operators that I haven't looked at since university...

func fasterRangeBitwiseAnd(left int, right int) int {
  shift := 0
  for left != right {
    left >>= 1
    right >>= 1
    shift++
  }
  return left << shift
}

func ShouldPassFasterBitwiseTest(left int, right int, testCase int) error {
  if x := fasterRangeBitwiseAnd(left, right); x != testCase {
    return fmt.Errorf("fail %d, %d. Expectd %d, got %d", left, right, testCase, x)
  }
  return nil
}

func main() {
  errored := false
  if err := ShouldPassBitwiseTest(5, 7, 4); err != nil {
    fmt.Println(err.Error())
    errored = true
  }
  if err := ShouldPassBitwiseTest(0, 0, 0); err != nil {
    fmt.Println(err.Error())
    errored = true
  }
  if err := ShouldPassBitwiseTest(1, 2147483647, 0); err != nil {
    fmt.Println(err.Error())
    errored = true
  }

  if err := ShouldPassFasterBitwiseTest(5, 7, 4); err != nil {
    fmt.Println(err.Error())
    errored = true
  }
  if err := ShouldPassFasterBitwiseTest(0, 0, 0); err != nil {
    fmt.Println(err.Error())
    errored = true
  }
  if err := ShouldPassFasterBitwiseTest(1, 2147483647, 0); err != nil {
    fmt.Println(err.Error())
    errored = true
  }

  if !errored {
    fmt.Println("All tests passed")
  }
}