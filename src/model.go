package src

import (
	"time"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type Db struct {
	User
	NewMovie
}

var x *xorm.Engine

type User struct {
	Id int64 `json:"id" xorm:"not null pk autoincr INT(10)"`
	Name string `json:"username" xorm:"not null default '' comment('用户名') VARCHAR(50)"`
	Age int `json:"age" xorm:"not null default 0 comment('年龄') INT(10)"`
	Passwd string `json:"password" xorm:"not null default '' comment('用户名') VARCHAR(50)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

type NewMovie struct {
	Id int64 `json:"id" xorm:"not null pk autoincr INT(10)"`
	User string `json:"username" xorm:"not null default '' comment('用户名') VARCHAR(50)"`
	Time string `json:"age" xorm:"not null default '' comment('时间') VARCHAR(50)"`
	Text string `json:"password" xorm:"not null comment('短评') TEXT"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

func GetEngine() (engin *xorm.Engine){
	if x==nil{
		host := "localhost"
		port := "3306"
		username := "root"
		password := "123456"
		database := "data"
		dataSourceName := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8", username, password, host, port, database)

		x, err := xorm.NewEngine("mysql", dataSourceName)
		    //建表
		if err = x.Sync2(new(NewMovie)); err != nil{
		        fmt.Println("error in create table user, ", err)
		    }
		if err != nil {
			panic(err)
		}
		return x
	}else {
		return x
	}

}




