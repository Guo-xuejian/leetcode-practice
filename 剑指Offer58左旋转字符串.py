class Solution:
    def reverseLeftWords(self, s: str, n: int) -> str:
        if n >= len(s) or n == 0:
            return s
        return s[n:] + s[:n]