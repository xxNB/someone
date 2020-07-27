package main 

import (
	"io/ioutil"
	"os"
)

func main(){
	f, err := os.Open("file")
	defer f.close()

	if err != nil{
		return
	}

	b, err := ioutil.ReadAll(f)
	println(string(b))
}