package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var res int = 0
	// sByte := []byte(s)
	if len(s) == 0 {
		return res
	}
	hashMap := make(map[byte]int)
	var left, right int = 0, 0
	for right < len(s) {
		if _, ok := hashMap[s[right]]; ok {
			left = max(left, hashMap[s[right]]+1)
		}
		hashMap[s[right]] = right
		res = max(res, right-left+1)
		right++
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(lengthOfLongestSubstring("hyhyhyhy"))
}
