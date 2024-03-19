
"""
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000


90 = XC
109 = CIX
CXIX
"""


"""
This is the shittest way to approach this problem, I'm keeping this in here for a lesson! 
See the go solution for better! 
"""
class Solution:
    def intToRoman(self, num: int) -> str:
        sol = ""

        lowest_digit = num % 10
        if lowest_digit == 0:
            pass
        elif lowest_digit <= 3:
            sol += lowest_digit*"I"
        elif lowest_digit == 4:
            sol += "IV"
        elif lowest_digit <= 8:
            sol += "V" + (lowest_digit%5) * "I"
        else:
            sol += ((10-lowest_digit)*"I") + "X"

        hundreds = num % 100
        if hundreds <= 9:
            pass
        elif hundreds <= 19:
            sol = "X" + sol
        elif hundreds <= 29:
            sol = "XX" + sol
        elif hundreds <= 39:
            sol = "XXX" + sol
        elif hundreds <= 49:
            sol = "XL" + sol
        elif hundreds <= 59: 
            sol = "L"  + sol
        elif hundreds <= 69: 
            sol = "LX" + sol
        elif hundreds <= 79:
            sol = "LXX" + sol
        elif hundreds <= 89:
            sol = "LXXX" + sol
        else:
            sol = "XC" + sol

        return sol
    
for i in range(0, 100):
    print(f"{i}: " + Solution().intToRoman(i))