package main

import (
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery"
	"net/http"
	"regexp"
	"strconv"
)

type content struct {
	user    string
	star    string
	comment string
	userful int
	date    string
}

func ways(url string) {
	client := &http.Client{}

	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36")
	response, _ := client.Do(reqest)
	document, err := goquery.NewDocumentFromReader(response.Body)
	document.Find("#comments div").Each(func(i int, selection *goquery.Selection) {
		fmt.Println("=======================")
		content := &content{}
		content.user = selection.Find("div.comment > h3 > span.comment-info > a").Text()
		content.comment = selection.Find("div.comment > p > span").Text()
		content.date = compressStr(selection.Find("span.comment-info > span.comment-time").Text())
		content.star, _ = selection.Find("span.comment-info > span.allstar30.rating").Attr("title")
		content.userful, err = strconv.Atoi(selection.Find("div.comment > h3 > span.comment-vote > span").Text())
		fmt.Println(content)
		//return false
		//return
		time.Sleep(2*time.Second)
	})

}

func compressStr(str string) string {
	if str == "" {
		return ""
	}
	//匹配一个或多个空白符的正则表达式
	reg := regexp.MustCompile("\\s+")
	return reg.ReplaceAllString(str, "")
}

// func main() {
// 	for i := 0; i < 2000; {
// 		var url string
// 		url = fmt.Sprintf("https://movie.douban.com/subject/30135113/comments?start=%s&limit=20&sort=new_score&status=P", i)
// 		ways(url)
// 		i = i + 20
// 	}
// }
