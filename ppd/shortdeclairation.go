package main


func f()(x, y int){
	return 1, 2
}


func main(){
	var x int 
	x, _ := f()
	// x, _ = f()
	// x, y := f()
	// x, y = f()
}