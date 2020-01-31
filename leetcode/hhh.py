class Solution:
    def smallestStringWithSwaps(self, s: str, pairs) -> str:
        pairs = sorted(pairs)
        for i in pairs:
            tmp1 = s[:i[0]]
            tmp2 = s[i[1]+1:]
            tmp3 = s[i[0]+1:i[1]]
            s = tmp1+s[i[1]]+tmp3+s[i[0]] + tmp2
            # breakpoint()
            print(s)
        return s

ss=Solution()

ss.smallestStringWithSwaps(s="dcab",pairs=[[0,3],[1,2],[0,2]])

