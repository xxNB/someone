package main

import "fmt"

func Test()  {
	subMapA := map[string]interface{}{"name":"jack","cid":123,"pid":"abc","number":"123456"}
	sub := map[string]interface{}{"duid":"dadadsa"}
	subMapA["token"] = sub
	fmt.Println(subMapA)
	for i, _ := range(subMapA){
		if(i=="cid"||i=="pid"){
			sub[i] = subMapA[i]
			// sub["pid"] = subMapA["pid"]
			subMapA["token"] = sub
			delete(subMapA, i)
		}
	fmt.Println(subMapA)
	}
}

func Test1()  {
	koko := make(map[string]map[string]interface{})
	koko["da"] = 123
	

	
}


func main()  {
	Test()
}