package main 

func main(){
	m := make(map[string]int)
	m["a"] = 1
	println(m)
	// if _, v := m["b"]; v != false { //B
	// 	println(v)
	// }
	x, y := m["a"]  // 1, true
	println(x, y)
	
}