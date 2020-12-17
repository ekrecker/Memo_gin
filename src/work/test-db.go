package main

import (
  "fmt"

  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Memo struct {
  ID 			  int 	  `gorm:"primary_key;not_null"`
  Date      string 	`gorm:"type:varchar(20);not_null"`
  Content   string 	`gorm:"type:varchar(100);not_null"`
  Status    string  `gorm:"type:varchar(20);not_null"` 
}

func getGormConnect() *gorm.DB {
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

  db.AutoMigrate(&Memo{})

  fmt.Println("db connected: ", &db)
  return db
}

func insertMemo(registerMemo *Memo) {
  db := getGormConnect()

  db.Create(&registerMemo)
  defer db.Close()
}

func deleteMemo(registerMemo *Memo) {
  db := getGormConnect()

  db.Debug().Delete(&registerMemo)
  defer db.Close()
}

func findAllMemo() []Memo {
  db := getGormConnect()
  var memos []Memo

  db.Order("ID asc").Find(&memos)
  defer db.Close()
  return memos
}

func main() {

 
  var memo = Memo {
    Date: "2020/11/18",
    Content: "テスト",
    Status: "To Do",
  }
	
	
  /*
  var memo = Memo {
    ID: 1,
  }
	*/
	
	
  insertMemo(&memo)
  //deleteMemo(&memo)

  resultMemos := findAllMemo()

  for i := range resultMemos {
    fmt.Printf("index: %d, ID: %d, Date: %s, Content: %s, Status: %s\n",
              i, resultMemos[i].ID, resultMemos[i].Date, resultMemos[i].Content, resultMemos[i].Status)
  }

}