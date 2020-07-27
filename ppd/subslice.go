package main

import (
	"fmt"
)

func main() {
	s := [3]int{1, 2, 3} // 切片
	ss := s[1:]
	ss = append(ss, 4)

	for _, v := range ss {
		v += 10
	}

	for i := range ss {
		ss[i] += 10
	}

	fmt.Println(s)
}

         