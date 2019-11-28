package main 

import (
    "fmt"
    "log"
    "net/http"
    // "string"
    // "sync"
    "time"
    "github.com/PuerkitoBio/goquery"
    _ "github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
)

type Movies struct {
    Id              int64 `xorm:"autoincr pk"`
    Title           string  // 书名
    Url             string  // url
    PicUrl          string  // pic url
    Actor            string  // 作者
    // Fraction        float64 // 豆瓣评分
    Fraction        string  // 豆瓣评分
    Type           string  // 出版社
    Area   string  // 原作名
    Language     string  // 译者
    Date string  // 出版年
    // Page            int64   // 页数
    Name            string   // 页数
    // Price float64            // 定价
    IMDb          string    // 定价
    Created        time.Time `xorm:"created"`
    Updated        time.Time `xorm:"updated"`
}

// #content > div > div.article > div > div.list-wp > div
func Page(engine *xorm.Engine, url string){
    fmt.Println("~~~~~~~~~~~~~~~")
    // baseUrl:="http://baidu.com"
    client:=&http.Client{}
    req, err := http.NewRequest("GET", url, nil)
    req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
    // req.Header.Add("Referer", url)
    // req.Header.Add("Cookie", "your cookie") // 也可以通过req.Cookie()的方式来设置cookie
    res, err := client.Do(req)
    defer res.Body.Close()
    doc, err := goquery.NewDocumentFromResponse(res)
    if err!=nil{
        log.Fatal("erroro%s", err)
    }
    fmt.Println(doc.Html())

    doc.Find("div.list-wp div").Each(func(i int, s *goquery.Selection){
        fmt.Println("@@@@@@@@@@@@@@@")
        href, _ := s.Find("a.item").Attr("href")
        // go 一个movie
        fmt.Println("@@@@@@@@@@@@@@@")
        fmt.Println(href)
        Movie(engine, href)
        time.Sleep(1 * time.Second)
    })
}


func Movie(engine *xorm.Engine, url string){
    doc, err := goquery.NewDocument(url)
    if err != nil {
        log.Fatal(err)
    }

    data := &Movies{}
    data.Title = doc.Find("h1 span").Text()
    fmt.Printf("moviename: %s\n", data.Title)
}

func main() {
    engine, err := xorm.NewEngine("mysql", "root:@/data?charset=utf8")  
    if err != nil {
        log.Fatal(err)
    }
    err = engine.Sync2(new(Movies))    // 同步表结构
    if err != nil {
        log.Fatal(err)
    }
    url:="https://movie.douban.com/explore#!type=movie&tag=热门&sort=recommend&page_limit=20&page_start=0"
    Page(engine, url)
}
