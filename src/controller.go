package src

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
)



type LoginForm struct {
	Name     string `form:"user" binding:"required"`
    Password string `form:"password" binding:"required"`
	Age int `form:"info" binding:"required"`
}

type Search struct{
	tags  string `form:"tags" binding:"required"`
	genres string `form:"genres" binding:"required"`
	countries string `form:"countries" binding:"required"`
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
       	//InsertUser(form.Name, form.Password, form.Age)
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
    }
}




func GetTop10(c *gin.Context){
	page := c.Param("page")
	pages, err := strconv.Atoi(page)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "code": 0, "res": err})
		return
	}
	res:=GetDoubanTop10(pages)
	c.JSON(http.StatusOK, gin.H{"status": "success", "code": 1, "res": res})
}

func GetShortCom(c *gin.Context)  {
	num := c.Param("num")
	page := c.Param("page")
	if num =="" || page==""{
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "code": 0, "res": nil})
		return
	}
	pages, _ := strconv.Atoi(page)
	nums, _ := strconv.Atoi(num)

	res := GetDetails(nums, pages)
	c.JSON(http.StatusOK, gin.H{"status": "success", "code": 1, "res": res})
}

func RedisTest(c *gin.Context)  {
	search := &Search{}
	res, err := UtilTest(search.tags, search.genres, search.countries)
	fmt.Println(err)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "code": 0, "res": fmt.Sprintf("%s", err)})
	}else{
	c.JSON(http.StatusOK, gin.H{"status": "success", "code": 1, "res": res})
	}
}
