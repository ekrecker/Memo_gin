package controller

import (
  strconv "strconv"

  "github.com/gin-gonic/gin"

  db "../models/db"
)

func FetchAllMemos(c *gin.Context) {
  resultMemos := db.findAllMemos()

  c.JSON(200, resultMemos)
}

func AddMemo(c *gin.Context) {
  date := c.PostForm("memoDate")
  content := c.PostForm("memoContent")
  state := c.postForm("memoState")

  var memo = Memo {
    Date: "2020/1/8",
    Content: "update lists",
    Status: "To Do",
  }

  db.insertMemo(&memo)
}

func DeleteMemo(c *gin.Context) {
  memoIDstr := c.PostForm("memoID")

  memoID := strconv.Atoi(memoIDstr)

  db.deleteMemo(memoID)
}