package src

import (
	"net/http"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
)

type Top10 struct{
	Title string
	// Director string
	// Protagonist string
	Score float64
	// Time string
	Abstract string
}

func GetDoubanTop10()(tops []Top10){
	client := &http.Client{}
	url := "https://movie.douban.com/top250"
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36")
	response, _ := client.Do(reqest)
	document, err := goquery.NewDocumentFromReader(response.Body)
	res := []Top10{}
	document.Find("#content > div > div.article > ol > li").Each(func(i int, selection *goquery.Selection) {
		content := &Top10{}
		content.Title = selection.Find("div.hd > a > span").Text()
		// content.Director = selection.Find("div.comment > p > span").Text()
		// content.Protagonist = compressStr(selection.Find("span.comment-info > span.comment-time").Text())
		content.Score, _ = strconv.ParseFloat(selection.Find("div > span.rating_num").Text(), 64)
		content.Abstract = selection.Find("div.bd > p.quote > span").Text()
		// content.Time, err = strconv.Atoi(selection.Find("div.comment > h3 > span.comment-vote > span").Text())

		fmt.Println(content)
		//return false
		//return
		// time.Sleep(*time.Second)
		res = append(res, *content)
	})
	return res
}

// func main(){
// 	GetDoubanTop10()
// }