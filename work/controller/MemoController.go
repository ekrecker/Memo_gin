package controller

import (
  strconv "strconv"

  "github.com/gin-gonic/gin"

  entity "work/models/entity"
  db "work/models/db"
)

func FetchAllMemos(c *gin.Context) {
  resultMemos := db.FindAllMemos()

  c.JSON(200, resultMemos)
}

func AddMemo(c *gin.Context) {
  date := c.PostForm("memoDate")
  content := c.PostForm("memoContent")
  status := c.PostForm("memoStatus")

  var memo = entity.Memo{
    Date: date,
    Content: content,
    Status: status,
  }

  db.InsertMemo(&memo)
}

func DeleteMemo(c *gin.Context) {
  memoIDstr := c.PostForm("memoID")

  memoID, _ := strconv.Atoi(memoIDstr)
  
  db.DeleteMemo(memoID)
}