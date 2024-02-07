### Solution Retrospective

**Date:** 07/02/2024  
**Language:** Typescript  
**Difficulty:** Medium (apparently)

Looking at this problem, the solution seemed (and turned out to be) very easy to understand. All I needed to do was create a 'counter' record of characters, sort the record by frequency, then print a string with the characters repeated by their frequency. There's probably a cheesier way to do this in another language, but TypeScript doesn't have a native `character` type, whereas in languages like C# where strings are basically an array of characters - which would've made it possibly easier to manipulate.

But anyway, here's the solution:

```typescript
function frequencySort(s: string): string {
    const counter: Record<string, number> = {}

    for(let i = 0; i < s.length; i++) {
        if(counter[s[i]] === undefined) {
            counter[s[i]] = 1;
        }
        else counter[s[i]] = counter[s[i]] + 1;
    }

    let outputStr = "";
    Object.entries(counter)
        .map(([k, v]) => ({
            char: k,
            count: v
        }))
        .sort((a, b) => b.count - a.count)
        .forEach((counter) => {
            outputStr = outputStr.concat(counter.char.repeat(counter.count));
        })
    
    return outputStr;
};
```

This took about 20 minutes to write up, I need to start doing this in more challenging languages. Enjoy! 😂
