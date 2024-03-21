class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        if len(strs) == 0 :
            return ''
        if len(strs) == 1:
            return strs[0]

        prefix = ''
        for i in range(1, len(strs[0])+1):
            possible_prefix = strs[0][:i]
            if all(map(lambda s: s.startswith(possible_prefix), strs)):
                prefix = possible_prefix
            else:
                break
        return prefix
