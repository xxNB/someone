package main

import "fmt"

func romanToInt(s string) int {
	a := map[byte]int{
		'I':1,
		'V':5,
		'X':10,
		'L':50,
		'C':100,
		'D':500,
		'M':1000,
	}
	ans := 0
	for i := range s{
		if (i <len(s) -1) && (a[s[i]] < a[s[i+1]]){
			ans -= a[s[i]]
		}else{
			ans += a[s[i]]
		}
	}
	return ans
}

func main(){
	fmt.Println(romanToInt("III"))
}