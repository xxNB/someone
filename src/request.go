package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	//"io"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)



func goqueyGet(url string, parms map[string]string) ( response *goquery.Document){
	client:=&http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil{
		fmt.Println(err)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
	// req.Header.Add("Referer", url)
	// req.Header.Add("Cookie", "your cookie") // 也可以通过req.Cookie()的方式来设置cookie
	res, err := client.Do(req)
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err!=nil{
		log.Fatal("erroro%s", err)
	}
	//fmt.Println(doc.Html())
	return doc
}

func gojsonGet(url string, parms map[string]string) ( response map[string]interface{}){
	client:=&http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil{
		fmt.Println(err)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
	// req.Header.Add("Referer", url)
	// req.Header.Add("Cookie", "your cookie") // 也可以通过req.Cookie()的方式来设置cookie
	res, err := client.Do(req)
	defer res.Body.Close()
	if err!=nil{
		log.Fatal("erroro%s", err)
	}
	//fmt.Println(doc.Html())
	content, err := ioutil.ReadAll(res.Body)
	if err != nil{
		log.Fatal(err)
	}
	stringRes := string(content)
	//fmt.Println(stringRes)
	err = json.Unmarshal([]byte(stringRes), &response)
	if err != nil{
		log.Fatal(err)
	}
	//fmt.Printf("%v", response)
	//return response
	return
}