# ()
# ()()     (())
# 
# ((()))   (())()   ()(()) ()()() (()())
# (((()))) ((()))() (())(()) ()((()))


# To get n, we take each solution to n-1 and:
# 1. put parenthese around it 
# 2. Put to the left 
# 3. put to the right
# 4. at the very end, de-duplicate everything

def generateSet(n): 
    if n == 1:
        return ["()"]
    subsol = generateSet(n-1)
    while len(subsol[0]) == (n-1)*2:
        sol = subsol.pop(0)
        subsol.append('()' + sol)
        subsol.append(sol + '()')
        subsol.append('(' + sol + ')')
    return subsol


class Solution:
    def generateParenthesis(self, n: int):
        return list(set(generateSet(n)))
    
if __name__ == '__main__':
    print(Solution().generateParenthesis(3))
    print(sorted(Solution().generateParenthesis(4)))

#["(((())))","((()()))","((())())","((()))()","(()(()))","(()()())","(()())()","(())()()","()((()))","()(()())","()(())()","()()(())","()()()()"]
#["(((())))","((()()))","((())())","((()))()","(()(()))","(()()())","(()())()","(())(())","(())()()","()((()))","()(()())","()(())()","()()(())","()()()()"]
# "(())(())(())"
# 1 -> 1 
# 2 -> 2
# 3 -> 5  = (2 * 3) - 1 
# 4 -> 14 = (5 * 3) - 1
# 5 -> 42 = (14*3)