package src

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"time"
	// "log"
	// "src"
)

type User struct {
	Id int64 `json:"id" xorm:"not null pk autoincr INT(10)"`
	Name string `json:"username" xorm:"not null default '' comment('用户名') VARCHAR(50)"`
	Age int `json:"age" xorm:"not null default 0 comment('年龄') INT(10)"`
	Passwd string `json:"password" xorm:"not null default '' comment('用户名') VARCHAR(50)"`
	// phone string
	// email string
	// isactive int64 `json:"isactive" xorm:"not null default 1 comment('用户名') INT(1)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}


type LoginForm struct {
	Name     string `form:"user" binding:"required"`
    Password string `form:"password" binding:"required"`
	Age int `form:"info" binding:"required"`
}

func AddPeople(c *gin.Context)  {
    form := &LoginForm{}
    // message := c.BindJSON("message")
    // nick := c.PostForm("nick") 
    if c.BindJSON(&form) == nil {
		if (form.Name == ""){
			c.JSON(http.StatusForbidden, gin.H{"status": "username must be needed"})
		}
		if (form.Password == ""){
			c.JSON(http.StatusForbidden, gin.H{"status": "passwod must be needed"})
		}
		fmt.Println("&&&&&&&&")
       	InsertUser(form.Name, form.Password, form.Age)
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
    }
}



var x *xorm.Engine

func GetEngine() (engin *xorm.Engine){
	if x==nil{
		host := "localhost"
		port := "3306"
		username := "root"
		// password := "123456"
		database := "data"
		dataSourceName := fmt.Sprintf("%s@(%s:%s)/%s?charset=utf8", username, host, port, database)

		x, err := xorm.NewEngine("mysql", dataSourceName)
		//     //建表
		//     err = engine.Sync2(new(User))
		//     if err != nil{
		//         fmt.Println("error in create table user, ", err)
		//     }
		// }
		if err != nil {
			panic(err)
		}
		return x
	}else {
		return x
	}

}


func InsertUser(name string, password string, age int){
	fmt.Println("engine", x)

	x = GetEngine()
	session := x.NewSession()
	//session := engine.NewSession()
	defer session.Close()
	// add Begin() before any action
	err := session.Begin()
	user := &User{}
	user.Name = name
	user.Passwd = password
	user.Age =age
	_, err = x.Insert(user)
	if err != nil {
		fmt.Println(err)
		session.Rollback()
		return
	}
	err = session.Commit()
	if err != nil {
		return
	}

}


func GetTop10(c *gin.Context){
	res:=GetDoubanTop10()
	c.JSON(http.StatusOK, gin.H{"status": "success", "code": 1, "res": res})
}

