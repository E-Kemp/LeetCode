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