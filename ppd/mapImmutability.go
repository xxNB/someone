package main 

type S struct {
	name string
}

func main(){
	m := map[string]*S{"x": &S{"one"}}
	// println(m["x"].name)
	m["x"].name = "two"
}