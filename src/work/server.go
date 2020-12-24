package main

import (
	//"fmt"
  //"net/http"

  "github.com/gin-gonic/gin"
  _ "github.com/jinzhu/gorm/dialects/mysql"

  controller "./controllers/"
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

  router.GET("./fetchAllMemos", controller.FetchAllMemos)

  router.POST("./addMemo", controller.AddMemo)

  router.POST("./deleteMemo", controller.deleteMemo)

  if err := router.Run(":8080"); != nil {
    log.Fatal("server Run Failed.: ", err)
  }
}