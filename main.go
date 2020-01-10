package main

import ("github.com/gin-gonic/gin"
		"someone/src"
)


func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
	r.POST("/addpeople", src.AddPeople)
    r.GET("/getTop10", src.GetTop10)
    r.GET("/getShortCom/:num", src.GetShortCom)
    r.Run() // listen and serve on 0.0.0.0:8080
}

