package src

import (
	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)



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




func GetTop10(c *gin.Context){
	res:=GetDoubanTop10()
	c.JSON(http.StatusOK, gin.H{"status": "success", "code": 1, "res": res})
}

func GetShortCom(c *gin.Context)  {
	num, _ := strconv.Atoi(c.Param("num"))
	res := GetDetails(num)
	c.JSON(http.StatusOK, gin.H{"status": "success", "code": 1, "res": res})
}

func RedisTest(c *gin.Context)  {
	res := UtilTest()
	c.JSON(http.StatusOK, gin.H{"status": "success", "code": 1, "res": res})
}
