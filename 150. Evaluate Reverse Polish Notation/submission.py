def spreadNewArray(tokens: List, replace: int, pos: int) -> List:
    return [*tokens[:pos-2], replace, *tokens[pos+1:]]

def getResult(tokens: List[str], pos: int) -> int:
    if len(tokens) == 1:
        return int(tokens[0])
    print(tokens)
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