package main

import (
	"fmt"
	//"github.com/PuerkitoBio/goquery"
	//"log"
	//"reflect"

	//"log"
)

func newHignMovieSelect() {
	parmas := make(map[string]string)
	//parmas["type"] = "热门"
	//parmas["sort"] = "recommend"
	//parmas["page_limit"] = "20"
	//parmas["page_start"] = "0"
	jsonResponse := gojsonGet("https://movie.douban.com/j/search_subjects?type=movie&tag=%E7%83%AD%E9%97%A8&sort=recommend&page_limit=20&page_start=0", parmas)
	//fmt.Println(reflect.TypeOf(jsonResponse["subjects"]))
	//fmt.Printf("%v", jsonResponse["subjects"])
	rr:= jsonResponse["subjects"].([]interface{})
	for _, movie := range (rr){
		movieInfo := movie.(map[string]interface{})
		if movieInfo["rate"].(string)>"7"{
			fmt.Println(movieInfo["title"])
		}
	}
}

func main()  {
	newHignMovieSelect()
}