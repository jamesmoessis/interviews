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
"""

class Solution:
    def convert(self, s: str, numRows: int) -> str:
        hops = 2*numRows - 2
        sol = ""
        for row in range(numRows):
            i = row
            while i < len(s):
                sol+=s[i]
                i += hops
        return sol

Solution().convert("PAYPALISHIRING", 3)