package ppd

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type Redis struct{
	// conn redis.conn
	// time string
}


func (rediss *Redis)GetRedisCon() (redis.conn){
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()
	return c
}

func (rediss *Redis) SetString(key string, value string) error{
	red := &Redis{}
	conn := redis.GetRedisCon()
	_, err := conn.Do("set", "url", "xxbandy.github.io")
	if err != nil{
		return err
	}
	return nil
}


// func main(){
// 	red := &Redis{}
// }