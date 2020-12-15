package db

type Memo struct {
  ID 			  int 	  `gorm:"primary_key;not_null"`
  Date      string 	`gorm:"type:varchar(20);not_null"`
  Content   string 	`gorm:"type:varchar(100);not_null"`
  Status    string  `gorm:"type:varchar(20);not_null"` 
}