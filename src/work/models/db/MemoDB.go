package db

import (
  "fmt"

  "github.com/jinzhu/gorm"

  entity "work/models/entity"
)

func GetGormConnect() *gorm.DB {
  DBMS := "mysql"
  USER := "admin"
  PASS := "password"
  PROTOCOL := "tcp(db:3306)"
  DBNAME := "go_database"
  CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
  db, err := gorm.Open(DBMS, CONNECT)

  if err != nil {
    panic(err.Error())
  }

  db.Set("gorm:table_options", "ENGINE=InnoDB")

  db.LogMode(true)

  //db.SingularTable(true)

  db.AutoMigrate(&entity.Memo{})

  fmt.Println("db connected: ", &db)
  return db
}

func InsertMemo(registerMemo *entity.Memo) {
  db := GetGormConnect()

  db.Create(&registerMemo)
  defer db.Close()
}

func DeleteMemo(memoID int) {
  memo := []entity.Memo{}

  db := GetGormConnect()
  db.Debug().Delete(&memo, memoID)
  defer db.Close()
}

func FindAllMemos() []entity.Memo {
  memos := []entity.Memo{}

  db := GetGormConnect()
  db.Order("ID asc").Find(&memos)
  defer db.Close()
  return memos
}