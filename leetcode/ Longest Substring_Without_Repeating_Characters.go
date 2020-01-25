func lengthOfLongestSubstring(s string) int {
    n, ans := len(s), 0
    maps := make(map[byte]int)
    for j,i:=0,0 ; j<n; j++{
        if _,ok := maps[s[j]]; ok{
            i = max(maps[s[j]], i)
        }
        ans = max(ans, j-i+1)
        maps[s[j]]= j+1
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}