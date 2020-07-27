package src

import (
	"bytes"
	"encoding/json"
	"fmt"
	simplejson "github.com/bitly/go-simplejson"
	redis "github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"regexp"
)

type Rediss struct {
	conn redis.Conn
}

func getConn() (conn redis.Conn, err error) {
	rediss := &Rediss{}
	if rediss.conn != nil {
		return rediss.conn, nil
	} else {
		c, err := redis.Dial("tcp", "127.0.0.1:6379")
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		// defer c.Close()
		return c, nil
	}
}

func (rediss *Rediss) set(key string, value string) error {
	c, err := getConn()
	if err != nil {
		fmt.Println(err)
	}
	_, err = c.Do("SET", key, value)
	// errors.New("something append")
	if err != nil {
		return errors.Wrap(err, "test set bug ")
	}
	return nil

}

func (rediss *Rediss) get(key string) (res string, err error) {
	c, err := getConn()
	if err != nil {
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

func PostJson(url string, data map[string]string) (res *simplejson.Json, err error) {
	redisCli := &Rediss{}
	fmt.Println("post a data successful.")
	redisRes, err := redisCli.get("doubanExample")
	if redisRes != "" {
		// fmt.Printf("response data:%v\n",string(respBody))
		res, err := simplejson.NewJson([]byte(redisRes))
		if err != nil {
			return res, errors.Wrap(err, "json err")
		}
		return res, nil
	} else {

		jsonStr, _ := json.Marshal(data)
		request, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36")
		//post数据并接收http响应
		resp, err := http.DefaultClient.Do(request)
		if err != nil {
			return nil, errors.Wrap(err, "http err")
		}
		respBody, _ := ioutil.ReadAll(resp.Body)
		err = redisCli.set("doubanExample", string(respBody))
		if err != nil {
			return nil, err
		}
		res, err := simplejson.NewJson([]byte(string(respBody)))
		return res, err
	}
}


func UtilTest(tags, genres, countries string) (res []interface{}, err error, ) {

	url := "https://movie.douban.com/j/new_search_subjects"
	data := map[string]string{
		"sort":      "U",
		"range":     "0,10",
		"tags":      tags,
		"start":     "0",
		"genres":    genres,
		"countries": countries,
	}
	redisRes, err := PostJson(url, data)
	if err != nil {
		return nil, err
	}
	resDataList, err := redisRes.Get("data").Array()
	return resDataList, err

}
