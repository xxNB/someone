package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
)

var (
	h   bool
	d   bool
	id  string
	m   string
	com bool
)

type Comment struct {
	user    string
	comment string
	vote    string
}

func init() {
	flag.BoolVar(&h, "h", false, "this help")

	flag.StringVar(&m, "m", "", "search movie")

	flag.BoolVar(&d, "d", false, "someting details")

	flag.StringVar(&id, "id", "", "book or movie id")

	flag.BoolVar(&com, "com", false, "comment")

	

	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, `doudouban version:1.0
		Usage: douban_api [-h help] [-m moviename] [-b bookname] [-d details] [-com comment]
		
		Options:
`)
	flag.PrintDefaults()
}

func searchMovie(name string) {
	url := fmt.Sprintf("https://movie.douban.com/j/subject_suggest?q=%s", name)
	id := getSearchId(url)
	fmt.Println(id)
}

func MovieDetailsByid(id string) {
	url := fmt.Sprintf("https://movie.douban.com/subject/%s", id)
	dom := urlToDocument(url)
	items := dom.Find("#info")
	items.EachWithBreak(func(index int, sel *goquery.Selection) bool {
		details := sel.Text()
		fmt.Println(details)
		return false
	})
	summary := dom.Find("#link-report > span")
	summary.EachWithBreak(func(index int, sel *goquery.Selection) bool {
		details := sel.Text()
		fmt.Println(details)
		return false
	})
}

func MovieCommentByid(id string) {
	url := fmt.Sprintf("https://movie.douban.com/subject/%s/comments?status=P", id)
	dom := urlToDocument(url)
	dom.Find("#comments > div").Each(func(i int, sel *goquery.Selection) {
		user := sel.Find("#comments > div > div.comment > h3 > span.comment-info > a").Text()
		comment := sel.Find("#comments > div > div.comment > p > span").Text()
		vote := sel.Find("#comments > div > div.comment > h3 > span.comment-vote > span").Text()
		if user != "" {
			fmt.Println("用户名:", user)
			fmt.Println("评论:", comment)
			fmt.Println("点赞数:", vote, "\n")
		}
	})
}

func getSearchId(url string) (id string) {
	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("NewRequest Error:", err)
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36")

	response, _ := client.Do(request)
	defer response.Body.Close()
	res, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
	}
	cn_json, _ := simplejson.NewJson(res)
	res_array, _ := cn_json.Array()
	if len(res_array) == 0 {
		fmt.Println("no this movie or book")
		return
	}
	keys := []string{}
	for k := range res_array[0].(map[string]interface{}) {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for index, di := range res_array {
		//取第一个
		newdi, _ := di.(map[string]interface{})
		for _, key := range keys {
			fmt.Printf("%s: %s\t", key, newdi[key])
		}
		if index == 0 {
			id = newdi["id"].(string)
		}
		fmt.Print("\n")
	}
	return id
}

func urlToDocument(url string) *goquery.Document {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("NewRequest Error:", err)
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36")
	request.Header.Add("Referer", "https://movie.douban.com/")
	request.Header.Add("Host", "movie.douban.com")

	response, _ := client.Do(request)
	if err != nil {
		error.Error(err)
	}
	//res , err := ioutil.ReadAll(response.Body)
	//if err !=nil{
	//	fmt.Println(err)
	//}
	//fmt.Println(string(res))
	dom, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		fmt.Println("NewDocumentFromReader Error:", err)
	}

	return dom

}

func main() {
	flag.Parse()

	if h {
		flag.Usage()
	}

	if m != "" {
		searchMovie(m)
	}

	if d && id != "" {
		MovieDetailsByid(id)
	}

	if com && id != "" {
		MovieCommentByid(id)
	}

	//MovieDetailsByid("27615233")

	//searchMovie("武林外传")

}
