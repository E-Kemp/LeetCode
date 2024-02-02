### Solution Retrospective

**Date:** 02/02/2024  
**Language:** C#  
**Difficulty:** Medium

This was quite a nice problem to brush up my C# skills. A bit of string array manipulation and a bit of thought into a loop condition, and we have ourselves a nice Leetcode problem! There's not much to say except the code pretty much explains itself...

```cs
public class Solution {
    public IList<int> SequentialDigits(int low, int high) {
        string NUMS = "123456789";
        int lowNumDigits = low.ToString().Length;
        int highNumDigits = high.ToString().Length;
        IList<int> output = new List<int>();
        
        for(int numDigits = lowNumDigits; numDigits <= highNumDigits; numDigits++) {
            bool complete = false;

            // iterations = (length + 1) - size
            for(int index = 0; index < (NUMS.Length + 1) - numDigits && !complete; index++) {
                int num = Int32.Parse(NUMS.Substring(index, numDigits));
                if(low <= num && num <= high) {
                    output.Add(num);
                }
                else if(num > high) {
                    complete = true;
                }
            }
        }
        return output;
    }
}
```

Maybe there's some optimisations to be made here or there, but I don't want to spend all night improving a solution that is apparently faster than _most_ other solutions people wrote up, where the key difference seems to be I'm treating it as string manipulation and others are doing lots of overcomplicated maths. It's a sliding door problem at the end of the day!
