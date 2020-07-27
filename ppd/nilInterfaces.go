package main

type S struct{}

func (s S) F() {}

type IF interface{
	F()
}

func InitType() S{
	var s S 
	return s
}

func InitPointer() *S{
	var s *S 
	return s 
}

func InitEfaceType() interface{} {
	var s S
	return s
}

func InitEfacePointer() interface{} {
	var s *S
	return s
}

func InitIfaceType() IF{
	var s S 
	return s 
}

func InitIfacePointer() IF{
	var s *S 
	return s 
}

func main(){
	println(InitPointer() == nil)     // true   结构体指针
	println(InitEfaceType() == nil)    // false  结构体
	println(InitEfacePointer() == nil)  // false  结构体
	println(InitIfaceType() == nil)   // false   返回接口 
	println(InitIfacePointer() == nil) // false  返回接口指针
}  



