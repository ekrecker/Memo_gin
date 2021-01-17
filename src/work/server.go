package main

import (
	//"fmt"
	"log"
  "net/http"

  "github.com/gin-gonic/gin"
  _ "github.com/jinzhu/gorm/dialects/mysql"

  controller "work/controller"
)

func main () {
  serve()
}

func serve () {
  router := gin.Default()

  router.Static("/views", "./views")
  router.StaticFS("/main", http.Dir("./views/static"))
  /*
  router.LoadHTMLGlob("views/static/index.html")
  router.GET("/", func(c *gin.Context){
    c.HTML(http.StatusOK, "/views/static/index.html", gin.H{})
  })
  */

  router.GET("/fetchAllMemos", controller.FetchAllMemos)

  router.POST("/addMemo", controller.AddMemo)

  router.POST("/deleteMemo", controller.DeleteMemo)

  if err := router.Run(":8080"); err != nil {
    log.Fatal("server Run Failed.: ", err)
  }
}