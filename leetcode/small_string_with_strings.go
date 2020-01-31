package main

import "fmt"

// func smallestStringWithSwaps(s string, pairs [][]int) string {
//     for _, i := range(pairs){
//         tmp := s[i[0]] 
//         s[i[0]] = s[i[1]]
//         s[i[1]] = tmp
// 	}
// 	fmt.Println(s)
//     return s
// }


func main(){
	// smallestStringWithSwaps("dcab", [[0,3],[1,2]])
    array := "zhangxin"

	array[1]="3"
	fmt.Println(array)
}