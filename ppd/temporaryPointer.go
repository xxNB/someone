package main 

const N = 3 

func main(){
	m := make(map[int]*int)

	for i := 0; i< N; i++{
		// a := i
		println(&i)
		m[i] = &i
	}
	println(len(m))
	for _, v := range m {
		print(*v)
	}
}