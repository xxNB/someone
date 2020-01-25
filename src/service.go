package src

import (
	"net/http"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"encoding/json"
)

type Top10 struct{
	Title string
	Score float64
	Abstract string
}

func GetDoubanTop10()(tops []Top10){
	client := &http.Client{}
	url := "https://movie.douban.com/top250"
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	res := []Top10{}
	redi := &Redis{}
	c, err := redi.Get("doubantop10")
	if c != nil && err!=nil{
		err := json.Unmarshal(c.([]byte), &res)
		if err != nil {
			fmt.Println("Umarshal failed:", err)
			return
		}
		return res
	}
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36")
	response, _ := client.Do(reqest)
	document, err := goquery.NewDocumentFromReader(response.Body)
	document.Find("#content > div > div.article > ol > li").Each(func(i int, selection *goquery.Selection) {
		content := &Top10{}
		content.Title = selection.Find("div.hd > a > span").Text()
		content.Score, _ = strconv.ParseFloat(selection.Find("div > span.rating_num").Text(), 64)
		content.Abstract = selection.Find("div.bd > p.quote > span").Text()

		fmt.Println(content)
		res = append(res, *content)
	})
	b, err:=json.Marshal(res)
	if err != nil {
        fmt.Println("JSON ERR:", err)
    }
	redi.Set("doubantop10", string(b))
	return res
}

func GetDetails(num int) (res []map[string]interface{}){
	client := &http.Client{}
	url :=fmt.Sprintf("https://movie.douban.com/subject/%v/comments?sort=new_score&status=P", num)
	
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36")
	response, _ := client.Do(reqest)
	var detailsRes []map[string]interface{}
	document, err := goquery.NewDocumentFromReader(response.Body)
	// "#comments > div:nth-child(1) > div.comment > p > span"
	document.Find("#comments > div").Each(func(i int, selection *goquery.Selection) {
		content := make(map[string]interface{})
		content["user"] = selection.Find("h3 > span.comment-info > a").Text()
		content["time"]  = compressStr(selection.Find("span.comment-info > span.comment-time").Text())
		content["text"] = selection.Find("div.comment > p > span").Text()
		if content["text"] != ""{
			fmt.Println(content)
			detailsRes = append(detailsRes, content)
		}
	})
	return detailsRes
}




// func main(){
// 	GetDetails("美丽人生")
// }

