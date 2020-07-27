package main

import (
	"fmt"
	"io/ioutil"
	// "time"
	"net/http"
)

func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
    type item struct {
        thumbfile string
        err       error
    }

    ch := make(chan item, len(filenames))
    for _, f := range filenames {
        go func(f string) {
            var it item
            it.thumbfile, it.err = https(f)
            ch <- it
        }(f)
    }

    for range filenames {
		// 消费 buffer
        it := <-ch
        if it.err != nil {
            return nil, it.err
        }
        thumbfiles = append(thumbfiles, it.thumbfile)
    }

    return thumbfiles, nil

}


func https(url string)(resp string, err error){
	res, err := http.Get(url)
	rr, err := ioutil.ReadAll(res.Body)
	return string(rr), err
}


func main(){
	res, err := makeThumbnails5([]string{"http://www.baidu.com","http://www.douban.com","http://www.baidu.com"})
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(res)
}