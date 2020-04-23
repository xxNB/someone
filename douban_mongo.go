package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"net/http"
	"strings"
	"time"
)

type BookContent struct {
	// id          bson.ObjectId `bson:"_id"`
	Book_Name    string    `json:"bookname" xorm:"not null default '' comment('书名') VARCHAR(50)"`
	Book_Comment string    `json:"bookcomment" xorm:"not null default '' comment('评论') VARCHAR(250)"`
	Book_Autor   string    `json:"bookautor" xorm:"not null default '' comment('作者') VARCHAR(50)"`
	Created      time.Time `xorm:"created"`
	Updated      time.Time `xorm:"updated"`
}

type Mysql struct {
	conn *xorm.Engine
}

func (sql *Mysql) getSqlcon() {
	if sql.conn != nil {
		return
	}
	fmt.Println("建立 mysql 连接")
	host := "localhost"
	port := "3306"
	username := "root"
	database := "data"
	dataSourceName := fmt.Sprintf("%s@(%s:%s)/%s?charset=utf8", username, host, port, database)

	x, err := xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	//建表
	//err = x.Sync2(new(BookContent))
	//if err != nil{
	//	fmt.Println("error in create table user, ", err)
	//}

	sql.conn = x
}

func (sql *Mysql) insert(name string, autor string, comment string) {
	sql.getSqlcon()
	session := sql.conn.NewSession()
	defer session.Close()
	err := session.Begin()
	item := &BookContent{}
	item.Book_Name = name
	item.Book_Autor = autor
	item.Book_Comment = comment
	_, err = sql.conn.Insert(item)
	if err != nil {
		fmt.Println(err)
		session.Rollback()
		return
	}
	err = session.Commit()
	fmt.Println("insert success~~")
	if err != nil {
		return
	}
}

func crawl(url string) {
	document := urlToDocument(url)
	urlist := []string{}
	document.Find("#subject_list > ul > li > div.info > h2 > a").Each(func(i int, selection *goquery.Selection) {
		url, _ := selection.Attr("href")
		urlist = append(urlist, url)
	})
	sql := &Mysql{}
	for _, url := range (urlist) {
		dom := urlToDocument(url)
		dom.Find("#wrapper").Each(func(i int, selection *goquery.Selection) {
			author := selection.Find("#info > span > a").Text()
			if author == "" {
				author = selection.Find("#info >  a").Text()
			}
			author = strings.Trim(author, "")
			comment := strings.Trim(selection.Find("#link-report > div > div > p").Text(), "")
			name := strings.Trim(selection.Find("#wrapper > h1 > span").Text(), "")
			fmt.Print(name, " ", author, " ", comment, )
			fmt.Println()
			go func() { sql.insert(name, author, comment) }()
		})
		time.Sleep(time.Second * 1)
	}
}

func urlToDocument(url string) *goquery.Document {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("NewRequest Error:", err)
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36")
	response, _ := client.Do(request)
	if err != nil {
		error.Error(err)
	}
	dom, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		fmt.Println("NewDocumentFromReader Error:", err)
	}

	return dom
}

func main() {
	for i := 20; i < 1000; {
		fmt.Println(i)
		crawl(fmt.Sprintf("https://book.douban.com/tag/历史?start=%v&type=T", i))
		i += 20
	}
}
