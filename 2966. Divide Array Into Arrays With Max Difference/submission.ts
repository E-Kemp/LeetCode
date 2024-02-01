function divideArray(nums: number[], k: number): number[][] {
    if(nums.length % 3 != 0) return []
    nums.sort((a, b) => a - b)
    const output: number[][] = []
    for(var i = 0; i < nums.length; i += 3) {
        if(nums[i + 2] - nums[i] > k) return []
        output.push([nums[i], nums[i + 1], nums[i + 2]])
    }
    return output
};