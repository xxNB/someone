package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"encoding/json"

)

type Redis struct {
	conn redis.Conn
	uri string
	// time int64
}


func (redi *Redis)ConConect()  {
	if redi.conn != nil{
		return
	}
	conn,err := redis.Dial("tcp","127.0.0.1:6379")
    if err != nil {
        fmt.Println("connect redis error :",err)
        return
    }
	defer conn.Close()
	redi.conn = conn
}

func (redi *Redis)Set(key string, value string) (error){
	redi.ConConect()
	fmt.Println(key, value)
	_, err := redi.conn.Do("SET", key, value)
	if err != nil{
		return fmt.Errorf("redis set error %s", err)
	}
	return nil
}


func (redi *Redis)Get(key string) (value string, err error){
	redi.ConConect()
	res, err := redis.String(redi.conn.Do("GET", key))
	if err != nil{
		return "", fmt.Errorf("redis set error %s", err)
	}
	return res, nil
}

type Res struct {
	Aiai string
	Bibi []int
}

func main(){
	redi := &Redis{}
	res := Res{Aiai:"peace && love", Bibi: []int{1,2,3,4}}
	b, err:=json.Marshal(res)
	fmt.Println(string(b))
	if err != nil {
        fmt.Println("JSON ERR:", err)
    }
	err = redi.Set("doubantop10", string(b))
	fmt.Println(err)
	ress, _:= redi.Get("doubantop10")
	fmt.Println(ress)
	// redi.Set("name", )
}