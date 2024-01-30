### Solution Retrospective

Well this was pretty slow, the average time of submission was 67ms, mine took 217ms. Let's explain why!

I first got excited looking at this problem of Reverse Polish Notation and thought I can recursively all a function to search the array for it's next operator. The more optimal way is to create a __new__ stack out of its tokens, and then pop the values as needed when an operator appears.

My submission (slow 😂):
```python
def spreadNewArray(tokens: List, replace: int, pos: int) -> List:
    return [*tokens[:pos-2], replace, *tokens[pos+1:]]

def getResult(tokens: List[str], pos: int) -> int:
    if len(tokens) == 1:
        return int(tokens[0])
    token = tokens[pos]
    replace = 0
    while pos < len(tokens):
        if token == "+":
            replace = int(tokens[pos-2]) + int(tokens[pos-1])
            return getResult(spreadNewArray(tokens, replace, pos), 0)
        elif token == "-":
            replace = int(tokens[pos-2]) - int(tokens[pos-1])
            return getResult(spreadNewArray(tokens, replace, pos), 0)
        elif token == "*":
            replace = int(tokens[pos-2]) * int(tokens[pos-1])
            return getResult(spreadNewArray(tokens, replace, pos), 0)
        elif token == "/":
            replace = int(int(tokens[pos-2]) / int(tokens[pos-1]))
            return getResult(spreadNewArray(tokens, replace, pos), 0)
        else:
            pos = pos + 1
            token = tokens[pos]
    return getResult(tokens, 0)

class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        return getResult(tokens, 0)
```

Sample (much moe Python-y!):
```python
class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        stack = []
        for t in tokens:
            if t not in ['+', '-', '*', '/']:
                stack.append(t)
            else:
                op1 = int(stack.pop())
                op2 = int(stack.pop())
                if t == '+':
                    stack.append(op1 + op2)
                elif t == '-':
                    stack.append(op2 - op1)
                elif t == '*':
                    stack.append(op2 * op1)
                else:
                    stack.append(int(float(op2)/op1))
        return int(stack[0])
```

There were some concepts and functions of Python that I missed in this, I wasn't aware of the the `.pop`, `.append` array methods in Python, and `if t in [x, y, z]` is a comparator I've never seen before, but would love to have in languages like Typescript!

Overall, I'm however happy with my solution, it gave me some good array manipulation experience and lots of thought into recursive functions. However reducing O(n) is important, and recursion doesn't necessariy make sense unless you're using a functional language (I wish I could use F# in this project!) or you're working with complex data structures like Trees. 
