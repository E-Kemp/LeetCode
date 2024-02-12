mod solution;

fn main() {
    let a = vec![3, 2, 3];
    let b = vec![2, 2, 1, 1, 1, 2, 2];
    let c = vec![1];
    assert_eq!(solution::Solution::majority_element(a), 3);
    assert_eq!(solution::Solution::majority_element(b), 2);
    assert_eq!(solution::Solution::majority_element(c), 1);
}
