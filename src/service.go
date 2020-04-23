package src

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
)

type Top10 struct {
	Title    string
	Score    float64
	Abstract string
}

func GetDoubanTop10(n int) (tops []Top10) {
	client := &http.Client{}
	url := fmt.Sprintf("https://movie.douban.com/top250?start=%v&filter=", (n-1)*25)
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	res := []Top10{}
	//redi := &Rediss{}
	//c, err := redi.get("doubantop10")
	//if c != "" && err!=nil{
	//	err := json.Unmarshal([]byte(c), &res)
	//	if err != nil {
	//		fmt.Println("Umarshal failed:", err)
	//		return
	//	}
	//	return res
	//}
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36")
	response, _ := client.Do(reqest)
	document, err := goquery.NewDocumentFromReader(response.Body)
	document.Find("#content > div > div.article > ol > li").Each(func(i int, selection *goquery.Selection) {
		content := &Top10{}
		content.Title = selection.Find("div.hd > a > span").Text()
		content.Score, _ = strconv.ParseFloat(selection.Find("div > span.rating_num").Text(), 64)
		content.Abstract = selection.Find("div.bd > p.quote > span").Text()
		res = append(res, *content)
	})
	//b, err:=json.Marshal(res)
	if err != nil {
		fmt.Println("JSON ERR:", err)
	}
	//redi.set("doubantop10", string(b))
	return res
}

func GetDetails(num int, page int) (res []map[string]string) {
	client := &http.Client{}
	url := fmt.Sprintf("https://movie.douban.com/subject/%v/comments?start=%d&limit=20&sort=new_score&status=P", num, (page-1)*25)
	engine := GetEngine()
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36")
	response, _ := client.Do(reqest)
	var detailsRes []map[string]string
	document, err := goquery.NewDocumentFromReader(response.Body)
	document.Find("#comments > div").Each(func(i int, selection *goquery.Selection) {
		content := make(map[string]string)
		content["user"] = selection.Find("h3 > span.comment-info > a").Text()
		content["time"] = compressStr(selection.Find("span.comment-info > span.comment-time").Text())
		content["text"] = selection.Find("div.comment > p > span").Text()
		if content["text"] != "" {
			detailsRes = append(detailsRes, content)
			go func() {
				_, err := engine.Insert(&NewMovie{User: content["user"], Time: content["time"], Text: content["text"]})
				if err !=nil{
					fmt.Println("Db insert error", err)
				}
			}()
		}
	})
	return detailsRes
}
