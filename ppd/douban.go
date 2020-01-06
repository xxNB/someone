package main

import (
	"fmt"
	//"github.com/PuerkitoBio/goquery"
	//"log"
	//"reflect"

	//"log"
)

type DouMovie struct {
	newMovie []*newMovie
}

type newMovie struct {
	title string
	director string
	sorce string
	info string
	actor string
	date string
}

func (dou *DouMovie)newHignMovieSelect(){
	parmas := make(map[string]string)
	//parmas["type"] = "热门"
	//parmas["sort"] = "recommend"
	//parmas["page_limit"] = "20"
	//parmas["page_start"] = "0"
	jsonResponse := gojsonGet("https://movie.douban.com/j/search_subjects?type=movie&tag=%E7%83%AD%E9%97%A8&sort=recommend&page_limit=20&page_start=0", parmas)
	//fmt.Println(reflect.TypeOf(jsonResponse["subjects"]))
	fmt.Printf("%v", jsonResponse["subjects"])
	var movies []*newMovie
	rr:= jsonResponse["subjects"].([]interface{})
	for _, movie := range (rr){
		movieInfo := movie.(map[string]interface{})
		if movieInfo["rate"].(string)>"7"{
			newmovie := &newMovie{
			}
			newmovie.sorce =movieInfo["rate"].(string)
			newmovie.title = movieInfo["title"].(string)
			fmt.Println(newmovie.title)
			other := dou.movieDetails(newmovie, movieInfo["url"].(string))
			movies = append(movies, other)
		}
	}
	douMovie := &DouMovie{}
	douMovie.newMovie = movies
	for _, movie := range(douMovie.newMovie) {
		fmt.Println("\n ============================ \n")
		fmt.Printf("%v", movie)
	}
	return
}

func (dou *DouMovie) movieDetails(newmovie *newMovie, url string) *newMovie{
	parmas := make(map[string]string)

	jsonQuery := goqueyGet(url, parmas)
	//"#info > span:nth-child(1) > span.attrs > a"
	//"#info > span.actor > span.attrs > span:nth-child(1) > a"
	//"#info > span.actor > span.attrs > span:nth-child(2) > a"

	newmovie.director = jsonQuery.Find("a[rel='v:directedBy']").Text()
	newmovie.actor = jsonQuery.Find("a[rel='v:starring']").Text()
	//"#info > span:nth-child(17)"
	newmovie.date = jsonQuery.Find("span[property='v:initialReleaseDate']").Text()
	newmovie.info = jsonQuery.Find("#link-report > span").Text()
	//if exists{
	//	fmt.Println(dou.date)
	//}
	//fmt.Println(newmovie)
	return newmovie
}


func main()  {
	ss := &DouMovie{}

	ss.newHignMovieSelect()
	//ss.movieDetails("https://movie.douban.com/subject/30413052/")
}
