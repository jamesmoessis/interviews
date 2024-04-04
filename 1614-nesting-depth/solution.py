class Solution:
    def maxDepth(self, s: str) -> int:
        n = len(s)
        if n <= 1 and s not in ('(', ')'):
            return 0
        
        current_depth = 0
        max_depth = 0

        for c in s:
            if c == "(":
                current_depth += 1
                if current_depth > max_depth:
                    max_depth = current_depth
            elif c == ")":
                current_depth -= 1

        return max_depth