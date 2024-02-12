### Solution Retrospective

**Date:** 12/02/2024  
**Language:** Rust  
**Difficulty:** Easy peasy

Wanting to get back into Rust, today's problem was a nice introduciton to a couple of concepts (despite the solution to today's problem being less than 10 lines long).

It's quite comforting to have some native array manipulation functions, I'm not sure if it will be as extensive as Javascript's array manipulation capabilities but at least having `.sort()` is a good start!

The `usize` type can't be treated as a usual integer - we can't apply some maths operators to it. Instead, it has a set of its own functions for manipulating, hence the use of `div_ceil()` allowing us to divide and round up to the next (not closest) whole number.

Overall, today was a nice quick challenge, applying a trick to find the majority element where it's guranteed that it will take over half of the array. Following is the code, have a good day!

```rust
impl Solution {
    pub fn majority_element(nums: Vec<i32>) -> i32 {
        let mut sort_nums = nums.clone();
        sort_nums.sort();
        return sort_nums[sort_nums.len().div_ceil(2) - 1];
    }
}
```
