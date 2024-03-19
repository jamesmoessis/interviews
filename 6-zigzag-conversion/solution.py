"""
https://leetcode.com/problems/zigzag-conversion/

PAYPALISHIRING

P   A   H   N
A P L S I I G
Y   I   R

PAHNAPLSIIGYIR

for 3 rows to get the first row, we hop 4 times, so n + 1


P     I    N
A   L S  I G
Y A   H R
P     I

PINALSIGYAHRPI

for 4 rows we hop 6 times, n+2

So what's the pattern? It has to go n-1 down, and then n-1 up.
So for n = 3, that's 2 + 2 = 4 hops.
So for n = 4, that's 3 + 3 = 6 hops.

So hops is (n-1)+(n-1) = 2n - 2

ahh but this is wrong it's not always the same amount of hops... 
need to revisit.

Ended up going for an easier strategy...
"""

class Solution:
    def convert(self, s: str, numRows: int) -> str:
        rows = []
        for i in range(numRows):
            rows.append([])
        increasing = True
        row = 0
        if numRows == 1:
            return s
        for c in s:
            rows[row].append(c)
            if increasing:
                row += 1
            else:
                row -= 1
            if row == numRows-1 or row == 0:
                increasing = not increasing
        sol = []
        for i in rows:
            for j in i:
                sol.append(j)
        return ''.join(sol)