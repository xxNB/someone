package main

import  (
	"fmt"
	"net/http"
	"log"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	tt "time"
)

type shortComment struct{
	Id     int64
	Author string  
	Comment string `text:"Comment" xorm:"varchar(1000) not null 'Comment'"` 
	Time string
}

var x *xorm.Engine

func GetEngin() *xorm.Engine{
	if x!=nil{
		return x
	}else{
	var err error
	fmt.Println("建立连接了～～～～")
	x, err := xorm.NewEngine("mysql", "root:@/data?charset=utf8mb4")
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	if err := x.Sync2(new(shortComment)); err != nil {
		log.Fatal("数据表同步失败:", err)
	}
	return x
}
}

func Insert(author string, comment string, time string) error {
	// ss := &shortComment{}
	// ss.Author = author
	// ss.Comment = comment
	// ss.Time = time
	var err error
	fmt.Println("开始了开始了", x)
	x = GetEngin()
	_, err = x.Insert(&shortComment{Author: author, Comment: comment, Time: time})
	return err
}


func GetAllurls() []string{
	// url_one := "https://movie.douban.com/subject/25887288/comments?start=%s&limit=20&sort=new_score&status=P"
	allUrls := []string{}
	for i:=0;i<1000;{
		fmt.Println(string(i))
		// fmt.Printf(url_one, string(i))
		url := fmt.Sprintf("https://movie.douban.com/subject/25887288/comments?start=%s&limit=20&sort=new_score&status=P", i)
		fmt.Println("shdoajd", url)
		allUrls = append(allUrls, url)
		fmt.Println("allUrls", allUrls)
		i=i+20
	}
	return allUrls
}


func main() {
	client:=&http.Client{}
	urls := GetAllurls()
	for _, i := range(urls){
		tt.Sleep(1 * tt.Second)
		fmt.Printf("正在爬取  url   %s", i)
	req, err := http.NewRequest("GET", i, nil)
	if err != nil{
		fmt.Println(err)
	}
    req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
    // req.Header.Add("Referer", url)
    // req.Header.Add("Cookie", "your cookie") // 也可以通过req.Cookie()的方式来设置cookie
    res, err := client.Do(req)
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromResponse(res)
	if err!=nil{
        log.Fatal("erroro%s", err)
	}
	fmt.Println("嘻嘻哈哈")
	// fmt.Println(doc.Html())
	// "#comments"
	doc.Find("#comments div.comment").Each(func(i int, s *goquery.Selection){
		fmt.Println("@@@@@@@@@@@@@@@")
		// fmt.Println(s.Html())
		people := s.Find("h3 > span.comment-info > a").Text()
		content := s.Find("div.comment  p  span").Text()
		time := s.Find("h3 > span.comment-info > span.comment-time").Text()
		fmt.Println("开始插入啊插入～～～")
		err := Insert(people, content, time)
		if err != nil{
			log.Fatal("erroro%s", err)
		}else{
			fmt.Println("插入成功~~~~")
		}
	})
}
}



