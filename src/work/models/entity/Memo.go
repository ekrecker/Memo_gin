package entity

type Memo struct {
  ID 			  int 	  `gorm:"primary_key;not_null"       json:"id"`
  Date      string 	`gorm:"type:varchar(20);not_null"  json:"date"`
  Content   string 	`gorm:"type:varchar(100);not_null" json:"content"`
  Status    string  `gorm:"type:varchar(20);not_null"  json:"status"`
}