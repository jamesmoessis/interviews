# ([])
# allowed a new open bracket, or close the 'current'
# keep a stack
# put bracket on stack, need to be matched for it to be popped

def isOpen(s):
    return s in ('(', '[', '{')

def mirrors(open, close):
    if open == '(':
        return close == ')'
    if open == '[':
        return close == ']'
    if open == '{':
        return close == '}'

class Solution:
    def isValid(self, s: str) -> bool:
        stack = []

        for c in s:
            if isOpen(c):
                stack.append(c)
                continue
            if len(stack) == 0:
                return False
            last_open = stack[-1]
            if not mirrors(last_open, c):
                return False
            else:
                stack.pop(-1)
        return len(stack) == 0
