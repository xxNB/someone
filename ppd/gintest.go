package main

import ("github.com/gin-gonic/gin"
        "net/http"
        "fmt"
)

// zhan wei fu
func path(c *gin.Context){
    name := c.Param("name")
    c.String(http.StatusOK, "hello %s", name)
}

// get ? &&

func GetParms(c *gin.Context){
    name := c.Query("lastname")
    xing := c.Query("firstname")
    c.String(http.StatusOK, "hello %s %s", xing, name)

}

// post

type LoginForm struct {
	User     string `form:"user" binding:"required"`
    Password string `form:"password" binding:"required"`
    Info []interface{} `form:"info" binding:"required"`
}

func PostParms(c *gin.Context)  {
    form := &LoginForm{}
    // message := c.BindJSON("message")
    // nick := c.PostForm("nick") 
    if c.BindJSON(&form) == nil {
        fmt.Println(form.Info)
        if form.User == "user" && form.Password == "password" {
            c.JSON(200, gin.H{"status": "you are logged in"})
        } else {
            c.JSON(401, gin.H{"status": "unauthorized"})
        }
    }
}


func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.POST("/tellmepost", PostParms)
    r.GET("/user/:name", path)
    r.GET("/parm", GetParms)
    r.Run() // listen and serve on 0.0.0.0:8080
}