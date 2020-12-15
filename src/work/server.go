package main

import (
	//"fmt"
  //"net/http"

  "github.com/gin-gonic/gin"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

func main () {
  serve()
}

func serve () {
  router := gin.Default()
  router.LoadHTMLGlob("views/static/index.html")

  router.GET("/", func(c *gin.Context){
    c.HTML(200, "index.html", gin.H{})
  })

  router.Run(":8080")
}