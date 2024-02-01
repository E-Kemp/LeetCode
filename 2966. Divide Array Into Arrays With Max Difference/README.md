### Solution Retrospective

My second solution in the project, this one took me down a rabbit hole that led to a stupidly simple solution...

The problem of dividing array into sub-arrays with a max difference at first felt like a sorting algorithm - I should've stuck with that. I first tried making a greedy algorithm that would loop through, swap values and re-order the array until I got what I wanted. Then I realised this should be much easier, I could just sort the array and check the groupings don't have a difference bigger than `k` provided. Not only that, with groupings of three, if the first and third values match the criteria, then the second would inherently fit too as it would sit between the first and third values.

This was my solution (this time in TypeScript, as I miss-assumed it would be harder, so I wanted to use a familiar language...)

```typescript
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
```

I could condense this further, but that wouldn't achieve much except make it less readable.
