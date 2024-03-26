def rec(s, n: int):
    # L + R
    if len(s) == n*2:
        return [s]
    lefts = 0
    rights = 0
    for c in s:
        if c == '(':
            lefts += 1
        elif c == ')':
            rights += 1
    if rights > lefts or lefts > n:
        return []
    sol = []
    if lefts+1 <= n:
        sol += rec(s+'(', n)
    if rights < lefts:
        sol += rec(s+')', n)
    return sol

class Solution:
    def generateParenthesis(self, n: int):
        return rec('(', n)
