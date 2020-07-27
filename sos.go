package main 

import "fmt"

func main(){
	s := "hello"
	for _, i := range s{
		fmt.Println(i)
		fmt.Printf("%q", i)
	}
}