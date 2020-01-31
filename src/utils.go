package src

import (
	"fmt"
	"net/http"
    "regexp"
    "io/ioutil"
    "encoding/json"
    "bytes"
    "github.com/pkg/errors"
    simplejson "github.com/bitly/go-simplejson"
    redis "github.com/gomodule/redigo/redis"

    // "reflect"

    // "strings"
)

// func DetailsRegex(name string, sourceStr string) (regexRes string){
// 	regexs := fmt.Sprintf("sid: (\d+), qcat: ''})'>%s", name)
// 	fmt.Println(regexs)
// 	r1:=regexp.MustCompile(regexs)
// 	b1:=r1.FindString(sourceStr)
// 	fmt.Println(b1)
// 	return b1
// }


type Rediss struct{
    conn redis.Conn
}

func getConn()(conn redis.Conn, err error){
    rediss := &Rediss{}
    if rediss.conn !=nil{
        return rediss.conn, nil
    }else{
        c, err := redis.Dial("tcp", "127.0.0.1:6379")
        if err != nil {
            fmt.Println(err)
            return nil, err
        }
        // defer c.Close()
        return c, nil
    }
}

func (rediss *Rediss) set(key string, value string){
    c, err := getConn()
    if err != nil{
    fmt.Println(err)
    }
    _, err = c.Do("SET", key, value)
    fmt.Println(err)
}


func (rediss *Rediss) get(key string) (res string, err error){
    c, err := getConn()
    if err != nil{
    fmt.Println(err)
    }
    return redis.String(c.Do("GET", key))
}


func compressStr(str string) string {
	if str == "" {
		return ""
	}
	//匹配一个或多个空白符的正则表达式
	reg := regexp.MustCompile("\\s+")
	return reg.ReplaceAllString(str, "")
}

func PostJson(url string, data map[string]string) (res *simplejson.Json, err error){

    redisCli := &Rediss{}
    
    fmt.Println("post a data successful.")
    redisRes, err := redisCli.get("doubanExample")
    if err!=nil{
        fmt.Println(err)
    }
    if redisRes != ""{
        // fmt.Printf("response data:%v\n",string(respBody))
        res, err := simplejson.NewJson([]byte(redisRes))
        return res, err
    }else{

        jsonStr, _ := json.Marshal(data)
        request, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
        request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36")
        //post数据并接收http响应
        resp,err :=http.DefaultClient.Do(request)

        if err!=nil{
            err = errors.WithStack(err)
            return nil, err
        }
        respBody,_ :=ioutil.ReadAll(resp.Body)
        fmt.Println(string(respBody))
        redisCli.set("doubanExample", string(respBody))
        res, err := simplejson.NewJson([]byte(string(respBody)))
        return res, err
    }
    }


func UtilTest() (res []interface{}){
    // ress := &Rediss{}
    // ress.set("zhangxin", "noway")
    // conn, err := ress.getConn()
    // if err!=nil{
    //     panic(err)
    // }
    // fmt.Println(conn)
    
    url := "https://movie.douban.com/j/new_search_subjects"
    data := map[string]string{
        "sort": "U",
        "range": "0,10",
        "tags": "电视剧",
        "start": "0",
        "genres": "悬疑",
        "countries": "美国",
    }
    redisRes, err := PostJson(url, data)
    if err!=nil{
        panic(err)
    }
    resDataList, err:= redisRes.Get("data").Array()
    if err !=nil{
        errors.WithStack(err)
        return
    }
    return resDataList

}
