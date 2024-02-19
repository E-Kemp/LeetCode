pub struct Solution {}

impl Solution {
    pub fn majority_element(nums: Vec<i32>) -> i32 {
        let mut sort_nums = nums.clone();
        sort_nums.sort();
        return sort_nums[sort_nums.len().div_ceil(2) - 1];
    }
}